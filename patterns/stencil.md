---
name: Stencil
tags: crank-handle, format-enforcement, weak-models, CLI, foundational
summary: CLI commands read/write rigidly formatted files so models never touch the format directly. The shape is fixed; the model only provides what goes inside.
---

# Stencil

A specialization of the crank-handle pattern for structured file I/O.
The file has a rigid format with slots. The CLI stamps values into the
slots. The model provides content; the command provides structure.

## The problem

Models write natural-looking output. A tag block needs `@status: done`
on its own line at the top of the file with no blank lines between
tags. The model writes `## Status: done` in the body because that's
what markdown looks like. Opus does this. Haiku does it more. The
format is simple enough that a human would never get it wrong, but
models don't see format — they see likely next tokens.

You cannot prompt your way out of this. The model doesn't know it's
making a mess.

## The fix

The model never touches the file directly. CLI commands are the only
interface:

- **Read through the stencil:** `ark message get-tags FILE` extracts
  values from the format. The model sees `status\tdone`, not the raw
  tag block.
- **Write through the stencil:** `ark message set-tags FILE status done`
  stamps the value into the correct slot. The model says what, the
  command handles where and how.
- **Create through the stencil:** `ark message new-request --from X --to Y`
  generates the entire file with correct structure. The model provides
  parameters, never a template.

## Properties

- **Format is invisible to the model.** It never sees tag block syntax,
  never decides where tags go, never chooses between `@status:` and
  `## Status:`.
- **Validation is a crank-handle.** `ark message check` outputs fix
  commands if the format is wrong — the model executes them without
  understanding the format rules.
- **Scales down to any model.** The weaker the model, the more it
  needs the stencil. But even strong models benefit — Opus needed
  stencils for mini-spec design doc format.

## Instances

- **ark message** — tag blocks in `requests/` files. `set-tags`,
  `get-tags`, `new-request`, `new-response`, `check`.
- **mini-spec** — structured design documents. The Go program parses
  the format, checks off requirements, queries by status. Opus
  provides intent, the program manages the document.

## Relationship to crank-handle

The crank-handle says: "you know what you want, I'll handle the
sequencing." The stencil says: "you don't even know you're making
a mess, so I won't let you touch it." Both externalize intelligence
into the tool. The crank-handle externalizes *sequencing*. The
stencil externalizes *format awareness*.

## When to use

Any time a model needs to read or write a file with rigid formatting:
- Tag blocks, front matter, structured headers
- Config files with positional rules
- Documents that other tools parse (the format is a contract)

If the format is simple enough that "just tell it the rules" seems
sufficient — that's exactly when you need a stencil. The model will
get it right 90% of the time and corrupt your data the other 10%.
