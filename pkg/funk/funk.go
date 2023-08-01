package funk

import "strconv"

func Map[T, T2 any](array []T, f func(T) T2) (result []T2) {
	for _, value := range array {
		result = append(result, f(value))
	}
	return
}

func StringToUint(str string) (uint, error) {
	strInt64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(strInt64), nil
}

func UintToString(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}
