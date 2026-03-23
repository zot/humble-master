# Store
**Requirements:** R1, R2, R3, R5, R6, R9, R10, R11, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25

Reads and writes PERSONA.jsonl and considering.json. All filesystem
I/O is here.

## Knows
- baseDir: path to ~/.claude/personal/
- persona: active persona name for filtering

## Does
- loadPhrases: read PERSONA.jsonl, filter by persona and type "phrase", return slice
- randomPhrases(n): return n random phrases (or all if fewer exist)
- randomByTag(tag): return one random phrase matching tag
- allTags: collect unique tags with counts across committed phrases
- loadCandidate: read considering.json, return candidate or nil
- saveCandidate: write considering.json
- clearCandidate: delete considering.json
- commitPhrase: append phrase record to PERSONA.jsonl, clear candidate
- rejectPhrase: append rejected record to PERSONA.jsonl, clear candidate
- ensureDir: create baseDir if needed

## Collaborators
- (none — leaf component)
