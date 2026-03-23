# Sequence: Consider (growth consider "PHRASE" "CONTEXT" OBSERVER)

```
CLI              Store            Renderer
 |                |                |
 |--loadCandidate->|               |
 |<--nil-----------|               |
 |                |                |
 |--saveCandidate->|               |
 |  (phase:proposed)|             |
 |                |                |
 |--renderCandidate(c)---------------->|
 |<--string--------------------------------|
 |--print-------->stdout          |
```

If candidate exists:
```
CLI              Store            Renderer
 |                |                |
 |--loadCandidate->|               |
 |<--candidate-----|               |
 |                |                |
 |--renderCandidate(c)---------------->|
 |  + "already in progress"       |
 |<--string--------------------------------|
 |--print-------->stderr          |
 |--exit(1)       |               |
```

# Sequence: Commit (growth commit)

```
CLI              Store            Renderer
 |                |                |
 |--loadCandidate->|               |
 |<--candidate-----|               |
 |                |                |
 |--check tags--->|                |
 |  (len > 0?)    |                |
 |                |                |
 |--commitPhrase-->|               |
 |                |--append JSONL->|
 |                |--delete json-->|
 |                |                |
 |--print "committed"-->stdout    |
```

# Sequence: Abort (growth abort "REASON")

```
CLI              Store            Renderer
 |                |                |
 |--loadCandidate->|               |
 |<--candidate-----|               |
 |                |                |
 |--rejectPhrase-->|               |
 |  (reason)      |                |
 |                |--append JSONL->|
 |                |--delete json-->|
 |                |                |
 |--print "rejected"-->stdout     |
```
