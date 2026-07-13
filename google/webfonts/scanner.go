package mywebfont

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Catalog struct {
	Families []Family `json:"families"`
}

type Family struct {
	Name     string    `json:"name"`
	Dir      string    `json:"dir"`
	Preview  string    `json:"preview,omitempty"`
	Variants []Variant `json:"variants"`
}

type Variant struct {
	File   string `json:"file"`
	Weight int    `json:"weight"`
	Style  string `json:"style"`
	Label  string `json:"label"`
}

var fontExts = map[string]bool{
	".ttf":   true,
	".otf":   true,
	".woff":  true,
	".woff2": true,
}

var weightName = map[string]int{
	"thin":       100,
	"extralight": 200,
	"ultralight": 200,
	"light":      300,
	"regular":    400,
	"normal":     400,
	"medium":     500,
	"semibold":   600,
	"demibold":   600,
	"bold":       700,
	"extrabold":  800,
	"ultrabold":  800,
	"black":      900,
	"heavy":      900,
}

var leadingWeight = regexp.MustCompile(`^(100|200|300|400|500|600|700|800|900)`)

func Scan(root string) (Catalog, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return Catalog{}, err
	}

	var families []Family
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		family, err := scanFamily(filepath.Join(root, entry.Name()), entry.Name())
		if err != nil {
			return Catalog{}, err
		}
		if len(family.Variants) > 0 {
			families = append(families, family)
		}
	}

	sort.Slice(families, func(i, j int) bool {
		return strings.ToLower(families[i].Name) < strings.ToLower(families[j].Name)
	})

	return Catalog{Families: families}, nil
}

func scanFamily(path, name string) (Family, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return Family{}, err
	}

	var variants []Variant
	var preview string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		file := entry.Name()
		ext := strings.ToLower(filepath.Ext(file))
		switch {
		case fontExts[ext]:
			variants = append(variants, parseVariant(file))
		case ext == ".png" && preview == "":
			preview = file
		}
	}

	sort.Slice(variants, func(i, j int) bool {
		if variants[i].Weight != variants[j].Weight {
			return variants[i].Weight < variants[j].Weight
		}
		if variants[i].Style != variants[j].Style {
			return variants[i].Style < variants[j].Style
		}
		return variants[i].File < variants[j].File
	})

	return Family{Name: name, Dir: name, Preview: preview, Variants: variants}, nil
}

func parseVariant(file string) Variant {
	base := strings.TrimSuffix(file, filepath.Ext(file))
	token := strings.ToLower(strings.NewReplacer("-", "", "_", "", " ", "").Replace(base))
	style := "normal"
	if strings.Contains(token, "italic") {
		style = "italic"
		token = strings.ReplaceAll(token, "italic", "")
	}

	weight := 400
	if token != "" {
		if m := leadingWeight.FindString(token); m != "" {
			if parsed, err := strconv.Atoi(m); err == nil {
				weight = parsed
			}
		} else if named, ok := weightName[token]; ok {
			weight = named
		}
	}

	label := fmt.Sprintf("%d %s", weight, style)
	if file == "regular.ttf" || file == "regular.woff2" || file == "regular.woff" || file == "regular.otf" {
		label = "regular"
	} else if strings.HasPrefix(file, "italic.") {
		label = "italic"
	}

	return Variant{File: file, Weight: weight, Style: style, Label: label}
}
