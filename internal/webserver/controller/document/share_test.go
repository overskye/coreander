package document

import (
	"strings"
	"testing"
)

func TestTruncateShareCommentCountsUnicodeCharacters(t *testing.T) {
	const maxCharacters = 10
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "cyrillic below limit", input: strings.Repeat("Ж", 6), want: strings.Repeat("Ж", 6)},
		{name: "emoji below limit", input: strings.Repeat("🙂", 4), want: strings.Repeat("🙂", 4)},
		{name: "CJK below limit", input: strings.Repeat("界", 4), want: strings.Repeat("界", 4)},
		{name: "cyrillic above limit", input: strings.Repeat("Ж", 12), want: strings.Repeat("Ж", maxCharacters)},
		{name: "emoji above limit", input: strings.Repeat("🙂", 12), want: strings.Repeat("🙂", maxCharacters)},
		{name: "CJK above limit", input: strings.Repeat("界", 12), want: strings.Repeat("界", maxCharacters)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := truncateShareComment(tc.input, maxCharacters); got != tc.want {
				t.Errorf("truncateShareComment() = %q, want %q", got, tc.want)
			}
		})
	}
}
