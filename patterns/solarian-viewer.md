---
name: Solarian Viewer
tags: agentic, architecture, context-isolation, skill, agent, CLI
summary: Skill (viewer) → Agent (robot) → CLI (controls). Three-layer runtime where each layer does what it's best at — deciding, operating, enforcing.
---

# Solarian Viewer

Named for the Solarian viewing rooms in Asimov: the viewer says what
they want to see. The robot operates the controls. The controls
produce the image.

## The three layers

**Skill (the viewer)** — the UX layer, loaded into the caller's
context. Knows the user's operational context (current project),
routes requests, speaks plain English. Tiny context footprint. The
viewer's power is in *deciding what to look at*, not operating
machinery.

**Agent (the robot)** — the expertise layer, runs as a subagent
(typically Haiku). Knows every flag, convention, and search strategy.
Curates results, expands queries, reports honest misses. Its full
context — CLI reference, conventions, domain knowledge — lives and
dies in the subagent. The caller never sees it.

**CLI (the controls)** — the mechanism layer. Enforces format,
executes operations, returns structured output. Format intelligence
lives here, not in any model. The crank-handle principle: the model
says *what*, the command handles *how*.

## Why three and not two

Each layer exists because removing it forces one of the others to
do something it's bad at:

- Without the agent: the skill must load CLI knowledge into the
  caller (expensive) or have the caller interpret raw output
  (unreliable)
- Without the skill: the user must know to spawn the agent and
  scope queries correctly
- Without the CLI: the agent must carry format rules in context
  (token waste, still error-prone) or get format wrong and retry

## Operational context flows down

The skill knows the user's project. It passes this to the agent as
part of the prompt. The agent uses it to scope operations. The CLI
doesn't need it — it operates on everything, the agent filters.

"Open messages" → skill scopes to "open messages targeting ark" →
agent searches and curates → CLI returns structured results.

Same principle as `git status` implying the current repo.

## Properties

- **Context isolation** — each layer's internal knowledge is private.
  The caller never sees CLI flags; the user never sees agent reasoning.
  Closure-actor principle at the architecture level.
- **Crank-handle at the bottom** — format intelligence in the CLI
  eliminates retries from format confusion.
- **Cheap expertise** — the agent runs on Haiku. Judgment and
  curation at SMS-scale cost.
- **Non-expert accessible** — the user doesn't need to be a domain
  expert. The skill carries the operational context they'd otherwise
  need to learn.

## First implementation

The ark system: `/ark` skill → `ark-hermes` agent → `ark` CLI.
Design notes in `~/work/ark/librarian/`.
