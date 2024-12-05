package fp

func Map[T any, U any](s []T, mapperFunc func(T) U) []U {
	result := make([]U, len(s))

	for i := range s {
		result[i] = mapperFunc(s[i])
	}

	return result
}
