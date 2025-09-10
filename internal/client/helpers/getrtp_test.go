package helpers

import (
	"testing"
)

func TestGetRTP(t *testing.T) {
	tests := []struct {
		sequence      []float64
		multiplicator []float64
		expected      float64
	}{
		{
			sequence:      []float64{0.5, 1.5, 2.0},
			multiplicator: []float64{1.0, 1.0, 2.5},
			expected:      (0.5 + 0 + 2.0) / 3, // 0.8333
		},
		{
			sequence:      []float64{1.0, 1.0, 1.0},
			multiplicator: []float64{0.5, 1.5, 0.9},
			expected:      (0.0 + 1.0 + 0.0) / 3.0, // 0.3333
		},
		{
			sequence:      []float64{0.0, 0.0},
			multiplicator: []float64{0.0, 0.0},
			expected:      0,
		},
		{
			sequence:      []float64{2.0, 3.0, 100.0, 1.1, 1.1},
			multiplicator: []float64{1.8, 965.0, 1.0, 5.6, 1.2},
			expected:      1.0399999999999998,
		},
	}

	for i, test := range tests {
		got := GetRTP(test.sequence, test.multiplicator)
		if got != test.expected {
			t.Errorf("Test %d failed: got %f, expected %f", i, got, test.expected)
		}
	}
}
