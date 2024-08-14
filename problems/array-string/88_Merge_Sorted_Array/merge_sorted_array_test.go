package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(T *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}

	var tests = []struct {
		name     string
		args     *args
		expected []int
	}{
		{
			name: "Example 1",
			args: &args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "Example 2",
			args: &args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			expected: []int{1},
		},
		{
			name: "Example 3",
			args: &args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			expected: []int{1},
		},
	}

	for _, test := range tests {
		T.Run(test.name, func(T *testing.T) {
			Merge(test.args.nums1, test.args.m, test.args.nums2, test.args.n)
			assert.Equal(T, test.expected, test.args.nums1)
		})
	}
}
