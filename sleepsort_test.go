package sleepsort

import "testing"

func TestIntSlice(t *testing.T) {
	// Prepare input
	input := []int{8, 1, 4, 2, 3, 3, 14}

	// Prepare output to check against
	output := []int{1, 2, 3, 3, 4, 8, 14}

	// Sort input with default item duration.
	sorted := IntSlice(input, 0)

	// Check sorted against output
	for i := range sorted {
		if sorted[i] != output[i] {
			t.Errorf("Invalid item #%d: is %d, should be %d", i, sorted[i], output[i])
		}
	}
}

func TestFloatSlice(t *testing.T) {
	// Prepare input
	input := []float64{0.8412, 0.08, 0.482, 0.20001, 0.3, 0.3001, 1.4}

	// Prepare output to check against
	output := []float64{0.08, 0.20001, 0.3, 0.3001, 0.482, 0.8412, 1.4}

	// Sort input with default item duration.
	sorted := FloatSlice(input, 0)

	// Check sorted against output
	for i := range sorted {
		if sorted[i] != output[i] {
			t.Errorf("Invalid item #%d: is %f, should be %f", i, sorted[i], output[i])
		}
	}
}
