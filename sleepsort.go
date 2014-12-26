package sleepsort

import "time"

// IntSlice sorts an integer slice. Sorting takes max(input)*itemDuration.
// itemDuration is 100ms by default.
func IntSlice(input []int, itemDuration time.Duration) []int {
	// Set the default duration if itemDuration is empty
	if itemDuration == 0 {
		itemDuration = time.Millisecond * 100
	}

	// Prepare a result slice
	var result []int

	// Set up a channel for responses
	channel := make(chan int, len(input))

	// Start the goroutines
	for i := 0; i < len(input); i++ {
		go func(i int) {
			time.Sleep(itemDuration * time.Duration(input[i]))
			channel <- input[i]
		}(i)
	}

	// Collect the result
	for i := 0; i < len(input); i++ {
		result = append(result, <-channel)
	}

	// Return the result
	return result
}

// FloatSlice sorts a float64 slice. Sorting takes max(input)*itemDuration.
// itemDuration is 1s by default.
func FloatSlice(input []float64, itemDuration time.Duration) []float64 {
	// Set the default duration if itemDuration is empty
	if itemDuration == 0 {
		itemDuration = time.Millisecond * 1000
	}

	// Prepare a result slice
	var result []float64

	// Set up a channel for responses
	channel := make(chan float64, len(input))

	// Start the goroutines
	for i := 0; i < len(input); i++ {
		go func(i int) {
			time.Sleep(time.Duration(float64(itemDuration) * input[i]))
			channel <- input[i]
		}(i)
	}

	// Collect the result
	for i := 0; i < len(input); i++ {
		result = append(result, <-channel)
	}

	// Return the result
	return result
}
