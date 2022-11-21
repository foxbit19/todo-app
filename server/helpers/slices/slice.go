package slices

// Filter use generics to get a typed slice to filter.
// It uses a func as argument to filter: if the func returns true, the
// element of the slice will be in the result slice.
func Filter[T any](slice *[]T, filterFunc func(i int) bool) *[]T  {
	var result []T

	for i := 0; i < len(*slice); i++ {
		if filterFunc(i) {
			result = append(result, (*slice)[i])
		}
	}

	return &result
}