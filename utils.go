package utils

func Contains(str string, slice []string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}

	return false
}
