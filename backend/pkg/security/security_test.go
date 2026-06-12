package security

import (
	"testing"
)

func TestSanitizeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Clean string unchanged",
			input:    "Baju Kaos Keren",
			expected: "Baju Kaos Keren",
		},
		{
			name:     "HTML script tags stripped",
			input:    "<script>alert('hack')</script>Baju Kaos Keren",
			expected: "Baju Kaos Keren",
		},
		{
			name:     "Embedded HTML elements stripped",
			input:    "<div>Hello <b>World</b></div>",
			expected: "Hello World",
		},
		{
			name:     "Onclick attributes stripped",
			input:    "<a href='http://test.com' onclick='stealCookies()'>Link</a>",
			expected: "Link",
		},
		{
			name:     "Pre-encoded HTML entities handled",
			input:    "&lt;script&gt;alert(1)&lt;/script&gt;Safe Text",
			expected: "Safe Text",
		},
		{
			name:     "Whitespace trimmed",
			input:    "   Trimmed Text   ",
			expected: "Trimmed Text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SanitizeString(tt.input)
			if got != tt.expected {
				t.Errorf("SanitizeString(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestSanitizeSQLWildcards(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Normal search string",
			input:    "macbook",
			expected: "macbook",
		},
		{
			name:     "Percent sign escaped",
			input:    "100%",
			expected: `100\%`,
		},
		{
			name:     "Underscore sign escaped",
			input:    "test_product",
			expected: `test\_product`,
		},
		{
			name:     "Backslash escaped",
			input:    `backslash\test`,
			expected: `backslash\\test`,
		},
		{
			name:     "Combined characters",
			input:    `100%_back\slash`,
			expected: `100\%\_back\\slash`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SanitizeSQLWildcards(tt.input)
			if got != tt.expected {
				t.Errorf("SanitizeSQLWildcards(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
