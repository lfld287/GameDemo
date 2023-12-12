package board

import "testing"

type testSample struct {
	original   []int
	expected   []int
	headToTail bool
}

var testList = []testSample{
	{
		original:   []int{2, 2, 8, 4},
		expected:   []int{4, 8, 4, 0},
		headToTail: false,
	},
	{
		original:   []int{2, 2, 8, 4},
		expected:   []int{0, 4, 8, 4},
		headToTail: true,
	},
	{
		original:   []int{0, 0, 0, 0},
		expected:   []int{0, 0, 0, 0},
		headToTail: false,
	},
	{
		original:   []int{0, 0, 0, 0},
		expected:   []int{0, 0, 0, 0},
		headToTail: true,
	},
	{
		original:   []int{2, 0, 0, 2},
		expected:   []int{0, 0, 0, 4},
		headToTail: true,
	},
	{
		original:   []int{2, 2, 8, 4},
		expected:   []int{4, 8, 4, 0},
		headToTail: false,
	},
}

func TestProcessLine(t *testing.T) {
	for _, sample := range testList {
		processed := processLine(sample.original, sample.headToTail)
		if len(processed) != len(sample.expected) {
			t.Errorf("expected %v, but got %v", sample.expected, processed)
		}
		for i := 0; i < len(sample.expected); i++ {
			if sample.expected[i] != processed[i] {
				t.Errorf("expected %v, but got %v", sample.expected, processed)
			}
		}
	}

}
