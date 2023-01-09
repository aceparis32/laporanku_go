package controllers

func Contains(a []string, source string) bool {
	for _, n := range a {
		if source == n {
			return true
		}
	}

	return false
}

func Find(a []string, source string) int {
	for i, n := range a {
		if source == n {
			return i
		}
	}

	return len(a)
}
