package funk

import "time"

func Map[T, T2 any](array []T, f func(T) T2) (result []T2) {
	for _, value := range array {
		result = append(result, f(value))
	}
	return
}

func CompareAboutTime(t1 time.Time, t2 time.Time) bool {
	return t1.Sub(t2).Seconds() < 1
}
