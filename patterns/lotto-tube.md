---
name: Lotto Tube
tags: agentic, event-loop, blocking, CLI, foundational
summary: A blocking CLI command that pops one work item from an application's event stream — the agent runs it in a loop to pull tasks
---

# Lotto Tube

A CLI command that blocks until the application has work for the
agent. The agent runs it in a loop: call, wait, receive, process,
repeat. The command is the tube — it sits there until a ball pops
up.

## The Shape

The application churns with activity — events, requests, state
changes. The lotto tube command connects the agent to that stream
one item at a time:

```
Agent                  Lotto Tube              Application
  |                        |                       |
  |-- run command -------->|                       |
  |                        |-- block (waiting) --->|
  |                        |                       | (churning...)
  |                        |                       | (event!)
  |                        |<-- JSON event --------|
  |<-- stdout + exit ------|                       |
  |                                                |
  | (process event)                                |
  |                                                |
  |-- run command again -->|                       |
  |                        |-- block (waiting) --->|
  ...
```

## Properties

- **Blocking, not polling.** The command blocks on the server
  side (long-poll, channel read, file watch). No sleep loops,
  no wasted cycles.
- **One *event* per invocation.** The command returns one event
  (or a batch), then exits. The agent decides what to do before
  asking for the next one. This is deliberate — it gives the
  agent a decision point between every item.
- **The agent's loop is event-to-event, never poll-to-poll.**
  Server-side long-poll timeouts, transient connection failures,
  port changes, and other infrastructure noise are absorbed
  *inside* the command (see Internal Resilience below). The agent
  only ever sees a real event on stdout, or a catastrophic-exit
  code (e.g. 52) meaning "the application is gone for good and
  you should stop." It never re-runs the command on an
  empty/timeout result, because empty results never reach it.
- **Structured payload on stdout.** The event is structured
  data. JSON is fine when the consumer is a script that
  programmatically dispatches; **markdown is preferred when an
  AI agent itself reads the output** (see Baby Food). The
  newsletter `wait` / `pull` commands emit a markdown work-item
  block — runId, kind, tab list as bullets — because the agent
  is the consumer. A `tee` to JSON for an automation hook is a
  separate concern.

## Prototype: `ark ui event`

```bash
# Agent runs this in a loop (background task)
Bash({cmd} event, run_in_background=true)

# Returns JSON: [{"app":"app-console","event":"select","name":"contacts"}]
#   when a real event arrives.
# Returns exit code 52 if the application has shut down for good
#   (agent stops looping).
# It does NOT return an empty array on timeout — internal long-poll
#   timeouts re-poll silently inside the command.
```

The agent reads the output, handles the event (which may involve
reading design docs, invoking skills, modifying code), then
restarts the listener.

## Internal Resilience

The tube presents a clean "one event per invocation" interface to
the agent, but internally it has a retry spine. The original
implementation (frictionless `.ui/mcp event`):

```bash
while true; do
    out="$(curl -s "http://127.0.0.1:$port/wait?timeout=120")"
    status=$?
    if [ -n "$out" ]; then
        # Server reconfigured or transient — re-read port, retry
        if echo "$out" | grep -q '"server_reconfigured"'; then
            sleep 1
            port=$(cat "$dir/mcp-port")
            continue
        fi
        echo "$out"   # real event — emit and exit
        break
    elif [ $status != 0 ]; then
        # Connection failed — server may have restarted
        sleep 1
        port=$(cat "$dir/mcp-port")
        continue
    fi
    # Empty body, status 0 = clean long-poll timeout → loop, don't exit
done
```

Note the empty/timeout case: the loop neither breaks nor errors —
it falls through and re-polls. Three classes of noise stay inside
the tube:

1. **Long-poll timeout** — server replied with no events; re-poll.
2. **Connection failure** — server restarting or briefly
   unreachable; sleep, re-read port, re-poll.
3. **Server reconfiguration** — port changed; re-read port,
   re-poll.

The agent never sees any of these. It only ever wakes for a real
event, or for a hard exit code that means "stop looping."

## Why Not a Persistent Connection?

WebSockets, gRPC streams, and stdin pipes all maintain a
persistent connection. The lotto tube deliberately breaks the
connection between events because:

1. **Decision point.** The agent processes event N before
   requesting event N+1. With a stream, events pile up
   and the agent falls behind.
2. **Crash recovery.** If the agent crashes mid-processing,
   the next invocation picks up fresh. No reconnection logic.
3. **Tool compatibility.** Claude Code's `Bash` tool runs
   commands and reads their output. A blocking command that
   returns JSON fits this interface perfectly. A persistent
   stream doesn't.

## Design Checklist

When building a lotto tube:

- [ ] Command blocks server-side (long-poll or channel), not
      client-side (sleep)
- [ ] Returns structured JSON on stdout — only when a real event
      arrives
- [ ] Server-side long-poll timeout absorbed *inside* the command
      (re-poll silently); the agent never sees an empty result
- [ ] Connection failures absorbed inside the command (sleep,
      re-read port, retry); the agent never sees them
- [ ] Distinct catastrophic-exit code (e.g. 52) for "application is
      shut down for good" so the agent stops looping
- [ ] Only one listener at a time (enforced server-side)
- [ ] Server-side long-poll timeout long enough to avoid churn,
      short enough to detect server death (~120s)

## Relationship to Sidecar Agent

The Lotto Tube is the delivery mechanism for the Sidecar Agent
pattern. The sidecar agent runs the tube in a loop, receives
work items, and delegates to subagents for processing.
