# Sequence: Session Startup (growth phrases N)

```
CLI              Store            Renderer
 |                |                |
 |--parseArgs---->|                |
 |  (n=5)         |                |
 |                |                |
 |--loadPhrases-->|                |
 |                |--read JSONL--->|
 |                |--filter------->|
 |<--[]Phrase-----|                |
 |                |                |
 |--randomPhrases(n)------------->|
 |<--[]Phrase-----|                |
 |                |                |
 |--allTags------>|                |
 |<--map[tag]int--|                |
 |                |                |
 |--renderSessionBlock(phrases, tags)-->|
 |<--string------------------------------|
 |                |                |
 |--print-------->stdout          |
```

# Sequence: Contemplation (growth contemplate TAG)

```
CLI              Store            Renderer
 |                |                |
 |--parseArgs---->|                |
 |  (tag="caution")|              |
 |                |                |
 |--randomByTag-->|                |
 |                |--read JSONL--->|
 |                |--filter tag--->|
 |<--*Phrase------|                |
 |                |                |
 |--renderContemplate(phrase)----------->|
 |<--string------------------------------|
 |                |                |
 |--print-------->stdout          |
```
