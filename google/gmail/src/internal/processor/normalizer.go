package processor

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jhillyerd/enmime"
	"github.com/user/goodbye-gmail/src/internal/models"
	"github.com/user/goodbye-gmail/src/internal/storage"
)

type Normalizer struct {
	fs *storage.FSConfig
	db *storage.DB
}

func NewNormalizer(fs *storage.FSConfig, db *storage.DB) *Normalizer {
	return &Normalizer{fs: fs, db: db}
}

// ProcessRaw takes a raw .eml payload, a gmail message ID and thread ID, and parses/saves it.
func (n *Normalizer) ProcessRaw(gmailID, threadID string, raw []byte) error {
	env, err := enmime.ReadEnvelope(bytes.NewReader(raw))
	if err != nil {
		return fmt.Errorf("enmime parse error for %s: %w", gmailID, err)
	}

	dateStr := env.GetHeader("Date")
	sentAt, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", dateStr)
	if err != nil {
		sentAt = time.Now() // Fallback if unparseable
	}

	msgID := env.GetHeader("Message-ID")
	if msgID == "" {
		msgID = gmailID // Fallback
	}

	subject := env.GetHeader("Subject")
	fromAddr := env.GetHeader("From")
	toAddr := env.GetHeader("To")
	ccAddr := env.GetHeader("Cc")

	hasAtt := len(env.Attachments) > 0 || len(env.Inlines) > 0

	// 1. Save raw .eml
	rawPath := n.fs.RawMessagePath(gmailID, sentAt)
	if err := writeFile(rawPath, raw); err != nil {
		return fmt.Errorf("failed writing raw message: %w", err)
	}

	// 2. Save normalized .txt
	txtPath := n.fs.NormalizedTextPath(gmailID, sentAt)
	if err := writeFile(txtPath, []byte(env.Text)); err != nil {
		return fmt.Errorf("failed writing normalized txt: %w", err)
	}

	// Prepare fragment text
	cleanText := ExtractCleanText(env.Text)

	msg := &models.Message{
		GmailID:            gmailID,
		MessageIDHeader:    msgID,
		ThreadID:           threadID,
		Direction:          "in",
		SentAt:             sentAt,
		Subject:            subject,
		FromJSON:           fmt.Sprintf(`{"address": %q}`, fromAddr),
		ToJSON:             fmt.Sprintf(`{"address": %q}`, toAddr),
		CcJSON:             fmt.Sprintf(`{"address": %q}`, ccAddr),
		BccJSON:            "",
		LabelsJSON:         "[]",
		RawPath:            rawPath,
		NormalizedTxtPath:  txtPath,
		NormalizedJSONPath: "",
		HasAttachments:     hasAtt,
		ExtractedText:      cleanText,
	}

	// 3. Save JSON metadata
	jsonPath := n.fs.NormalizedJSONPath(gmailID, sentAt)
	b, _ := json.MarshalIndent(msg, "", "  ")
	if err := writeFile(jsonPath, b); err != nil {
		return fmt.Errorf("failed writing normalized json: %w", err)
	}
	msg.NormalizedJSONPath = jsonPath

	// 4. Update SQLite Message
	dbID, err := n.db.SaveMessage(msg)
	if err != nil {
		return fmt.Errorf("sqlite save message error: %w", err)
	}

	// 4.5 Link Contacts safely
	n.parseAndRegisterContacts(dbID, fromAddr, "from", sentAt)
	n.parseAndRegisterContacts(dbID, toAddr, "to", sentAt)
	n.parseAndRegisterContacts(dbID, ccAddr, "cc", sentAt)

	// 5. Extract and Index Attachments
	for _, att := range env.Attachments {
		if err := n.processAttachment(dbID, msg, att); err != nil {
			return err
		}
	}
	for _, att := range env.Inlines {
		if err := n.processAttachment(dbID, msg, att); err != nil {
			return err
		}
	}

	return nil
}

func (n *Normalizer) parseAndRegisterContacts(msgID int64, rawHeader string, role string, sentAt time.Time) {
	if rawHeader == "" {
		return
	}
	addrs, _ := mail.ParseAddressList(rawHeader)
	for _, a := range addrs {
		if a.Address == "" {
			continue
		}
		cid, err := n.db.RegisterContact(strings.ToLower(a.Address), a.Name, sentAt)
		if err == nil {
			_ = n.db.LinkMessageContact(msgID, cid, role)
		}
	}
}

func (n *Normalizer) processAttachment(msgID int64, msg *models.Message, att *enmime.Part) error {
	sum := sha256.Sum256(att.Content)
	hashHex := fmt.Sprintf("%x", sum)

	attPath := n.fs.AttachmentPath(hashHex, att.FileName, msg.SentAt)
	if err := writeFile(attPath, att.Content); err != nil {
		return fmt.Errorf("failed writing attachment %s: %w", att.FileName, err)
	}

	dbAtt := &models.Attachment{
		MessageRef:     msgID,
		ThreadID:       msg.ThreadID,
		SentAt:         msg.SentAt,
		Filename:       att.FileName,
		StoredFilename: hashHex + "__" + att.FileName,
		MimeType:       att.ContentType,
		SizeBytes:      int64(len(att.Content)),
		Sha256:         hashHex,
		DiskPath:       attPath,
		ContactEmail:   "", // To map later
	}

	_, err := n.db.SaveAttachment(dbAtt)
	return err
}

func writeFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
