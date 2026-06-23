package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Conversation struct {
	Title      string          `json:"title"`
	CreateTime float64         `json:"create_time"`
	Mapping    map[string]Node `json:"mapping"`
}

type Node struct {
	ID      string   `json:"id"`
	Message *Message `json:"message,omitempty"`
}

type Message struct {
	ID         string     `json:"id"`
	Author     Author     `json:"author"`
	CreateTime float64    `json:"create_time"`
	Content    Content    `json:"content"`
}

type Author struct {
	Role string `json:"role"`
}

type Content struct {
	ContentType string        `json:"content_type"`
	Parts       []interface{} `json:"parts"`
}

func main() {
	outputDir := "markdown_output"

	files, err := filepath.Glob("conversations-*.json")
	if err != nil {
		fmt.Printf("Erreur scan : %v\n", err)
		return
	}

	if len(files) == 0 {
		if _, err := os.Stat("conversations.json"); err == nil {
			files = append(files, "conversations.json")
		} else {
			fmt.Println("Erreur : Fichiers conversations-*.json introuvables.")
			return
		}
	}

	_ = os.MkdirAll(outputDir, os.ModePerm)
	fmt.Printf("Analyse de %d fichiers JSON...\n", len(files))

	totalChats := 0
	writtenFiles := 0

	for _, jsonFile := range files {
		data, err := os.ReadFile(jsonFile)
		if err != nil {
			continue
		}

		var conversations []Conversation
		if err := json.Unmarshal(data, &conversations); err != nil {
			continue
		}

		totalChats += len(conversations)

		for _, conv := range conversations {
			if conv.Title == "" {
				conv.Title = "Untitled Conversation"
			}

			// Extraction directe sans passer par l'arbre de nœuds
			orderedMessages := getSortedMessagesDirect(conv.Mapping)
			if len(orderedMessages) == 0 {
				continue
			}

			safeTitle := sanitizeFilename(conv.Title)
			filename := fmt.Sprintf("%s.md", safeTitle)
			filePath := filepath.Join(outputDir, filename)

			if _, err := os.Stat(filePath); err == nil {
				filename = fmt.Sprintf("%s-%d.md", safeTitle, int64(conv.CreateTime))
				filePath = filepath.Join(outputDir, filename)
			}

			var md strings.Builder
			md.WriteString(fmt.Sprintf("# %s\n\n", conv.Title))
			
			if conv.CreateTime > 0 {
				t := time.Unix(int64(conv.CreateTime), 0)
				md.WriteString(fmt.Sprintf("> **Créé le :** %s\n\n", t.Format("2006-01-02 15:04:05")))
			}
			md.WriteString("---\n\n")

			for _, msg := range orderedMessages {
				role := strings.ToUpper(msg.Author.Role)
				if role == "SYSTEM" || role == "" {
					continue
				}

				if role == "USER" {
					role = "👤 **User**"
				} else if role == "ASSISTANT" {
					role = "🤖 **ChatGPT**"
				}

				md.WriteString(fmt.Sprintf("### %s\n\n", role))
				
				for _, part := range msg.Content.Parts {
					// Extraction textuelle ultra-permissive
					if strPart, ok := part.(string); ok {
						md.WriteString(strPart)
						md.WriteString("\n")
					} else {
						// Si OpenAI a imbriqué du JSON (fichiers, outils)
						bytes, _ := json.Marshal(part)
						md.WriteString(string(bytes))
						md.WriteString("\n")
					}
				}
				md.WriteString("\n---\n\n")
			}

			err = os.WriteFile(filePath, []byte(md.String()), 0644)
			if err == nil {
				writtenFiles++
			}
		}
	}

	fmt.Printf("✅ Terminé ! %d fichiers Markdown créés sur %d conversations.\n", writtenFiles, totalChats)
}

// Extraction plate et tri chronologique
func getSortedMessagesDirect(mapping map[string]Node) []*Message {
	var list []*Message

	for _, node := range mapping {
		if node.Message == nil {
			continue
		}
		
		// On ignore les messages vides ou systèmes immédiatement
		role := strings.ToUpper(node.Message.Author.Role)
		if role == "SYSTEM" || role == "" {
			continue
		}

		// On s'assure qu'il y a du contenu (n'importe quel type)
		if len(node.Message.Content.Parts) > 0 {
			list = append(list, node.Message)
		}
	}

	// Tri chronologique basé sur le timestamp de chaque message
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreateTime < list[j].CreateTime
	})

	return list
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

