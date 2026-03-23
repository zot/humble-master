# CLI
**Requirements:** R1, R4, R5, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R26, R27, R28

Parses arguments, dispatches to subcommands, handles --persona flag.

## Knows
- args: command-line arguments
- persona: active persona name (from flag or default)

## Does
- parseArgs: extract --persona flag, identify subcommand and args
- dispatch: route to the appropriate handler
- printUsage: show help text

## Collaborators
- Store: all data access
- Renderer: all output formatting
