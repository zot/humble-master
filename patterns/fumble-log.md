---
name: Fumble Log
tags: agentic, diagnostics, weak-models, stencil, foundational
summary: Append-only log of stencil-parse failures with the input that caused them. Silent ride-along. Without it, format-tightening is guesswork; with it, the agent's recurring fumbles become evidence.
---

# Fumble Log

When the CLI rejects an agent-authored stencilled artifact, the
agent retries — usually correctly. The recovery turn costs tokens
and leaves no trace. Multiply by every run, every agent, every
phase: a meaningful budget burns invisibly.

The Fumble Log makes those rejections visible. Each parse failure
appends a record: timestamp, run id, phase, the input that was
rejected, the parser's errors. Days later, you grep the log,
notice that 30% of failures are the same off-by-one bullet shape,
and either tighten the prompt or relax the parser. The format
evolves from evidence, not guesswork.

## The technique

```
cache/.cc/parse-errors.log                   ← append-only
=== 2026-05-06T14:22:01.231Z run=run-x phase=DISCOVER ===
Errors:
  - line 14: unknown source field "Tldr" — expected one of: Page title, Published, ...
  - line 22: source heading found before any cluster (## ...) heading
--- input ---
<the markdown the agent wrote>
--- end ---
```

Implementation is small: when parse / validation fails, append a
record with the input. Continue with the existing retry path —
the log is parallel to the agent's experience, not in it.

```javascript
function logParseFailure(cacheDir, phase, runId, input, errors) {
  try {
    const logPath = path.join(cacheDir, '.cc', 'parse-errors.log');
    fs.mkdirSync(path.dirname(logPath), { recursive: true });
    fs.appendFileSync(logPath, formatEntry(phase, runId, input, errors));
  } catch {
    // Logging is best-effort; a failure here must not break the pipeline.
  }
}
```

## Properties

- **Silent ride-along.** The agent only sees the retry prompt
  with the parse errors. The log is for human review later. No
  prompt bloat, no behavior change.
- **Best-effort.** Logging failures must never break the
  pipeline. Disk full, permission denied, anything — the CLI
  swallows the error and continues.
- **Append-only.** Every entry is evidence. Nothing is edited or
  truncated. Log rotation is a rare manual step, not automation.
- **Co-located with the artifact, not centralized.** The newsletter
  log lives in `cache/.cc/`, alongside the artifacts that produced
  it. Diagnostic context stays with the run it diagnoses.
- **Includes the input verbatim.** Without the input, the errors
  are unactionable. With it, you can re-run the parser locally,
  reproduce the failure, and try a fix.

## What you do with the log

Periodically:

1. `cat cache/.cc/parse-errors.log | grep "^  -"` — what errors
   recur? An error that fires in 60% of failures is a prompt
   problem; tighten the example.
2. Pick a recurring failure, copy its input, paste into a unit
   test. Now the parser has a regression test grounded in a real
   failure shape.
3. If the same error keeps appearing despite prompt tightening,
   the *parser* is too strict. Relax it (e.g. accept both
   `**Published:**` and `**Published date:**`).
4. Tag fixed failures with a date in the prompt's changelog so
   you can A/B the rate before vs. after.

The point isn't dashboards. It's *evidence over assumption* when
deciding how to evolve the format-agent contract.

## When to use it

Anywhere a CLI parses agent-authored structured input:

- Stencil-pattern files (newsletter `clusters.md`, ark message
  tag blocks, mini-spec design docs)
- CLI subcommands that take stdin from agents
- Validation steps in any agentic pipeline where the agent
  produces structured artifacts

The threshold is low — even a few entries a week pay for
themselves in one prompt-tightening pass.

## When not to use it

- The CLI never rejects anything (the format is permissive). No
  rejections, nothing to log.
- The agent's input is unstructured prose with no parse step. The
  whole pipeline is downstream of the agent; there's no parser
  rejection to observe.

## Relationship to other patterns

- **Stencil** — Fumble Log is Stencil's longitudinal diagnostic.
  Stencil enforces the shape; Fumble Log records every miss.
  Without the log, Stencil's correctness over time depends on
  the maintainer's intuition.
- **Baby Food** — the asymmetric pattern on the read side. Baby
  Food failures are quieter (token bloat, not parse errors), so
  they're harder to log; Fumble Log is for the loud kind.
- **Fumble Onboarding** — different pattern, similar instinct:
  the fumble is data. Fumble Onboarding turns the first fumble
  into a teaching moment; Fumble Log turns the Nth fumble into
  format evidence.
- **Visible Substrate** — both surface invisible-by-default
  state. Visible Substrate makes data structures visible to the
  eye; Fumble Log makes recovery turns visible to the auditor.

## Origin

Proposed by Bill, 2026-05-06, during the inside-out newsletter
migration. Symptom: he noted that mini-spec's CLI occasionally
chokes on agent-authored design docs, and the agent quietly
re-edits to fix it. The fumble was real; the cost was real;
nobody could see either. The log is the seeing.
