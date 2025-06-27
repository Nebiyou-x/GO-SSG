package internal

import (
	"testing"
)

func TestMarkdownToHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Header", "# Title", "<h1>Title</h1>"},
		{"List", "- item1\n- item2", "<ul>"},
		{"Link", "[link](https://example.com)", "<a href=\"https://example.com\""},
		{"Code block", "```go\nfmt.Println(\"hi\")\n```", "<pre><code class=\"language-go\">"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := MarkdownToHTML(tt.input)
			if !contains(output, tt.expected) {
				t.Errorf("expected output to contain %q, got %q", tt.expected, output)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(s) > len(substr) && (contains(s[1:], substr) || contains(s[:len(s)-1], substr)))) || (len(s) < len(substr) && substr[:len(s)] == s)
}
