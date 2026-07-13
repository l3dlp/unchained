# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a monorepo containing standalone Go CLI tools for converting and archiving data from major AI/email providers into portable, offline-first formats. Each tool is organized by provider and can be built and run independently.

**Core Principle:** Files first. Convert proprietary formats to open, queryable, human-readable files. No external dependencies where possible.

### Projects

- **`openai/chatgpt-json2md/`** — Converts ChatGPT JSON exports to per-conversation Markdown files. Input: `conversations.json` or `conversations-*.json`. Output: `markdown_output/`.
- **`openai/chatgpt-dat2assets/`** — Detects and restores correct file extensions from `.dat` files using magic byte detection (PNG, JPEG, WebP, WAV).
- **`anthropic/claude-go/`** — Converts Claude JSON exports to per-conversation Markdown files. Input: `conversations.json`. Output: `markdown_output/`.
- **`google/gmail/`** — Goodbye Gmail: converts Gmail accounts (active or archived) to a local, searchable corpus. Raw `.eml` files, normalized JSON/TXT, SQLite FTS5 indexing, attachment extraction. Most complex project with internal packages and external dependencies.

## Building & Running

Each project is self-contained and can be built from its own directory.

### Simple Projects (chatgpt-json2md, chatgpt-dat2assets, claude-go)

These are single-file, zero-dependency tools using only the Go standard library.

```bash
# Run directly
go run main.go

# Build binary
go build -o binary-name

# Run built binary
./binary-name
```

### Goodbye Gmail (google/gmail)

This project has internal packages and external dependencies. It's driven by CLI subcommands.

```bash
cd google/gmail

# Authenticate with Google (opens browser for OAuth)
go run ./... auth

# Fetch raw emails as .eml files (with resumable checkpoints)
go run ./... fetch --since "2024-01-01" --workers 10

# Process emails: parse, extract attachments, build SQLite index
go run ./... normalize

# Search local FTS5 index
go run ./... search "keywords"

# Rebuild or fix the SQLite index
go run ./... index

# Export databases to JSON
go run ./... export

# List all extracted attachments (with SHA256 hashes)
go run ./... attachments

# Build binary
go build -o goodbye-gmail ./...
```

## Project Structure & Architecture

### Simple Converters (OpenAI ChatGPT, Anthropic Claude)

- **Pattern:** Read JSON → Parse → Sanitize filenames → Write Markdown
- **Key behaviors:**
  - Untitled conversations → `Untitled Conversation.md`
  - Filename collisions resolved by appending identifier (ChatGPT: epoch; Claude: UUID prefix)
  - Filename safety: strips or replaces `/ \ : * ? " < > |`; truncates titles >100 chars
  - Messages sorted chronologically by `create_time` (ChatGPT) or already ordered (Claude)
  - Status/system messages and French labels (`Créé le :`, `Analyse de…`) are baked in
  - Non-string message parts (tool calls, attachments) → JSON-marshaled inline

### DAT to Assets

- **Pattern:** Glob for `.dat` files → detect magic bytes (first 4–12 bytes) → rename to correct extension
- **Supported formats:** PNG (89 50 4E 47), JPEG (FF D8 FF), WebP/WAV (RIFF header)
- **Behavior:** Only renames if magic bytes match; silently skips unknown formats

### Goodbye Gmail (google/gmail)

**Layered architecture:**

1. **Config layer** (`config/`): OAuth2 credentials and token storage (YAML + JSON).
2. **Storage layer** (`internal/storage/`):
   - `fs.go` — Filesystem abstraction (raw `.eml` storage, normalized outputs, attachments with SHA256 prefixes).
   - `sqlite.go` — Database schema and raw SQL operations.
   - `sqlite_repo.go` — Repo-pattern wrapper over SQLite (queries, indexing, FTS5).
3. **Model layer** (`internal/models/`): Email, contact, attachment structures.
4. **Processor layer** (`internal/processor/`):
   - `normalizer.go` — Parse `.eml` files, extract plain text, convert HTML → text.
   - `fragments.go` — Signature and quoted-reply detection and stripping.

**Key design choices:**
- **File-first:** Raw `.eml` and normalized JSON/TXT live on disk; SQLite is the index, not the source of truth.
- **Deduplication:** Attachments hashed via SHA256; stored once per unique hash.
- **Concurrency:** Worker pool pattern for fetching (respects HTTP 429 backoffs natively).
- **Resumable sync:** `.sync_checkpoint` file tracks progress; `fetch` can resume mid-export.
- **Contact parsing:** Deep extraction of `From/To/Cc` fields; signatures and quoted replies mapped in metadata.

**Supported dependencies:**
- `enmime` — MIME email parsing.
- Official Google Gmail API (OAuth2).

## Development Guidelines

### Code Organization

- **Simple tools:** Single `main.go`, standard library only. No external packages needed.
- **Goodbye Gmail:** Organized into `internal/` packages by concern (config, storage, models, processor). Follow Go's `internal/` convention: packages here can only be imported from within the same module.

### Input/Output Conventions

- **All tools** assume input files are in the current working directory (or specified paths).
- **Output** goes to `markdown_output/` (converters) or `data/` subdirectories (Goodbye Gmail).
- **No flags/config files** for simple converters; they auto-detect input (e.g., `conversations*.json`).

### Error Handling

- Simple tools: Printf to stderr and early return on critical errors.
- Goodbye Gmail: Structured error handling with context (API errors, file I/O, parse failures).

### Testing

- Simple tools: Manual testing with real export files (they're small).
- Goodbye Gmail: Unit tests on normalizer and fragment parsing (uses `internal/` packages).
- Run: `go test ./...` from the project root.

### Go Version

All projects declare `go 1.26+` in `go.mod`. Newer versions will likely work; older versions are untested.

## Common Tasks

### Adding a New Converter

1. Create a new provider directory (e.g., `providers/{name}/`).
2. Write a single `main.go` using only standard library.
3. Define JSON structs for the export format.
4. Parse → sanitize filenames → write Markdown (or chosen format).
5. Update README with usage and output format.

### Debugging Goodbye Gmail

- **OAuth not opening browser?** Check `config/account.yaml` exists and has valid credentials.
- **Fetch stalls?** Check `.sync_checkpoint` file; it records progress. Delete to restart from beginning.
- **SQLite errors?** Run `goodbye-gmail index` to rebuild from the JSON filesystem.
- **Attachment extraction issues?** Check `data/attachments/` directory exists and disk space available.

### Modifying Converters

- **Change output format:** Edit the Markdown template in the loop where messages are written.
- **Support new message types:** Extend the sender/role check (e.g., `if role == "custom"...`).
- **Handle collisions differently:** Modify the collision-detection logic (currently appends identifier).

## License

All code is MIT-licensed. See `LICENSE` at the repo root.
