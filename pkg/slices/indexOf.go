package slices

func IndexOf[V comparable](s []V, n V) int {
	for i, v := range s {
		if v == n {
			return i
		}
	}
	return -1
}
