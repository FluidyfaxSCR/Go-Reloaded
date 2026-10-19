package main

import "testing"

func TestTransformLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "number conversions",
			input: "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			want:  "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:  "uppercase and punctuation",
			input: "Ready, set, go (up) !",
			want:  "Ready, set, GO!",
		},
		{
			name:  "lowercase",
			input: "I should stop SHOUTING (low)",
			want:  "I should stop shouting",
		},
		{
			name:  "capitalize",
			input: "Welcome to the Brooklyn bridge (cap)",
			want:  "Welcome to the Brooklyn Bridge",
		},
		{
			name:  "counted case modifier",
			input: "this is so exciting (up, 2)",
			want:  "this is SO EXCITING",
		},
		{
			name:  "counted case modifier without comma",
			input: "this is so exciting (up 2)",
			want:  "this is SO EXCITING",
		},
		{
			name:  "punctuation spacing",
			input: "over there , and then BAMM !!",
			want:  "over there, and then BAMM!!",
		},
		{
			name:  "punctuation groups",
			input: "I was thinking ... You were right",
			want:  "I was thinking... You were right",
		},
		{
			name:  "quoted phrase spacing",
			input: "describe me: ' awesome '",
			want:  "describe me: 'awesome'",
		},
		{
			name:  "possessive apostrophe spacing",
			input: "James' book",
			want:  "James' book",
		},
		{
			name:  "article before vowel",
			input: "I've seen a elephant",
			want:  "I've seen an elephant",
		},
		{
			name:  "article conversion after case modifier",
			input: "a elephant (up)",
			want:  "an ELEPHANT",
		},
		{
			name:  "hex alone",
			input: "1E (hex) files were added",
			want:  "30 files were added",
		},
		{
			name:  "bin alone",
			input: "It has been 10 (bin) years",
			want:  "It has been 2 years",
		},
		{
			name:  "capital article",
			input: "There it was. A amazing rock!",
			want:  "There it was. An amazing rock!",
		},
		{
			name:  "article before h",
			input: "a honest man",
			want:  "an honest man",
		},
		{
			name:  "tag followed by comma",
			input: "it was the best of times, it was the worst of times (up) , it was",
			want:  "it was the best of times, it was the worst of TIMES, it was",
		},
		{
			name:  "tag glued to comma",
			input: "it was the age of foolishness (cap, 6) , it was",
			want:  "It Was The Age Of Foolishness, it was",
		},
		{
			name:  "count larger than words",
			input: "hello world (up, 6)",
			want:  "HELLO WORLD",
		},
		{
			name:  "multi word quote",
			input: "As Elton John said: ' I am the most well-known homosexual in the world '",
			want:  "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name:  "leading punctuation glued to next word",
			input: "over there ,and then",
			want:  "over there, and then",
		},
		{
			name:  "interrogation group",
			input: "really !? yes",
			want:  "really!? yes",
		},
		{
			name:  "negative modifier count",
			input: "alpha bravo charlie (up -1) !",
			want:  "[alpha bravo charlie (up -1) !]=err syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := transformLine(test.input); got != test.want {
				t.Errorf("transformLine(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}

func TestTransformTextPreservesNewlines(t *testing.T) {
	input := "first line (up, 2)\nsecond line (low)\n"
	want := "FIRST LINE\nsecond line\n"
	if got := transformText(input); got != want {
		t.Errorf("transformText(%q) = %q, want %q", input, got, want)
	}
}
