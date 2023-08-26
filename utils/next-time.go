package utils

import "time"

func NextYear(mo time.Month, d, h, m int) func(t time.Time) time.Time {
	return func(t time.Time) time.Time {
		a := time.Date(t.Year(), mo, d, h, m, 0, 0, time.UTC)
		if a.After(t) {
			return a
		}
		return time.Date(t.Year()+1, mo, d, h, m, 0, 0, time.UTC)
	}
}
