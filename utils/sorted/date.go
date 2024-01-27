package sorted

import "time"

type DateFunc func(i, j time.Time) bool

var Date DateFunc = DateDesc

func DateDesc(i, j time.Time) bool {
	return i.After(j)
}

func DateAsc(i, j time.Time) bool {
	return i.Before(j)
}
