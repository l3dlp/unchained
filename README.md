# Getting Back Your Data

**Languages:** [English](README.md) • [Français](README.fr.md) • [Español](README.es.md) • [Português](README.pt.md) • [Italiano](README.it.md)

You've been giving your thoughts, memories, and emails to corporations for over a decade. They've mined it, analyzed it, sold it to the highest bidder. Now you want it *back*. In a format that doesn't require their permission to read. On *your* machine. Forever.

This repo is your toolkit.

## The Tools

### ChatGPT JSON to Markdown
`openai/chatgpt-json2md`

Your ChatGPT export is sitting there as a wall of JSON. We turn it into clean, readable Markdown files—one per conversation. No bullshit, no cloud lookups, no "sorry the API is down." Just markdown you can grep, edit, and throw into Obsidian, Logseq, or your plaintext vault of choice.

- Zero dependencies. Standard library only.
- Handles multiple export files. Auto-detects batch exports.
- Cleans up filenames. Sorts messages by timestamp.
- Everything is *yours* the second it leaves OpenAI's server.

```bash
go run main.go
```

Read the [full docs](openai/chatgpt-json2md/README.md).

### DAT to Assets
`openai/chatgpt-dat2assets`

Your ChatGPT export contains images, audio, documents—but they're all renamed `.dat` because OpenAI figured you'd never reclaim them. We sniff the magic bytes and give them back their real names. PNG, JPEG, WebP, WAV. One pass, done.

```bash
go run main.go
```

### Claude JSON to Markdown
`anthropic/claude-go`

Same deal as ChatGPT. Anthropic gives you JSON, we give you readable Markdown. Keeps your conversation history portable and yours.

```bash
go run main.go
```

### Google Webfonts Manager
`google/webfonts`

Ever grabbed fonts from Google Fonts and ended up with a shoebox of TTF files? This CLI generates a static HTML GUI where you browse your local font library, pick variants, and copy the `@font-face` CSS blocks. Favorites saved in localStorage. Works offline. Multi-language.

```bash
go run ./cmd/mywebfont -root webfonts -out webfonts.html -base-url webfonts
```

Read the [full docs](google/webfonts/README.md).

### Goodbye Gmail
`google/gmail`

This one is different. It doesn't just grab your emails—it *archives* them. Pulls raw `.eml` files directly from the Gmail API, parses them locally, extracts every attachment, indexes everything in SQLite with FTS5 full-text search, and builds a filesystem you own outright.

- **Blazing fast.** Concurrent workers, respects rate limits with backoff.
- **Deduped.** Each unique attachment stored once, hashed via SHA256.
- **Indexed.** FTS5 search. Instant. Local. No API calls.
- **Resumable.** Crashes? Restarts from your checkpoint. No re-fetching.
- **Deep parsing.** Extracts contacts, signatures, quoted replies.

```bash
cd google/gmail
go run ./... auth          # OAuth into Google
go run ./... fetch --workers 10  # Pull your emails
go run ./... normalize     # Parse and index
go run ./... search "keywords"  # Instant search
```

Read the [full docs](google/gmail/README.md).

---

## Why This Matters

You didn't realize it when you were typing, but you were building an artifact. Every conversation with an AI was a decision point, an idea refined, a problem solved. Your emails are a chronicle of your life: relationships, work, struggles, growth.

These aren't just files. They're *your intellectual property*. And right now, they're locked behind login screens and terms of service that can change on a Tuesday.

We're in the endgame of the platform era. Corporations are tightening their grip because they know what they've built is valuable. They're not giving refunds. So we extract. We convert. We archive.

This repo is a declaration: *your data is yours*. Not temporarily. Not conditionally. Yours. Stored in formats that predate the web. Formats that will still be readable when the servers rot.

---

## Installation

Each project is standalone. Pick what you need.

```bash
# Clone
git clone https://github.com/l3dlp/unchained.guru.git
cd unchained.guru/rip

# Simple converters: no setup needed
cd openai/chatgpt-json2md
go run main.go

# Goodbye Gmail: needs Go 1.26+
cd google/gmail
go run ./... auth
```

**Requires:** Go 1.26 or later (older versions may work; untested).

---

## Formats

All output is **portable**:
- Markdown: readable in any editor, searchable with grep
- JSON: queryable, parseable, future-proof
- `.eml`: RFC 5322 standard. Import into Thunderbird, Apple Mail, anything
- SQLite: no vendor lock-in. Use the CLI, or query with Python/Perl/Go

You own the format. You own the data.

---

## The Unix Manifesto, Circa 1996

These tools follow what Unix users have known for 30 years:

- **Use boring formats.** Text. JSON. SQLite. They outlive fashion.
- **Do one thing well.** Extract. Convert. Index. Nothing more.
- **No cloud required.** Everything runs offline. No telemetry. No calls home.
- **Respect the user.** Assume you know what you're doing. No hand-holding, no guardrails.
- **Chain it together.** Pipe to grep, awk, jq. Combine with other tools. That's the point.

We're not building a platform. We're giving you **tools**. What you do with your data is your business.

---

## License

MIT. Use it. Fork it. Improve it. Don't give it back to corporations.

---

## Contribute

Find a bug? Data corrupted? New provider to support? Send a PR. No code of conduct. No governance committees. Just good faith and clean diffs.

---

**Your data. Your machine. Your rules.**
