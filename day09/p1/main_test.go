package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StreamGroupScore(t *testing.T) {
	tests := []struct {
		stream   string
		expected int
	}{
		{
			stream:   `{}`,
			expected: 1,
		},
		{
			stream:   `{{{}}}`,
			expected: 6,
		},
		{
			stream:   `{{},{}}`,
			expected: 5,
		},
		{
			stream:   `{{{},{},{{}}}}`,
			expected: 16,
		},
		{
			stream:   `{<a>,<a>,<a>,<a>}`,
			expected: 1,
		},
		{
			stream:   `{{<ab>},{<ab>},{<ab>},{<ab>}}`,
			expected: 9,
		},
		{
			stream:   `{{<!!>},{<!!>},{<!!>},{<!!>}}`,
			expected: 9,
		},
		{
			stream:   `{{<a!>},{<a!>},{<a!>},{<ab>}}`,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.stream, func(t *testing.T) {
			score := getStreamGroupScore([]byte(tt.stream))
			assert.Equal(t, tt.expected, score)
		})
	}
}
