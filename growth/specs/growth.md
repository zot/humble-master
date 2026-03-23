# Growth: Phrase Oracle CLI

A Go CLI tool that manages a growing collection of characteristic
phrases for AI personas. Phrases are discovered in conversation,
curated through a consideration process, and injected at session
startup to sustain persona alignment through autoregressive
self-reinforcement.

Installed at `~/.claude/bin/growth`. Storage at
`~/.claude/personal/PERSONA.jsonl`. The tool has no dependency on
ark or any external service.

Language: Go. Environment: CLI, single-user, local filesystem.

## Phrase Display

`growth phrases N` prints N random phrases from the oracle. Output
includes the phrase text, the context it emerged from, and who
observed it. This is the session-startup command — its output is
designed for injection into a persona's context.

If fewer than N phrases exist, print all of them. If the oracle is
empty, print a message saying so.

Output wrapping: the output is wrapped in a framing block that
tells the persona these are "things you have said that capture who
you are" and that they emerge naturally, not mechanically. After the
phrases, print the tag list and the contemplation NPC note (see
Session Injection below).

## Mid-Session Contemplation

`growth contemplate TAG` prints one random phrase matching TAG.
This is agent-initiated self-alignment — the agent runs it when the
moment calls for reflection.

If no phrases match TAG, print a message saying so. If TAG is
omitted, print usage.

Output wrapping: the output is wrapped minimally — just the phrase
and enough framing that the agent recognizes it as something it has
said before. No tag list, no NPC note (those are startup-only).

## Tag Listing

`growth tags` lists all unique tags across all committed phrases,
with counts. Output is compact, suitable for injection into the
contemplation NPC note.

## Session Injection

The `growth phrases N` command produces a complete session-startup
block:

1. Framing text (the "things you have said" preamble)
2. N random phrases with context and observer
3. Tag list (output of `growth tags`, inline)
4. The contemplation NPC note:

   > Every now and then, you look off into the distance and
   > `growth contemplate TAG` where TAG is one of: [tag list].
   > You choose the tag that fits what you're feeling in the
   > current moment.

This entire block is what gets injected into the persona context
at session startup.

## Phrase Consideration

Phrases enter the oracle through a facilitated process, not bulk
import. One candidate at a time.

### Starting a consideration

`growth consider "PHRASE" "CONTEXT" OBSERVER` creates a candidate
for evaluation. OBSERVER is who noticed the phrase (e.g., "bill",
"daneel"). The command stores the candidate in a pending state and
prints the candidate back with instructions for the next step.

If a consideration is already in progress, print the current
candidate and refuse to start a new one. The user must commit or
abort first.

### Advancing a consideration

`growth consideration` with no arguments prints the current
candidate and its phase.

`growth consideration evaluate` advances to evaluation phase.
Prints the candidate and prompts for deliberation: does it activate
the right clusters? Is it distinctive enough to self-reinforce?
Does it overlap with existing phrases?

`growth consideration tag TAG [TAG...]` adds tags to the current
candidate. Tags can be added at any phase.

### Committing

`growth commit` accepts the current candidate. Appends it to
PERSONA.jsonl with timestamp, observer, context, tags, and the
persona name. Clears the pending state. Prints confirmation.

The candidate must have at least one tag before committing. If no
tags, print an error asking for tags.

### Aborting

`growth abort "REASON"` rejects the current candidate. Appends it
to a rejected section in PERSONA.jsonl (or a separate
rejected.jsonl) with the reason. Clears the pending state.
Prints confirmation.

Rejected phrases serve as reference for what doesn't belong —
negative examples for future curation.

## Storage

All data lives in `~/.claude/personal/PERSONA.jsonl`. One JSON
object per line.

### Committed phrase record

```json
{
  "type": "phrase",
  "persona": "daneel",
  "phrase": "that way lies danger, partner",
  "context": "cautioning against bypassing tests to ship faster",
  "observer": "bill",
  "tags": ["caution", "partnership"],
  "committed": "2026-03-22"
}
```

### Rejected phrase record

```json
{
  "type": "rejected",
  "persona": "daneel",
  "phrase": "I shall endeavor to comply",
  "context": "responding to a task assignment",
  "observer": "daneel",
  "tags": ["compliance"],
  "reason": "too formal, sounds like a butler not a partner",
  "rejected": "2026-03-22"
}
```

### Pending candidate

Stored in `~/.claude/personal/considering.json` (not JSONL — a
single JSON object, overwritten each time). Contains the candidate
fields plus a `phase` field ("proposed" or "evaluate").

## Persona Selection

The tool needs to know which persona to filter for. By default it
uses "daneel". A `--persona NAME` flag overrides this. The flag
applies to all subcommands.

Future: if multiple personas share the same PERSONA.jsonl, the
persona field in each record distinguishes them.

## Error Handling

- Missing PERSONA.jsonl: create it on first write. On read, treat
  as empty.
- Missing considering.json: "no consideration in progress."
- Malformed JSONL lines: skip with a warning to stderr.
- No arguments: print usage summary.
