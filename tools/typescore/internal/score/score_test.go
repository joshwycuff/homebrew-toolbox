package score

import (
	"testing"
)

func TestScore(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		desc     string
	}{
		// Empty string
		{"", 0, "Empty input, score should be zero"},

		// Single character, base case
		{"a", 1, "Single home row letter 'a'"},
		{"A", 3, "Shifted home row letter 'A'"},

		{"rf", 3, "Different chars same finger (leftIndex)"},
		{"12", 9, "Both symbols on top row, different fingers"},

		// Repeating character
		{"aa", 3, "Repeating home row char - extra penalty"},
		{"AA", 7, "Repeating shifted char - penalty with shift multiplier"},

		// Shifted chars: increases score
		{"!", 8, "Shifted pinky char (!), top row, doubled for shift"},
		{"Q", 4, "Shifted top-row Q on left hand"},
		{"{", 4, "Shifted bracket (left pinky), top row, shift"},
		{"Z", 4, "Shifted bottom row char"},

		// Space (thumb)
		{" ", 1, "Single spacebar"},

		// Mixed easy and difficult chars
		{"asdf", 4, "Home row, all different fingers"},
		{"qaz", 8, "All pinky (different row), so pinky coefficient"},
		{"ZXCV", 13, "Bottom row left hand with shift"},

		// Pinkies
		{";;", 3, "Right pinky, repeated (home row)"},
		{"PP", 9, "Right pinky, shifted"},

		// Across both hands
		{"vbnm", 6, "Bottom row left to right hand, repeat index fingers"},

		// Punctuation
		{"'", 2, "Home row right pinky character"},
		{",", 1, "Right middle, bottom row"},
		{"__", 17, "Repeating shifted '_' symbol on pinky, top row"},

		// Sequential numbers
		{"1234567890", 45, "All digits, moving across top row and both hands"},
		{"!!!", 26, "Three exclamation, shifted top row pinky"},

		// Unknown/zero case
		{"©", 100, "Unknown character, fallback case"},

		// Mixed: includes shift, row, finger overlaps, repeats
		{"tt", 3, "Repeating same finger, index, top row"},
		{"ttt", 5, "Triplet repeat, index, top row"},
		{"gh", 2, "Adj. home row, different hands"},
		{"qazwsx", 13, "Alternating rows, using same fingers"},
		{"asdfASDF", 16, "Home row, lower and uppercase (shift applied on second half)"},

		// Top row, shifted and regular
		{"!@#$%^&*()", 69, "All top row, shifted symbols"},

		{"The quick brown fox jumps over the lazy dog", 53, "Typical sentence, all chars"},
		{"!qAz @wSx #eDc $rFv %tGb ^yHn *iK, (oL. )p:/", 160, "Terrible"},
	}

	for _, tt := range tests {
		got := Score(tt.input)
		if got != tt.expected {
			t.Errorf("Score(%q): got %d, want %d (%s)", tt.input, got, tt.expected, tt.desc)
		}
	}
}
