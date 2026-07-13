package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Dossier cible (adaptez le chemin si nécessaire)
	targetDir := "."

	files, err := filepath.Glob(filepath.Join(targetDir, "*.dat"))
	if err != nil {
		fmt.Printf("Erreur lors de la recherche des fichiers .dat: %v\n", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("Aucun fichier .dat trouvé dans le dossier ciblé.")
		return
	}

	fmt.Printf("Analyse de %d fichiers .dat...\n", len(files))
	renamedCount := 0

	for _, oldPath := range files {
		ext, err := detectExtension(oldPath)
		if err != nil || ext == "" {
			// Extension inconnue ou erreur, on ne touche à rien
			continue
		}

		// Remplacement de l'extension .dat par la vraie extension détectée
		newPath := oldPath[:len(oldPath)-4] + ext
		err = os.Rename(oldPath, newPath)
		if err == nil {
			renamedCount++
		}
	}

	fmt.Printf("✅ Terminé ! %d fichiers .dat ont retrouvé leur vraie extension.\n", renamedCount)
}

// detectExtension lit l'en-tête (Magic Bytes) pour deviner le format réel
func detectExtension(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 12 octets suffisent pour identifier PNG, JPEG, WebP et WAV (RIFF)
	buffer := make([]byte, 12)
	_, err = io.ReadFull(file, buffer)
	if err != nil && err != io.ErrUnexpectedEOF {
		return "", err
	}

	// Détection PNG
	if len(buffer) >= 4 && buffer[0] == 0x89 && buffer[1] == 0x50 && buffer[2] == 0x4E && buffer[3] == 0x47 {
		return ".png", nil
	}

	// Détection JPEG
	if len(buffer) >= 3 && buffer[0] == 0xFF && buffer[1] == 0xD8 && buffer[2] == 0xFF {
		return ".jpg", nil
	}

	// Détection formats conteneurs RIFF (WebP ou WAV)
	if len(buffer) >= 12 && string(buffer[0:4]) == "RIFF" {
		format := string(buffer[8:12])
		if format == "WEBP" {
			return ".webp", nil
		}
		if format == "WAVE" {
			return ".wav", nil
		}
	}

	return "", nil
}

