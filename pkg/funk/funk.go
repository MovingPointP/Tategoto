package funk

func Map[T, T2 any](array []T, f func(T) T2) (result []T2) {
	for _, value := range array {
		result = append(result, f(value))
	}
	return
}
