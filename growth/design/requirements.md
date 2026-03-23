# Requirements

## Feature: Phrase Display
**Source:** specs/growth.md — Phrase Display

- **R1:** `growth phrases N` prints N random committed phrases with phrase text, context, and observer
- **R2:** If fewer than N phrases exist, print all of them
- **R3:** If the oracle is empty, print a message saying so
- **R4:** Output includes framing preamble, phrases, inline tag list, and contemplation NPC note

## Feature: Mid-Session Contemplation
**Source:** specs/growth.md — Mid-Session Contemplation

- **R5:** `growth contemplate TAG` prints one random phrase matching TAG
- **R6:** If no phrases match TAG, print a message saying so
- **R7:** If TAG is omitted, print usage
- **R8:** Output wrapping is minimal — phrase and brief framing only

## Feature: Tag Listing
**Source:** specs/growth.md — Tag Listing

- **R9:** `growth tags` lists all unique tags across committed phrases with counts

## Feature: Phrase Consideration
**Source:** specs/growth.md — Phrase Consideration

- **R10:** `growth consider "PHRASE" "CONTEXT" OBSERVER` creates a pending candidate
- **R11:** If a consideration is already in progress, refuse and show the current candidate
- **R12:** `growth consideration` with no arguments prints the current candidate and phase
- **R13:** `growth consideration evaluate` advances to evaluation phase with deliberation prompts
- **R14:** `growth consideration tag TAG [TAG...]` adds tags to the current candidate
- **R15:** `growth commit` appends the candidate to PERSONA.jsonl and clears pending state
- **R16:** Commit requires at least one tag; error if none
- **R17:** `growth abort "REASON"` rejects the candidate with reason and clears pending state
- **R18:** Rejected phrases are stored for future reference as negative examples

## Feature: Storage
**Source:** specs/growth.md — Storage

- **R19:** Committed phrases stored in ~/.claude/personal/PERSONA.jsonl as type "phrase"
- **R20:** Rejected phrases stored in PERSONA.jsonl as type "rejected"
- **R21:** Pending candidate stored in ~/.claude/personal/considering.json
- **R22:** Each record includes a persona field for multi-persona filtering
- **R23:** Missing PERSONA.jsonl created on first write; treated as empty on read
- **R24:** Missing considering.json prints "no consideration in progress"
- **R25:** Malformed JSONL lines skipped with warning to stderr

## Feature: Persona Selection
**Source:** specs/growth.md — Persona Selection

- **R26:** Default persona is "daneel"
- **R27:** `--persona NAME` flag overrides default for all subcommands

## Feature: CLI
**Source:** specs/growth.md — Error Handling

- **R28:** No arguments prints usage summary
