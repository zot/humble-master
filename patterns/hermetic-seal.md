---
name: Hermetic Seal
tags: agentic, subagents, tool-discipline, hooks, weak-models, foundational
summary: Constrain a subagent's tool access through narrative framing plus a PreToolUse hook — the prompt shapes intent, the hook enforces boundaries
---

# Hermetic Seal

A subagent that can only interact with the world through a specific CLI tool,
enforced at two layers: a narrative framing that shapes the model's intent, and
a PreToolUse hook that blocks everything else.

## The Problem

Weak models (Haiku) ignore tool restrictions in prompts. Eight iterations of
rules, world-models, consequences, and identity framing all failed — Haiku
follows the rules for 1-2 calls, then escapes to familiar tools (find, ls,
Read, grep) when results don't satisfy the query. `tools:` frontmatter
allowlists and `disallowedTools:` don't work — the system provides base tools
regardless.

## The Solution: Two Layers

### Layer 1: Narrative framing (shapes intent)

Don't say "don't use X." Instead, frame the approved tool as the *technique*
that works. The model reaches for the right tool first because the narrative
makes it the obviously correct choice.

```markdown
**Your tools are `~/.ark/ark` commands — search, fetch, message, files.**
Your job is to determine whether the answer exists in the collection
using these tools, and if so, what it is. Both outcomes are equally
valuable: "here's what it says" and "it's not in the collection" are
both successful results.

Files in other projects are not directly accessible — attempting to
access them will usually fail. The hermetic technique: `ark fetch`
retrieves any indexed file regardless of which project it belongs to.
```

Key principles:
- **"Not found" is a successful outcome** — removes the frustration that
  drives tool escape. The model isn't failing, it's answering.
- **Name the technique** — "hermetic technique" gives the approved path
  an identity. The model copies named patterns more reliably than rules.
- **Concrete examples** — show actual commands, not abstract descriptions.
  Haiku copies examples better than it follows instructions.
- **Word depriming** — never use tool names as English words in the prompt.
  "find" activates `find`. "read" activates `Read`. Use "search for",
  "uncover", "retrieve", "fetch" instead.

### Layer 2: PreToolUse hook (enforces boundaries)

The hook catches what the narrative misses. It runs on every tool call and
blocks anything that isn't an approved CLI command.

```yaml
# In the agent's frontmatter
hooks:
  PreToolUse:
    - matcher: "Bash|Read|Grep|Glob|Search"
      hooks:
        - type: command
          command: "$CLAUDE_PROJECT_DIR/scripts/guard.sh"
```

The guard script:

```bash
#!/bin/bash
INPUT=$(cat)
TOOL=$(echo "$INPUT" | jq -r '.tool_name')

if [ "$TOOL" = Bash ]; then
    CMD=$(echo "$INPUT" | jq -r '.tool_input.command')
    if echo "$CMD" | grep -q '^\s*~/.ark/ark\b'; then
      exit 0  # allow
    fi
fi

jq >&2 -n '{
  hookSpecificOutput: {
    hookEventName: "PreToolUse",
    permissionDecision: "deny",
    permissionDecisionReason: "Use ~/.ark/ark commands instead."
  }
}'
exit 2
```

The denial message teaches: when the model hits the wall, the reason tells it
what to do instead. Haiku actually follows this guidance — in testing, a
blocked `find` was immediately followed by `ark files` with the correct glob.

## The Guard as Bootstrap

The guard script also serves as the subagent's bootstrap mechanism via
**Fumble Onboarding** (see fumble-onboarding.md). Lifecycle hooks don't
fire in subagents, so the first denied tool call carries a crank-handle
instruction telling the agent to fetch its own skill reference. The
agent definition stays clean — persona and guard only — and works from
any project directory.

## Why Two Layers

Either layer alone is insufficient:
- **Narrative only**: Haiku escapes after 2-3 calls when frustrated
- **Hook only**: The model wastes calls hitting the wall repeatedly, and the
  denial messages consume context. With narrative framing, most calls are
  correct on the first try.

Together: the narrative produces 80-90% correct tool selection, the hook
catches the remaining escapes and teaches the model to self-correct.

## Three Permission Layers (Not Two)

The guard hook is layer 2 of *three* layers an agent call must pass:

1. **Narrative framing** — shapes which tool the model reaches for
2. **PreToolUse hook (guard)** — default-deny; selectively allows safe tools/commands
3. **Claude Code permission system** — selectively denies anything not in `allowedTools`

The hook runs *in addition to* the permission system, not instead of it.
A hook that exits 0 (allow) only passes its own gate — the permission
system still checks whether the tool is pre-approved. In background agents,
there's no user to approve prompts, so unapproved tools auto-deny.

**Implication**: every command the guard allows must *also* have a matching
`allowedTools` pattern in settings.json (project or user level). The guard
opens the first door, the permission system opens the second. Both must
open for the call to execute.

### Permission Pattern Gotchas

`allowedTools` patterns use glob matching against the tool invocation.
Discovered limitations:

- **Heredocs break patterns.** `Bash(~/.ark/ark message *)` will NOT match
  `~/.ark/ark message new-request ... <<'BODY'` — the heredoc syntax
  prevents the glob from matching. This is why the scaffold→read→write
  pattern exists.
- **Tilde vs absolute path.** `Bash(~/.ark/ark *)` and
  `Bash(/home/user/.ark/ark *)` are separate patterns — you may need both
  depending on how the model expands the path.
- **Subcommands need separate patterns.** `Bash(~/.ark/ark *)` may not
  match `~/.ark/ark fetch --wrap`. Add explicit patterns:
  `Bash(~/.ark/ark fetch *)`, `Bash(~/.ark/ark message *)`, etc.
- **No per-agent scoping.** Permission patterns are session-wide. You
  can't grant Bash to one agent without granting it to all. The guard
  hook is what provides per-agent differentiation.

## Selective Path Access

The guard's default is deny-all. It can selectively allow non-Bash
tools on specific paths, expanding the agent's capabilities while
keeping the seal intact:

```bash
# Allow Read/Write on requests/ paths only
if [ "$TOOL" = Read ] || [ "$TOOL" = Write ]; then
    FPATH=$(echo "$INPUT" | jq -r '.tool_input.file_path')
    if echo "$FPATH" | grep -q '/requests/'; then
        exit 0
    fi
fi
```

This requires adding the tools to the agent's `tools:` frontmatter
(e.g., `tools: Bash, Read, Write`) and matching `allowedTools` patterns
in settings.json (e.g., `Write(requests/**)`, `Read(requests/**)`).

## Scaffold→Read→Write Pattern

When a CLI command creates a file with structured metadata (tag blocks,
frontmatter), split creation from content authoring:

1. **Scaffold**: CLI creates the file with correct metadata (tags, dates,
   IDs). No stdin, no heredoc — just flags.
2. **Read**: Agent reads the scaffold to capture the metadata block.
3. **Write**: Agent writes the full file — metadata (preserved exactly)
   plus body content.

This works around the heredoc permission problem while keeping the CLI
as the source of truth for metadata format. The model never touches
the tag syntax directly — it only adds content after the tags.

## What Doesn't Work

Tested and failed (12 iterations, documented):
- **Rules and prohibitions**: "NEVER use find" — ignored after 2 calls
- **Forbidden lists**: naming tools to avoid primes their use
- **World-model lies**: "files don't exist on disk" — disproved by `ls`
- **Consequences**: "wastes user's time" — too abstract for Haiku
- **Identity framing**: "betrayal of your duty" — activated research, not compliance
- **Challenge framing**: "solve this puzzle using only these tools" — marginal improvement
- **Cheating framing**: "using other tools is cheating" — Haiku cheated immediately
- **Tool count limits**: "3 commands maximum" — switched to non-ark tools after 3 ark calls

## Word Priming

A critical discovery: tool-name words in the prompt activate the corresponding
tools. "Find what is hidden" makes Haiku reach for `find`. "Read the contents"
activates `Read`. This is not metaphorical — replacing "find" with "uncover"
in the persona section measurably reduced `find` usage.

Principle: audit the entire agent prompt for words that are also tool names.
Replace them with synonyms everywhere except in actual CLI command references.

## Approval Comment Leakage

When a user approves/denies a subagent's tool call with a comment, the comment
is injected into the subagent's context, not the parent's. Mentioning a concept
in an approval comment can send the subagent on a research tangent about that
concept. Keep approval comments minimal or empty.

## Origin

Named for the triple meaning: Hermes (the agent) + hermetic seal (containment)
+ Hermetic tradition (accessing hidden knowledge through proper technique and
naming). Developed over 12 iterations constraining ark-hermes, a Haiku-powered
search agent for the ark digital zettelkasten.
