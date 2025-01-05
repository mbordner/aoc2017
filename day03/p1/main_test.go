package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetDistance(t *testing.T) {
	tests := []struct {
		num      int
		expected int
	}{
		{
			num:      1,
			expected: 0,
		},
		{
			num:      2,
			expected: 1,
		},
		{
			num:      3,
			expected: 2,
		},
		{
			num:      15,
			expected: 2,
		},
		{
			num:      16,
			expected: 3,
		},
		{
			num:      17,
			expected: 4,
		},
		{
			num:      18,
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.num), func(t *testing.T) {
			d := getDistanceFromCenter(test.num)
			assert.Equal(t, test.expected, d)
		})
	}
}

func Test_GetLevel(t *testing.T) {
	tests := []struct {
		num           int
		expectedLevel int
		expectedCount int
		expectedMin   int
		expectedMax   int
	}{
		{
			num:           1,
			expectedLevel: 0,
			expectedCount: 1,
			expectedMin:   1,
			expectedMax:   1,
		},
		{
			num:           2,
			expectedLevel: 1,
			expectedCount: 8,
			expectedMin:   2,
			expectedMax:   9,
		},
		{
			num:           6,
			expectedLevel: 1,
			expectedCount: 8,
			expectedMin:   2,
			expectedMax:   9,
		},
		{
			num:           12,
			expectedLevel: 2,
			expectedCount: 16,
			expectedMin:   10,
			expectedMax:   25,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.num), func(t *testing.T) {

			level, count, min, max := getLevel(test.num)
			assert.Equal(t, test.expectedLevel, level)
			assert.Equal(t, test.expectedCount, count)
			assert.Equal(t, test.expectedMin, min)
			assert.Equal(t, test.expectedMax, max)

		})
	}
}
