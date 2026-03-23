# Renderer
**Requirements:** R1, R4, R5, R8, R9, R12, R13, R28

Formats all output. Keeps formatting logic out of CLI and Store.

## Knows
- (stateless — all input via parameters)

## Does
- renderSessionBlock(phrases, tags): full startup injection — preamble, phrases, tag list, NPC note
- renderContemplate(phrase): minimal mid-session output — phrase with brief framing
- renderTags(tags): compact tag list with counts
- renderCandidate(candidate): show current candidate with phase and instructions
- renderEvaluate(candidate): evaluation prompts for deliberation
- renderUsage: help text

## Collaborators
- (none — leaf component, called by CLI)
