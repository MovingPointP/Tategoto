package funk

import "time"

func Map[T, T2 any](array []T, f func(T) T2) (result []T2) {
	for _, value := range array {
		result = append(result, f(value))
	}
	return
}

func CompareAboutTime(t1 time.Time, t2 time.Time) bool {
	isYearEqual := t1.Year() == t2.Year()
	isMonthEqual := t1.Month() == t2.Month()
	isDayEqual := t1.Day() == t2.Day()
	isHourEqual := t1.Hour() == t2.Hour()
	isMinuteEqual := t1.Minute() == t2.Minute()
	return isYearEqual && isMonthEqual && isDayEqual && isHourEqual && isMinuteEqual
}
