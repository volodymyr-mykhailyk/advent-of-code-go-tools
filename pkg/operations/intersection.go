package operations

func Intersection[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]int)

	for _, v := range a {
		hash[v] = 1
	}

	for _, v := range b {
		if hash[v] == 1 {
			hash[v] += 1
			set = append(set, v)
		}
	}

	return set
}
