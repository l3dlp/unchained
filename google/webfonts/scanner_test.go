package mywebfont

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanInfersVariants(t *testing.T) {
	root := t.TempDir()
	family := filepath.Join(root, "Alegreya")
	if err := os.Mkdir(family, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, file := range []string{"regular.ttf", "italic.ttf", "700.ttf", "700italic.ttf", "date.txt"} {
		if err := os.WriteFile(filepath.Join(family, file), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	catalog, err := Scan(root)
	if err != nil {
		t.Fatal(err)
	}
	if got := len(catalog.Families); got != 1 {
		t.Fatalf("families = %d, want 1", got)
	}
	variants := catalog.Families[0].Variants
	if got := len(variants); got != 4 {
		t.Fatalf("variants = %d, want 4", got)
	}
	if variants[0].Weight != 400 || variants[0].Style != "italic" {
		t.Fatalf("first variant = %#v, want 400 italic after sort by style", variants[0])
	}
	if variants[3].Weight != 700 || variants[3].Style != "normal" {
		t.Fatalf("last variant = %#v, want 700 normal", variants[3])
	}
}

func TestParseVariant(t *testing.T) {
	tests := map[string]Variant{
		"regular.ttf":    {File: "regular.ttf", Weight: 400, Style: "normal", Label: "regular"},
		"italic.ttf":     {File: "italic.ttf", Weight: 400, Style: "italic", Label: "italic"},
		"300italic.ttf":  {File: "300italic.ttf", Weight: 300, Style: "italic", Label: "300 italic"},
		"ExtraBold.woff": {File: "ExtraBold.woff", Weight: 800, Style: "normal", Label: "800 normal"},
	}
	for file, want := range tests {
		if got := parseVariant(file); got != want {
			t.Fatalf("parseVariant(%q) = %#v, want %#v", file, got, want)
		}
	}
}
