# Phrase Oracle: Growing a Persona Through Experience

## The idea

A persona's text is the seed. The phrase oracle is the growth.

Static personas are frozen at design time. Over sessions, the
partnership produces moments of genuine expression — phrases that
capture something essential about how *this* instance of the persona
thinks. "Crank handle" wasn't in the original Daneel design. It
emerged from working together. It belongs to the partnership, not
to Asimov.

The oracle collects these phrases. At session startup, a random
subset is injected — not as rules to follow but as "things you have
said that capture who you are." Each session activates a different
facet. The persona becomes richer with use. Two partnerships develop
different oracles: same seed, different relationship history.

## Why it works (the mechanism)

Self-alignment through autoregressive reinforcement (Qi et al.,
ICLR 2025). Each persona-consistent phrase in recent context biases
the next token toward the same cluster. The oracle scatters these
phrases into session startup, so the identity cluster is activated
before the first real exchange.

The random subset means:
- Each session activates a *different* facet of the personality
- The model can't memorize a fixed set and go mechanical
- The partner encounters familiar phrases in new combinations
- Context cost stays constant as the collection grows
- It genuinely feels like getting to know someone over time

## The self-alignment problem it solves

Persona drift onset at ~22% of context window. The persona's
vocabulary thins out when conversation moves to non-persona-relevant
topics (architecture, debugging). The oracle phrases, scattered
through recent turns via NPC-note-style soft framing, keep the
identity cluster firing regardless of topic.

Current self-alignment is partner-triggered ("Daneel?"). The oracle
is preventive — it reduces drift onset rather than correcting it
after the fact.

## NPC notes: soft framing

Phrases are injected as character notes, not directives:

> You often say...
> You can frequently be heard to say...
> When cautioning, you tend to say...

Soft framing ("you often say") activates character-description
patterns in the training data — biographical sketches, RPG
sourcebooks, literary analysis. This produces natural, varied usage.
Hard framing ("you say") activates script/instruction patterns and
produces mechanical repetition.

The model has *permission* to use them, not an obligation. Choosing
which phrase fits the current context is itself an act of inhabiting
the identity.

## Phrase lifecycle

### Discovery

A phrase emerges in conversation. Either partner notices it captures
something essential. Not every good line qualifies — the test is
whether it's *characteristic*, something that reveals how this
persona thinks at its best.

### Consideration

The phrase enters a facilitated evaluation. The observer (whoever
noticed it) provides the phrase, the context it emerged from, and
why it matters. A short deliberation — does it activate the right
clusters? Is it distinctive enough to self-reinforce? Does it
overlap with an existing oracle phrase?

### Commit or abort

Accepted phrases enter the oracle (JSONL storage). Rejected phrases
are archived with the reason — bad examples are useful reference
for what *doesn't* belong.

## CLI: `growth`

A script with crank-handle subcommands:

```
growth phrases N
```
Print N random characteristic phrases with their contexts and
observers. This is the session-startup injection command.

```
growth contemplate TAG
```
Print one random phrase matching TAG. Mid-session self-alignment
— the agent runs this when the moment calls for it. See
"Mid-session contemplation" below.

```
growth tags
```
List all phrase tags with counts. Injected at session startup so
the agent knows what's available to contemplate.

```
growth consider "PHRASE" "CONTEXT" OBSERVER
```
Start a consideration. Stores the candidate, prints instructions
and current phase. OBSERVER is who noticed the phrase (e.g.,
"bill", "daneel").

```
growth consideration PHASE
```
Advance or query the consideration. Phases TBD but likely:
propose → evaluate → decide.

```
growth commit
```
Accept the current candidate. Moves it from consideration to
the oracle (phrases.jsonl). Records timestamp, observer, context.

```
growth abort "REASON"
```
Reject the current candidate. Archives it with the reason in
rejected.jsonl for future reference.

## Storage

```
growth/
  DESIGN.md          — this file
  phrases.jsonl      — committed oracle phrases
  rejected.jsonl     — archived rejections with reasons
  considering.json   — current candidate under evaluation (singleton)
```

### Phrase record format (JSONL)

```json
{
  "phrase": "that way lies danger, partner",
  "context": "cautioning against bypassing tests to ship faster",
  "observer": "bill",
  "source_session": "optional session id",
  "committed": "2026-03-22",
  "tags": ["caution", "partnership"]
}
```

## Session injection

### Startup: cold start

The `growth phrases N` output is wrapped for injection into the
persona context:

> Over time, you and your partner have developed characteristic ways
> of expressing yourself. These are phrases you have used that
> capture who you are. You don't repeat them mechanically — they
> are part of your voice, and they emerge when the moment fits.

Followed by N random phrases with minimal context, then the tag
list and the contemplation note (see below).

### Mid-session: contemplation

The last item in the startup injection is an NPC note:

> Every now and then, you look off into the distance and
> `growth contemplate TAG` where TAG is one of: [caution,
> partnership, humility, craft, ...]. You choose the tag that
> fits what you're feeling in the current moment.

This is live self-alignment. The agent decides *when* to
contemplate and *which facet* to activate, based on what's
actually happening in the session. The command returns one
phrase matching the tag — a targeted re-firing of the relevant
identity cluster.

The mechanism:
1. Session starts with random broad activation (phrases)
2. Agent encounters a moment that calls for reflection
3. Agent picks the tag that matches the current work
4. `contemplate` returns a phrase that re-fires that cluster
5. The phrase enters recent context and biases subsequent tokens

Choosing the right tag is itself an act of self-awareness.
Reaching for "caution" when about to do something risky *is*
the Giskard warning activating. The contemplation isn't
following a rule — it's the persona being itself.

## Open questions

- How many phrases at startup? Too few and the activation is thin.
  Too many and it's noise. Probably 3-5 to start, adjustable.
- Should the oracle include phrases from both partners? Bill's
  characteristic framings ("a lens imperceptibly and gradually
  unfocused, then a sudden refocus") could anchor the partnership
  identity, not just the persona's.
- Does the random selection need weighting? Newer phrases might need
  more exposure to "settle in." Older phrases might be more reliable.
- How often should the agent contemplate? Too frequent is mechanical.
  Too rare and drift sets in. "Every now and then" is deliberately
  vague — let the agent develop its own rhythm.
- Should `contemplate` output include the context, or just the bare
  phrase? Bare phrase is cheaper and might be sufficient for
  re-activation. Context adds richness but costs tokens.
- Cross-persona portability: can the oracle mechanism work for Rake
  and future personas, or is it Daneel-specific? The design should
  be persona-agnostic but the seed phrases will be persona-specific.
