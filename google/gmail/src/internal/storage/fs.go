package storage

import (
	"path/filepath"
	"time"
)

type FSConfig struct {
	BaseDir string
}

func NewFS(baseDir string) *FSConfig {
	return &FSConfig{BaseDir: baseDir}
}

func (fs *FSConfig) RawMessagePath(id string, sent time.Time) string {
	year := sent.Format("2006")
	month := sent.Format("2006-01")
	return filepath.Join(fs.BaseDir, "raw", year, month, id+".eml")
}

func (fs *FSConfig) NormalizedTextPath(id string, sent time.Time) string {
	year := sent.Format("2006")
	month := sent.Format("2006-01")
	return filepath.Join(fs.BaseDir, "normalized", year, month, id+".txt")
}

func (fs *FSConfig) NormalizedJSONPath(id string, sent time.Time) string {
	year := sent.Format("2006")
	month := sent.Format("2006-01")
	return filepath.Join(fs.BaseDir, "normalized", year, month, id+".json")
}

func (fs *FSConfig) AttachmentPath(hash string, filename string, sent time.Time) string {
	year := sent.Format("2006")
	month := sent.Format("2006-01")
	storedName := hash + "__" + filename
	return filepath.Join(fs.BaseDir, "attachments", year, month, storedName)
}

func (fs *FSConfig) ThreadDirPath(threadID string) string {
	return filepath.Join(fs.BaseDir, "threads", threadID)
}

func (fs *FSConfig) ContactDirPath(email string) string {
	return filepath.Join(fs.BaseDir, "contacts", email)
}
