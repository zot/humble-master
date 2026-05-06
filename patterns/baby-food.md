---
name: Baby Food
tags: agentic, format, weak-models, markdown, foundational
summary: Agent-facing inputs are markdown views of structured data. The CLI does the chewing — the model never parses JSON, never balances braces, never looks up keys.
---

# Baby Food

The agent reads markdown. Always. If the underlying data is JSON,
YAML, a SQLite row, or a wire-format blob, the CLI renders a
markdown view of it for the agent to consume. The kitchen mashes
the peas; the agent eats.

## The asymmetry

Models speak markdown natively. Headings, bullets, labeled lines —
that's the shape of half their training corpus. They produce
markdown without thinking and consume it without strain.

JSON they have to *chew*. Find the matching brace. Track which key
in the schema is the right one for this slot. Escape the strings.
Compose the comma-separation. None of that is hard for a model,
but every gram of attention spent on syntax is a gram not spent on
the task. Strong models pay the price quietly; weak models fumble
visibly.

The same applies on the way in. Telling an agent "Read
`cache/foo.json`" hands it raw food. The model parses to extract
the field it needs and burns tokens summarizing what it found.
Worse, the prompt that taught it to write that file probably
embedded the schema as inline JSON — now there's JSON in the
prompt *and* JSON in the file, and the agent has to reconcile
them.

## The fix

Every agent-facing read is markdown:

- **A view of structured data is rendered as markdown** in the
  phase prompt or in a `.md` file the prompt points to. Tab list
  becomes a bulleted list. Cluster output becomes
  `## Cluster Title` / `**Theme:** …` / `### Source: <url>` with
  labeled bullets for fields. The agent sees the data the way a
  human reading the file would see it.
- **Files the agent reads end in `.md`** when the format is
  agent-facing. JSON files exist on disk for the renderer / the
  validator / the next CLI hop, but the agent never opens them.
- **Schemas don't appear in prompts.** The shape of the agent's
  output is shown by example (a markdown stencil), not by pasted
  JSON Schema. The schema is the CLI's concern; the example is
  the agent's.

Disk format stays whatever fits the consumer. The renderer wants
JSON; the SQLite query wants rows; the WebSocket wants a binary
blob. Whatever the on-disk shape, the *agent-facing wire format*
is markdown.

## Properties

- **The model's time is the budget.** Every chewing turn is a
  recovery turn that didn't need to happen. Pre-chewed food
  is faster, cheaper, and surfaces the actual content sooner.
- **Human bonus is real, not incidental.** `cat cache/foo.md`
  is readable; `cat cache/foo.json | jq` requires a manual.
  Diagnostics improve for free.
- **Symmetric on the output side.** When the agent has to produce
  structured data, it writes a markdown stencil — see Stencil for
  that direction.
- **The asymmetry is not weakness-specific.** Opus benefits as
  much as Haiku does. Opus pays the JSON-chewing cost more
  quietly, but it pays it.

## Instances

- **Newsletter discover phase.** The prompt embeds the tab list
  as a markdown bullet list. The agent writes `cache/clusters.md`
  in a stencilled markdown shape; the CLI parses it into
  `cache/clusters.json` for the renderer. Net: the agent writes
  and reads only markdown; JSON is CLI-internal.
- **Newsletter research phase.** The prompt embeds the discovery
  clusters from `cache/clusters.md` directly — the research
  agent never opens `cache/clusters.json`.
- **`newsletter health`.** The CLI reads several JSON state
  files (`cache/connection.json`, server status, JSONL
  metadata) and renders one markdown report the agent reads at
  the top of each cycle. The agent never sees the underlying JSON.

## Failure mode

The diagnostic that says "we got this wrong" is the agent
inspecting its own files:

```
node -e "const d = JSON.parse(require('fs').readFileSync('cache/clusters.json','utf8'));
         console.log('clusters:', d.clusters.length); ..."
```

If the agent has reason to run that command, the CLI failed to
plate the data. Either the upstream JSON is leaking into the
agent's prompt, or the file the agent was told to read is JSON
when it should be a markdown view.

## Relationship to other patterns

- **Stencil** — the symmetric output-side pattern. Baby Food is
  what the agent reads; Stencil is what the agent writes. Both
  externalize format awareness into the CLI.
- **Crank Handle** — every crank-handle prompt is markdown by
  default. Baby Food is the rule that keeps it that way as the
  prompt grows.
- **Visible Substrate** — both render the underlying data in the
  consumer's medium. Visible Substrate renders for the human's
  eye; Baby Food renders for the model's tongue.
- **Fumble Log** — when the symmetric (Stencil) side fails, the
  Fumble Log records it. Baby Food failures are quieter — they
  show up as token bloat, not parse errors.

## Origin

Named with Bill, 2026-05-06, during the inside-out newsletter
migration. Symptom: a session was running
`node -e ...JSON.parse(...clusters.json)` to inspect its own
output before advancing. The CLI had asked the agent to read JSON.
The fix made every agent-facing artifact markdown; the symptom
went away. A little silliness in the name was deliberate — AI
takes it fine, and the metaphor is exactly right.
