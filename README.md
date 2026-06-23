# chatgpt-json2md

A tiny Go program that turns your ChatGPT data export into clean, per-conversation Markdown files — ready to drop into Obsidian, Logseq, or any other plain-text vault.

No dependencies. No config. No telemetry. One file, standard library only.

## What it does

1. Looks for `conversations-*.json` in the current directory. If none are found, falls back to `conversations.json`.
2. Parses each file as an array of ChatGPT conversations.
3. For every conversation:
   - Skips messages with the `system` role and messages with an empty role.
   - Sorts the remaining messages chronologically by `create_time`.
   - Writes a Markdown file named after the conversation title.
4. Drops every result into a `markdown_output/` directory next to the binary.

## Output format

Each `.md` file looks like this:

```markdown
# <conversation title>

> **Créé le :** 2025-04-12 18:32:07

---

### 👤 **User**

<message content>

---

### 🤖 **ChatGPT**

<message content>

---
```

Notes:

- Untitled conversations are saved as `Untitled Conversation.md`.
- Filename collisions are resolved by appending the conversation's `create_time` epoch: `My Chat-1712946727.md`.
- Filenames are sanitized: `/ \ : * ? " < > |` are stripped or replaced with `-`, and titles longer than 100 characters are truncated.
- Status messages (`Analyse de…`, `Terminé !`) and the `Créé le :` label are in French. The data itself stays untouched.
- Message `parts` that are plain strings are written as-is. Non-string parts (tool calls, attachments, structured payloads) are JSON-marshaled and dropped inline so nothing silently disappears.

## How to get your ChatGPT export

1. Open ChatGPT → **Settings** → **Data Controls** → **Export data**.
2. Wait for the email from OpenAI.
3. Download and unzip the archive. You will get one or more `conversations-*.json` files (or a single `conversations.json` on older exports).

## Usage

```sh
# put your export JSON file(s) next to the binary
go run main.go

# or build and run
go build -o chatgpt-json2md
./chatgpt-json2md
```

Then point Obsidian (or your editor of choice) at the `markdown_output/` folder.

## Requirements

- Go 1.26+ (declared in `go.mod`; older versions will likely build fine but are untested).
- No external Go modules. Only `encoding/json`, `os`, `path/filepath`, `sort`, `strings`, `time`, `fmt`.

## License

MIT — see [LICENSE](LICENSE).
