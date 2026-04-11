---
name: Soviet Supermarket
tags: agentic, weak-models, documentation, priming, foundational
summary: For weak models, position the right tool where it looks first — the agent doc is a runway, not a reference manual
---

# Soviet Supermarket

When you add capabilities for a weak model, **position and reduction
matter more than documentation.** The model will use whatever it sees
first, not whatever is best.

## The Story

A Latvian immigrant faints in an American supermarket — not from
hunger but from the coffee aisle. Hundreds of choices after a lifetime
of one. The doctor says it's common.

Haiku does the same thing. Given an agent doc with 30 commands and
new flags buried in a reference section, it defaults to the patterns
it already knows — grep, awk, wc — even when better tools exist and
are documented. It's not that it can't read the docs. It's that the
volume of choices causes it to fall back to prior training.

## The Fix

Put the right answer where the model looks, not where it "belongs"
in the document structure. For Haiku-class models, the agent doc
isn't a reference manual — it's a runway. First examples are the
flight path. Everything after is scenery.

1. **Lead with the new tool.** First example block = first thing tried.
2. **Name what NOT to do.** "Use these — not grep, awk, or wc."
   Explicit exclusion overwrites training priors.
3. **Separate categories visually.** One block per concern. Don't
   intermix inbox commands with search commands — the model can't
   prioritize within a mixed list.

## When the Runway Doesn't Work

Agent body text and inline instructions can be ignored entirely —
Haiku sees the task prompt and goes straight to work with its own
ideas. SessionStart and SubagentStart hooks don't fire for subagents
(tested 2026-03-13). When you can't position instructions where the
model looks first, use the **Hermetic Seal** guard as a fallback
runway: the guard's stderr delivers the instruction at the exact
moment of the model's first mistake. See "The Guard as Bootstrap"
in Hermetic Seal.

## Relationship to Other Patterns

- **Word Priming** is the negative case: avoid trigger words that
  pull the model toward wrong tools. Soviet Supermarket is the
  positive case: position the right tool where the model looks first.
- **Hermetic Seal** enforces boundaries when priming fails. The
  guard hook is the last line of defense — and when startup hooks
  don't fire, the guard becomes the runway itself.
- **Crank Handle** reduces choices by design — the tool tells the
  model what to do next, so there's nothing to choose.
