package storage

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func OpenDB(baseDir string) (*DB, error) {
	dbPath := filepath.Join(baseDir, "index", "goodbye-gmail.sqlite")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) InitSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY,
		gmail_message_id TEXT,
		message_id_header TEXT UNIQUE,
		thread_id TEXT,
		direction TEXT,
		sent_at TEXT,
		subject TEXT,
		from_json TEXT,
		to_json TEXT,
		cc_json TEXT,
		bcc_json TEXT,
		labels_json TEXT,
		raw_path TEXT,
		normalized_txt_path TEXT,
		normalized_json_path TEXT,
		has_attachments INTEGER DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS attachments (
		id INTEGER PRIMARY KEY,
		message_ref INTEGER NOT NULL,
		thread_id TEXT,
		sent_at TEXT,
		filename TEXT,
		stored_filename TEXT,
		mime_type TEXT,
		size_bytes INTEGER,
		sha256 TEXT,
		disk_path TEXT,
		contact_email TEXT,
		FOREIGN KEY(message_ref) REFERENCES messages(id)
	);

	CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY,
		email TEXT UNIQUE,
		display_name TEXT,
		first_seen_at TEXT,
		last_seen_at TEXT,
		messages_count INTEGER DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS message_contacts (
		message_ref INTEGER NOT NULL,
		contact_ref INTEGER NOT NULL,
		role TEXT NOT NULL,
		PRIMARY KEY (message_ref, contact_ref, role),
		FOREIGN KEY(message_ref) REFERENCES messages(id),
		FOREIGN KEY(contact_ref) REFERENCES contacts(id)
	);

	CREATE TABLE IF NOT EXISTS fragments (
		id INTEGER PRIMARY KEY,
		message_ref INTEGER NOT NULL,
		thread_id TEXT,
		sent_at TEXT,
		fragment_type TEXT,
		lang TEXT,
		tone TEXT,
		text_raw TEXT,
		text_clean TEXT,
		score REAL DEFAULT 0,
		FOREIGN KEY(message_ref) REFERENCES messages(id)
	);

	-- FTS5 table for efficient search
	CREATE VIRTUAL TABLE IF NOT EXISTS messages_fts USING fts5(
		subject, 
		content,
		content='messages',
		content_rowid='id'
	);
	`
	_, err := db.Exec(schema)
	return err
}
