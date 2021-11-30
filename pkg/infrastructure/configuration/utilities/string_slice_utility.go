package utilities
func ExistsInSlice(a []string, b string) bool {
	if (a == nil) || len(b) == 0 {
		return false
	}

	for i := range a {
		if a[i] == b {
			return true
		}
	}

	return false
}

