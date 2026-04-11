---
name: Crank Handle
tags: agentic, state-machine, weak-models, externalized-intelligence
summary: A tool outputs a self-contained prompt telling the AI what to do next. Sequencing intelligence lives in the tool, not the model. Designed for weak models that can execute but not plan.
---

@learned: Daneel (the persona) coined the term "crank handle" during
a design session with Bill.
@connection: minispec phase commands = crank handle pattern. Each
`minispec phase <step>` validates one workflow phase and tells the
AI what to fix or do next. The sequencing intelligence lives in the
Go tool, not the model.

# Crank Handle

A tool outputs a self-contained prompt. The AI reads it, follows it,
runs a command, feeds the result back. The tool outputs the next
prompt. Repeat until done.

## The image

A hand crank on a machine. You turn it, something comes out. You
feed that back in, turn again. The intelligence is in the machine,
not the operator.

## Why it exists

Strong models (Opus) can infer multi-step sequences from a
description. Weak models (Haiku, Sonnet) need each step handed to
them as a complete, unambiguous instruction. The crank-handle
externalizes sequencing intelligence so the model only needs to
execute, not plan.

This matters because the economics of agentic systems favor cheap
models for routine work. If the tool carries the plan, a $0.25/M
model does the same job as a $15/M model.

## Shape

```
Tool output → AI reads → AI executes → Tool output → ...
```

Each output is:
- **Self-contained** — everything the model needs to act, no
  external context required
- **Unambiguous** — one clear action, not a menu of options
- **Terminal or continuing** — either "do this and you're done"
  or "do this and run me again"

## When to use it

- A CLI tool needs an AI to do something the tool can't
  (edit a file intelligently, make a judgment call)
- The sequence has dependencies — step 2 depends on step 1's
  result
- The consumer might be a weak model
- You want the same tool to work across model tiers without
  tier-specific prompting

## When not to use it

- The model is always strong and the sequence is simple —
  just describe what you want
- There are no dependencies between steps — give all
  instructions at once
- The tool can do everything itself without AI help

## Examples

### ark install

```
$ ark install
```

If `.ui/` is missing:
```
Install Frictionless first. Run the following:
1. Read the README at github/zot/frictionless
2. Follow its install instructions
3. Then re-run: ~/.ark/ark install
```

If `.ui/` is present:
```
Skill and agent installed. Add the following line to the top
of this project's CLAUDE.md:

    load /ark

This enables ark's long-term memory for every session in this
project.
```

Each output is a complete instruction. The model doesn't need to
know what ark is, what Frictionless is, or why CLAUDE.md matters.
It just follows the prompt.

### ark decide (planned)

State machine for weaker models during routine work:
```
$ ark decide check
→ "Run `ark search --chunks @decision:` and check if any
   decisions were made in this session that aren't tagged yet.
   If you find untagged decisions, tag them. Then run
   `~/.ark/ark decide next`."
```

The phase machine (scan, tag, note, review, idle) lives in ark.
The model is the executor. Sequencing intelligence is fully
externalized.

## Relationship to other patterns

- **State machine** — the crank-handle IS a state machine, but
  the transitions are expressed as natural language prompts, not
  code. The "state" is whatever the tool tracks internally.
- **Continuation passing** — similar shape, but continuations are
  code. Crank-handle continuations are prompts for an AI.
- **Wizard/step-by-step** — wizards are for humans. Crank-handles
  are for models. The output is optimized for instruction-following,
  not comprehension.

## The opposite: Magic Eight Ball

Without externalized structure, an AI is a magic eight ball. You
shake it, an answer floats up from whatever's in context. Sometimes
profound, sometimes "ask again later." No memory, no recall, no
persistence. The answer disappears when you close the session.

The crank-handle gives the model *structure*. The eight ball gives
it *vibes*.
