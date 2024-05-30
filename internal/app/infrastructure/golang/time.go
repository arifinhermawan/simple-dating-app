package golang

import "time"

func (g *Golang) GetTimeNow() time.Time {
	return time.Now()
}

func (g *Golang) GetMidnight(input time.Time) time.Time {
	result := time.Date(input.Year(), input.Month(), input.Day(), 0, 0, 0, 0, input.Location())

	return result
}
