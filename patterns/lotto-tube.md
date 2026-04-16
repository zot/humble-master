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
- **One item per invocation.** The command returns one event (or
  a batch), then exits. The agent decides what to do before
  asking for the next one. This is deliberate — it gives the
  agent a decision point between every item.
- **Timeout exits cleanly.** On timeout (no events within the
  window), the command exits 0 with an empty result. The agent
  restarts it. No error, no retry logic.
- **JSON on stdout.** The event is structured data. The agent
  parses it and dispatches.

## Prototype: `ark ui event`

```bash
# Agent runs this in a loop (background task)
Bash({cmd} event, run_in_background=true)

# Returns JSON: [{"app":"app-console","event":"select","name":"contacts"}]
# Or empty array [] on timeout
# Or exit code 52 if server restarted
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
done
```

Connection failures, port changes, and transient server states
loop silently inside the tube. The agent never sees them — it
just gets a clean event or a timeout. The tube absorbs
infrastructure noise.

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

- [ ] Command blocks server-side (long-poll or channel), not client-side (sleep)
- [ ] Returns structured JSON on stdout
- [ ] Exits 0 on timeout with empty result
- [ ] Distinct exit code for "server restarted" (agent should restart loop)
- [ ] Only one listener at a time (enforced server-side)
- [ ] Timeout long enough to avoid churn, short enough to detect server death (~120s)

## Relationship to Sidecar Agent

The Lotto Tube is the delivery mechanism for the Sidecar Agent
pattern. The sidecar agent runs the tube in a loop, receives
work items, and delegates to subagents for processing.
