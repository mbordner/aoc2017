package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StreamGarbageLen(t *testing.T) {
	tests := []struct {
		stream   string
		expected int
	}{
		{
			stream:   `{<>}`,
			expected: 0,
		},
		{
			stream:   `{<random characters>}`,
			expected: 17,
		},
		{
			stream:   `{<<<<>}`,
			expected: 3,
		},
		{
			stream:   `{<{!>}>}`,
			expected: 2,
		},
		{
			stream:   `{<!!>}`,
			expected: 0,
		},
		{
			stream:   `{<!!!>>}`,
			expected: 0,
		},
		{
			stream:   `{<{o"i!a,<{i<a>}`,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.stream, func(t *testing.T) {
			root := getStreamRootGroup([]byte(tt.stream))
			assert.Equal(t, tt.expected, root.GarbageLen())
		})
	}
}
