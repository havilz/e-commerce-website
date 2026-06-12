package security

import (
	"html"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

var (
	// strictPolicy removes ALL HTML tags, returning plain text.
	strictPolicy = bluemonday.StrictPolicy()
)

// SanitizeString strips all HTML, Javascript, and CSS styling tags from the input string.
// It also unescapes common HTML entities to prevent double-encoding.
func SanitizeString(input string) string {
	if input == "" {
		return ""
	}
	
	// Remove all HTML tags
	sanitized := strictPolicy.Sanitize(input)
	
	// Unescape HTML entities (e.g. &lt; -> <) just in case the input had pre-encoded characters,
	// then clean it again to be absolutely sure no new tags were introduced.
	sanitized = html.UnescapeString(sanitized)
	sanitized = strictPolicy.Sanitize(sanitized)
	
	return strings.TrimSpace(sanitized)
}

// SanitizeSQLWildcards escapes the wildcard characters '%' and '_' in SQLite LIKE query parameters.
// This prevents users from querying all records using simple wildcards.
func SanitizeSQLWildcards(input string) string {
	// Escape backslashes first, then escape % and _ using a backslash as the escape character.
	replacer := strings.NewReplacer(
		`\`, `\\`,
		`%`, `\%`,
		`_`, `\_`,
	)
	return replacer.Replace(input)
}
