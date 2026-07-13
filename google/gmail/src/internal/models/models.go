package models

import "time"

type Message struct {
	ID                 int64
	GmailID            string
	MessageIDHeader    string
	ThreadID           string
	Direction          string // "in" or "out"
	SentAt             time.Time
	Subject            string
	FromJSON           string
	ToJSON             string
	CcJSON             string
	BccJSON            string
	LabelsJSON         string
	RawPath            string
	NormalizedTxtPath  string
	NormalizedJSONPath string
	HasAttachments     bool

	// Fields for FTS index
	ExtractedText string
}

type Attachment struct {
	ID             int64
	MessageRef     int64
	ThreadID       string
	SentAt         time.Time
	Filename       string
	StoredFilename string
	MimeType       string
	SizeBytes      int64
	Sha256         string
	DiskPath       string
	ContactEmail   string
}

type Contact struct {
	ID            int64
	Email         string
	DisplayName   string
	FirstSeenAt   time.Time
	LastSeenAt    time.Time
	MessagesCount int
}

type Fragment struct {
	ID           int64
	MessageRef   int64
	ThreadID     string
	SentAt       time.Time
	FragmentType string
	Lang         string
	Tone         string
	TextRaw      string
	TextClean    string
	Score        float64
}
