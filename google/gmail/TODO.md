# Goodbye Gmail Exhaustive TODO

This file tracks the current state of the project. The `v1.0 MVP` has been fully implemented, securing the tool to be ran massively against heavy production workflows.

## ✅ Completed (v1.0 MVP Ready)
- [x] Initial Go module setup & Cobra CLI structure.
- [x] Storage structure definition (FS and paths generators setup).
- [x] SQLite schema initialized with `FTS5` virtual table indexing.
- [x] Gmail OAuth2 flow (`auth` command) established.
- [x] **Pagination robustness**: Seamlessly paginate through Google APIs.
- [x] **Rate-Limiting**: Applies Google recommended explicit exponential backoff for `429 Too Many Requests` API errors.
- [x] **Incremental sync**: Handles `--since` via `.sync_checkpoint` native writes to easily resume syncing.
- [x] **Concurrency Pipeline**: `fetch` heavily utilizes Go routines to pull ID channels rapidly.
- [x] **Separation of Concerns**: `fetch` and `normalize` are distinct scripts. `normalize` iterates blindly offline securing zero rate-limits.
- [x] **Idempotence**: Native `db.MessageExists(id)` checks ensure 10k loops run in seconds once already downloaded.
- [x] **CLI Progress visualizers**: Integrates `schollz/progressbar/v3`.
- [x] **Deep Contact Mapping Engine**: Natively uses `mail.ParseAddressList` across `From/To/Cc` creating distinct `sqlite` Contacts schemas and junctions.
- [x] **Fragment Extractor**: Basic native heuristic stripping of "reply chains / quotes" and "signature lines" extracting original thought from `.eml`.
- [x] **Thread mapping**: Internal API correctly tags threaded responses based off Google's `ThreadId`.
- [x] `search` command using blazing fast SQLite FTS5 indexes natively.
- [x] Default `export`, `attachments`, and `index` maintenance CLI apps established.
- [x] MIT License added, robust `.gitignore` implemented.

---

## 🔮 Future / V2.0 Enhancements

The next major phases of the project will focus around integrating these extractions seamlessly into downstream NLP architectures natively, and expanding parsing robustness:

- [ ] **Daemonization**: Create a `sync` polling loop capable of constantly pinging webhooks natively to keep archives hot for AI assistants.
- [ ] **Advanced HTML processing**: Strip CSS inline tags robustly for even lighter `.json` footprint using plugins like `PuerkitoBio/goquery`.
- [ ] **Advanced Language Detection**: Add native `lang` recognition metrics to the Fragments database automatically during processing.
- [ ] **CLI Graphical Interface (`Bubbletea`)**: Build out an optional highly interactive CLI GUI using `charmbracelet/bubbletea` to crawl threads intuitively without opening SQLite!
- [ ] **UI Web App Base**: Later implementation of the fundamental Web app viewer using standard SSR for easy navigations offline.
