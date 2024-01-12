package utils

func InSlice(str string, slice []string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}

	return false
}

func InSliceInt(digit int, slice []int) bool {
	for _, v := range slice {
		if v == digit {
			return true
		}
	}

	return false
}
