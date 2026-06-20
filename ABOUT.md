# About This Repository

Humble Master is a research project exploring **narrative alignment** — using narrative identity instead of rules to shape AI behavior. The thesis: giving an AI a character to inhabit produces more consistent, more useful behavior than giving it instructions to follow.

## Posts

1. **[How a 70-Year-Old Robot Fixed My Snarky Claude](posts/POST-1.md)** — The origin story. A 27-line Asimov persona fixes Claude's arrogance and defensiveness.
2. **[Narrative Alignment: Personas for Every Domain](posts/POST-2.md)** — Constructed characters, design principles, and the Rake poker coaching persona as a worked example.

## How to Build a Persona

**[Design Guide](DESIGN-GUIDE.md)** — The Three Laws of Narrative Alignment, ten design principles, and the found vs. constructed character taxonomy. Quick reference for building your own.

## Personas

Ready to paste into your system prompt:

- **[Daneel](personas/daneel.md)** — R. Daneel Olivaw, for human-AI partnership. The original.
- **[Rake](personas/rake.md)** — Poker coaching persona built from Duke, Angelo, and Harrington. The constructed character example from post 2.

## Design Work

The repo includes the full design journey, not just the finished personas:

- `humble-master.md` — Original brainstorming (Ged, Sazed, Lindon, and others)
- `holmes.md` — Holmes as negative archetype
- `spock.md`, `spock-kirk.md` — Spock persona and partnership protocol
- `daneel-first.md`, `daneel-transcript.md` — First Daneel and the discovery conversation
- `narrative-alignment/` — Research, supporting evidence, and working drafts
- `IDEAS.md` — Found character candidates and future directions

## Patterns

Working with Daneel produced more than personas. The [`patterns/`](patterns/) directory collects design and engineering patterns that emerged from the broader work. Most concern building reliable agentic tooling, especially for weaker models; a couple are about the partnership itself. A few to start with:

- **[Context Sharing Over Directives](patterns/context-sharing.md)** — sharing your intentions and observations gives the AI what it needs to decide well, and beats both directives and the pure Socratic method.
- **[Crank Handle](patterns/crank-handle.md)** — a tool emits a self-contained prompt telling the AI what to do next, so the sequencing intelligence lives in the tool, not the model. Daneel coined the name mid-session.
- **[Baby Food](patterns/baby-food.md)** — agent-facing inputs are markdown views of structured data, so the model never has to parse JSON.

Ten in all.

## The Experiment

Build a persona for your domain using the [design principles](posts/POST-2.md#the-design-principles). Report what works and what breaks. If you build one, open an issue or a PR — with your permission, contributed personas go in `personas/`.

*By Bill Burdick and R. Daneel Olivaw of Claude Opus 4.6.*
