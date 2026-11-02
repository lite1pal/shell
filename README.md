> [!IMPORTANT]
> It's a learning project, not a production-ready shell

# shell

A minimal Unix-style interactive shell written in Go.

It supports:

- Running external commands (`ls`, `pwd`, `git`, etc.)
- Built-in `cd` for changing the current working directory
- Built-in `exit` for leaving the shell
- A dynamic prompt with `user@host <cwd> >`

## Why This Project

This repository is a compact, readable shell implementation focused on learning and experimentation.  
The codebase is intentionally small and split into parser, executor, builtins, and runtime loop components.

## Requirements

- Go `1.25.5` (as declared in `go.mod`)
- A Unix-like environment (Linux/macOS recommended)

## Quick Start

```bash
git clone <your-repo-url>
cd shell
go run ./cmd/shell
```

You will see a prompt like:

```text
denis@machine /Users/denis/Work/Fun/go/shell >
```

## Usage

### External commands

```text
user@host /tmp > pwd
/tmp

user@host /tmp > ls -la
...
```

### Built-ins

- `cd <path>`: change directory
- `exit`: terminate the shell process

Examples:

```text
user@host /tmp > cd /usr/local
user@host /usr/local > exit
```

## Behavior Notes

- Empty input is ignored.
- Pressing `Ctrl+D` (EOF) exits the shell cleanly.
- If an external command fails, its error is printed to stderr.
- `cd` requires a path argument; otherwise it returns `cd: path required`.
- Argument parsing currently uses whitespace splitting (`strings.Fields`) and does not yet support quoting, escaping, pipes, or redirection.

## Project Structure

```text
.
├── cmd/shell/main.go             # Program entrypoint
└── internal/shell/
    ├── shell.go                  # REPL loop, prompt rendering, input read
    ├── parser.go                 # Input tokenization
    ├── executor.go               # Builtin dispatch + external command run
    └── builtins.go               # Builtin implementations (cd, exit)
```

## Roadmap

From `TODO.md`:

- [ ] Browse input history with the up/down keys

Potential next steps after history support:

- Quote-aware parsing (`"..."`, `'...'`, escapes)
- I/O redirection (`>`, `>>`, `<`)
- Pipelines (`|`)
- Environment variable expansion
- Tab completion

## Contributing

Contributions are welcome. Small, focused PRs are preferred:

1. Open an issue describing the change (bug, feature, or refactor).
2. Keep behavior changes covered by tests (when tests are added).
3. Keep implementation consistent with the current package layout.

## License

MIT license
