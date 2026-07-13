# Goodbye Gmail

**Goodbye Gmail** is an open source Go tool that transforms a Gmail account (active or archive) into a local, clean, readable, queryable, and reusable corpus. 

It does not seek to recreate Gmail. Instead, it returns an entire email history to the world of files, plain text, SQLite indexing, and Unix automation.

**Principle:** *Files first. SQLite second. Web later.*

## Features

- **Blazing Fast Networking:** Pulls raw `.eml` files heavily using concurrent goroutine worker pools. Respects HTTP 429 timeouts natively using Exponential Backoffs.
- **Offline Normalization:** Emails are parsed purely locally (`enmime`). Heavy HTML is converted to readable plain text (`.txt`) and stable JSON (`.json`) metadata.
- **File-First Storage:** Raw `.eml`, normalized strings, and attachments are organized purely on the filesystem.
- **Organic Deduplication:** Attachments are extracted exactly once per email and hashed using `SHA-256`.
- **Deep Contact Parsing:** Maps `From/To/Cc` fields natively to SQLite indices, extracting out signatures and quoted reply chains in the process.
- **SQLite Index:** Everything is referenced in a lightweight SQLite database natively equipped with `FTS5` (Full-Text Search).

## Setup & Verification

Goodbye Gmail relies on the official Gmail API via OAuth 2.0.

1. **Obtain Google Cloud Credentials:**
   - Go to your [Google Cloud Console](https://console.cloud.google.com/).
   - Create a new project (e.g., "Goodbye Gmail").
   - Enable the **Gmail API**.
   - Go to "Credentials" > "Create Credentials" > "OAuth client ID".
   - Select "Desktop app" (or equivalent) and create.
   - Download the JSON credential file.

2. **Configure Local Account:**
   - Create the config directory: `mkdir -p config`
   - Create `config/account.yaml` based on your downloaded credentials:
     ```yaml
     email: "your.email@gmail.com"
     client_id: "YOUR_CLIENT_ID"
     secret: "YOUR_CLIENT_SECRET"
     ```

## Usage

Goodbye Gmail is entirely command-line driven. 

```bash
# 1. Authenticate your account (opens browser / prompts for URL token)
goodbye-gmail auth

# 2. Fetch raw .eml payloads directly to disk (using 10 concurrent workers)
# Automatically creates/resumes from .sync_checkpoint 
goodbye-gmail fetch --since "2024-01-01" --workers 10

# 3. Process raw .eml offline, extract attachments, map contacts, and build local SQLite tables
goodbye-gmail normalize

# 4. Search your local index instantly via FTS5!
goodbye-gmail search "your search keywords"

# Uncorrupt/Rebuild your SQLite index from the JSON filesystem natively
goodbye-gmail index

# Export generic databases into CLI JSON format
goodbye-gmail export

# List all extracted and hashed attachments locally recorded
goodbye-gmail attachments
```

## Structure

```text
goodbye-gmail/
  config/
    account.yaml
    token.json
  data/
    raw/           # Original .eml files
    normalized/    # Clean .txt and .json files
    attachments/   # All downloaded files, prefixed with SHA256 hashes
    .sync_checkpoint # Internal state for resuming syncs
  index/
    goodbye-gmail.sqlite
```
