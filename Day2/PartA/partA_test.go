package day2partA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intCode(t *testing.T) {

	var tests = []struct {
		title string
		input []int
		out   []int
	}{
		{
			"a",
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		},
		{
			"b",
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		},
		{
			"c",
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			"d",
			[]int{1,1,1,4,99,5,6,0,99},
			[]int{30,1,1,4,2,5,6,0,99},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			out := intCode(test.input)
			assert.Equal(t, test.out, out)
		})
	}

}
