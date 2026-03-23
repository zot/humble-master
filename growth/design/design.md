# Growth CLI — Design

## Intent

A single-binary Go CLI that manages a phrase oracle for AI persona
self-alignment. Minimal components — this is a small tool that reads
and writes JSONL/JSON files and formats output.

## Artifacts

### CRC Cards
- [ ] crc-CLI.md → `src/main.go`
- [ ] crc-Store.md → `src/store.go`
- [ ] crc-Renderer.md → `src/render.go`

### Sequences
- [ ] seq-phrases.md → `src/main.go`, `src/store.go`, `src/render.go`
- [ ] seq-consider.md → `src/main.go`, `src/store.go`

## Cross-cutting Concerns

### File Paths
Storage dir is `~/.claude/personal/`. The tool resolves `~` to
`$HOME`. Creates the directory if it doesn't exist on first write.

### Persona Filtering
All read operations filter JSONL records by the `persona` field
matching the active persona (default "daneel", override with
`--persona`).

## Gaps

(populated after implementation)
