package processor

import (
	"regexp"
	"strings"
)

// StripQuotes removes standard email reply quotes correctly.
func StripQuotes(text string) string {
	lines := strings.Split(text, "\n")
	var out []string

	// Basic heuristic: stop capturing when we hit "On <date>, <person> wrote:"
	// Or standard "-- " signature blocks.
	quoteRegex := regexp.MustCompile(`^(On\s.*wrote:|Le\s.*a\sécrit\s?:|---.*\sOriginal Message\s.*---)`)
	sigRegex := regexp.MustCompile(`^--\s*$`)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if quoteRegex.MatchString(trimmed) {
			break
		}
		if sigRegex.MatchString(trimmed) {
			break
		}
		// Also stop if the line itself is a prefix quote
		if strings.HasPrefix(trimmed, ">") && len(trimmed) > 1 {
			continue // Or break entirely if we want to drop the whole block
		}
		out = append(out, line)
	}
	return strings.TrimSpace(strings.Join(out, "\n"))
}

// ExtractCleanText returns a significantly cleaner version of the email.
func ExtractCleanText(rawText string) string {
	return StripQuotes(rawText)
}
