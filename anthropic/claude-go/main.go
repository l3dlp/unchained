package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Structure adaptée au JSON d'export d'Anthropic Claude
type ClaudeConversation struct {
	UUID        string          `json:"uuid"`
	Name        string          `json:"name"`
	CreatedAt   string          `json:"created_at"` // Claude utilise des chaînes ISO 8601 (ex: "2024-03-15T10:20:30Z")
	ChatMessages []ClaudeMessage `json:"chat_messages"`
}

type ClaudeMessage struct {
	UUID      string          `json:"uuid"`
	Sender    string          `json:"sender"` // "human" ou "assistant"
	CreatedAt string          `json:"created_at"`
	Text      string          `json:"text"`
}

func main() {
	outputDir := "markdown_output"
	inputFile := "conversations.json"

	if _, err := os.Stat(inputFile); err != nil {
		fmt.Printf("Erreur : Le fichier %s est introuvable.\n", inputFile)
		return
	}

	_ = os.MkdirAll(outputDir, os.ModePerm)
	fmt.Printf("Analyse du fichier %s (Export Claude)...\n", inputFile)

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return
	}

	var conversations []ClaudeConversation
	if err := json.Unmarshal(data, &conversations); err != nil {
		fmt.Printf("Erreur de parsing JSON : %v\n", err)
		return
	}

	totalChats := len(conversations)
	writtenFiles := 0

	for _, conv := range conversations {
		if conv.Name == "" {
			conv.Name = "Untitled Conversation"
		}

		// Pas besoin de trier, Claude fournit déjà les messages dans l'ordre chronologique
		if len(conv.ChatMessages) == 0 {
			continue
		}

		safeTitle := sanitizeFilename(conv.Name)
		filename := fmt.Sprintf("%s.md", safeTitle)
		filePath := filepath.Join(outputDir, filename)

		// Gestion des collisions
		if _, err := os.Stat(filePath); err == nil {
			// En cas de doublon, on utilise l'UUID de la conversation pour différencier
			shortUUID := conv.UUID
			if len(shortUUID) > 8 {
				shortUUID = shortUUID[:8]
			}
			filename = fmt.Sprintf("%s-%s.md", safeTitle, shortUUID)
			filePath = filepath.Join(outputDir, filename)
		}

		var md strings.Builder
		md.WriteString(fmt.Sprintf("# %s\n\n", conv.Name))
		
		if conv.CreatedAt != "" {
			// Parsing du format ISO 8601 de Claude
			if t, err := time.Parse(time.RFC3339, conv.CreatedAt); err == nil {
				md.WriteString(fmt.Sprintf("> **Créé le :** %s\n\n", t.Format("2006-01-02 15:04:05")))
			}
		}
		md.WriteString("---\n\n")

		for _, msg := range conv.ChatMessages {
			role := strings.ToLower(msg.Sender)
			if role == "system" || role == "" {
				continue
			}

			if role == "human" {
				role = "👤 **User**"
			} else if role == "assistant" {
				role = "🤖 **Claude**"
			} else {
				role = fmt.Sprintf("👤 **%s**", msg.Sender)
			}

			md.WriteString(fmt.Sprintf("### %s\n\n", role))
			md.WriteString(msg.Text)
			md.WriteString("\n\n---\n\n")
		}

		err = os.WriteFile(filePath, []byte(md.String()), 0644)
		if err == nil {
			writtenFiles++
		}
	}

	fmt.Printf("✅ Terminé ! %d fichiers Markdown créés sur %d conversations.\n", writtenFiles, totalChats)
}

func sanitizeFilename(name string) string {
	replacer := strings.NewReplacer(
		"/", "-", "\\", "-", ":", "", "*", "", "?", "", "\"", "", "<", "", ">", "", "|", "",
	)
	safe := replacer.Replace(name)
	if len(safe) > 100 {
		safe = safe[:100]
	}
	return strings.TrimSpace(safe)
}
