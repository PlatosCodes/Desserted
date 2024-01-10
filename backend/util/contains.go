package util

// contains checks if a slice contains an element
func Contains(slice []int64, element int64) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}
