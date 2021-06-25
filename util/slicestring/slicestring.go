package slicestring

func exists(ss []string, f func(string) bool) bool {
	for _, s := range ss {
		if f(s) {
			return true
		}
	}
	return false
}

func Contains(ss []string, keyword string) bool {
	return exists(
		ss,
		func(s string) bool {
			return s == keyword
		},
	)
}
