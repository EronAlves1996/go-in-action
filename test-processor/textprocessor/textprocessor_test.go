package textprocessor

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "two-word sentence",
			input:    "Bad call",
			expected: "call Bad",
		},
		{
			name:     "senetence with punctuation",
			input:    "Sun is great! Right!",
			expected: "Right! great! is Sun",
		},
		{
			name:     "single word",
			input:    "john",
			expected: "john",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "String with multiple spaces between words",
			input:    "Sorry,   I can make   this",
			expected: "this make can I Sorry,",
		},
	}

	for _, test := range tests {
		d := MyProcessor{}
		t.Run(test.name, func(t *testing.T) {
			n := d.ReverseWords(test.input)
			if n != test.expected {
				t.Errorf("[%s] For '%s', Expected '%s', got '%s'", test.name, test.input, test.expected, n)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	d := MyProcessor{}
	c := d.CountVowels("blablabla")
	if c != 3 {
		t.Errorf("For '%s', Expected '%d', got '%d'", "blablabla", 3, c)
	}
}

func BenchmarkReverseWords(b *testing.B) {
	p := MyProcessor{}
	s := "the quick brown fox jumps over the lazy dog"
	for b.Loop() {
		p.ReverseWords(s)
	}
}
