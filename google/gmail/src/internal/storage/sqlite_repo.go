package storage

import (
	"fmt"
	"time"

	"github.com/user/goodbye-gmail/src/internal/models"
)

// SaveMessage inserts or ignores a message using the message_id_header as unique.
func (db *DB) SaveMessage(m *models.Message) (int64, error) {
	query := `
	INSERT INTO messages (
		gmail_message_id, message_id_header, thread_id, direction,
		sent_at, subject, from_json, to_json, cc_json, bcc_json,
		labels_json, raw_path, normalized_txt_path, normalized_json_path,
		has_attachments
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(message_id_header) DO UPDATE SET
		labels_json=excluded.labels_json,
		has_attachments=excluded.has_attachments
	RETURNING id;
	`
	
	sentAtStr := m.SentAt.Format("2006-01-02T15:04:05Z07:00")
	hasAtt := 0
	if m.HasAttachments {
		hasAtt = 1
	}

	var id int64
	err := db.QueryRow(query,
		m.GmailID, m.MessageIDHeader, m.ThreadID, m.Direction,
		sentAtStr, m.Subject, m.FromJSON, m.ToJSON, m.CcJSON, m.BccJSON,
		m.LabelsJSON, m.RawPath, m.NormalizedTxtPath, m.NormalizedJSONPath,
		hasAtt,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to save message: %w", err)
	}

	// Update search index
	ftsQuery := `
	INSERT INTO messages_fts(rowid, subject, content) 
	VALUES (?, ?, ?)
	ON CONFLICT(rowid) DO UPDATE SET
		subject=excluded.subject,
		content=excluded.content;
	`
	_, _ = db.Exec(ftsQuery, id, m.Subject, m.ExtractedText) // Ignore FTS errors for now

	return id, nil
}

// SaveAttachment saves an attachment metadata.
func (db *DB) SaveAttachment(att *models.Attachment) (int64, error) {
	query := `
	INSERT INTO attachments (
		message_ref, thread_id, sent_at, filename, stored_filename,
		mime_type, size_bytes, sha256, disk_path, contact_email
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	sentAtStr := att.SentAt.Format("2006-01-02T15:04:05Z07:00")
	
	res, err := db.Exec(query,
		att.MessageRef, att.ThreadID, sentAtStr, att.Filename, att.StoredFilename,
		att.MimeType, att.SizeBytes, att.Sha256, att.DiskPath, att.ContactEmail,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// MessageExists checks idempotence quickly.
func (db *DB) MessageExists(gmailID string) bool {
	var id int64
	err := db.QueryRow("SELECT id FROM messages WHERE gmail_message_id = ?", gmailID).Scan(&id)
	return err == nil
}

// RegisterContact registers a contact identity, incrementing counts.
func (db *DB) RegisterContact(email, displayName string, sent time.Time) (int64, error) {
	query := `
	INSERT INTO contacts (email, display_name, first_seen_at, last_seen_at, messages_count)
	VALUES (?, ?, ?, ?, 1)
	ON CONFLICT(email) DO UPDATE SET
		display_name = CASE WHEN excluded.display_name != '' THEN excluded.display_name ELSE contacts.display_name END,
		first_seen_at = min(contacts.first_seen_at, excluded.first_seen_at),
		last_seen_at = max(contacts.last_seen_at, excluded.last_seen_at),
		messages_count = contacts.messages_count + 1
	RETURNING id;
	`
	tStr := sent.Format(time.RFC3339)
	var id int64
	err := db.QueryRow(query, email, displayName, tStr, tStr).Scan(&id)
	return id, err
}

// LinkMessageContact binds a message to a contact role.
func (db *DB) LinkMessageContact(msgID, contactID int64, role string) error {
	q := `INSERT OR IGNORE INTO message_contacts (message_ref, contact_ref, role) VALUES (?, ?, ?)`
	_, err := db.Exec(q, msgID, contactID, role)
	return err
}
