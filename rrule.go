package dateutil

import (
	"time"
)

type Frequency string

const (
	Yearly Frequency = "YEARLY"
)

type Weekday string

const (
	Monday    Weekday = "MO"
	Tuesday           = "TU"
	Wednesday         = "WE"
	Thursday          = "TH"
	Friday            = "FR"
	Saturday          = "SA"
	Sunday            = "SU"
)

type RRule struct {
	Frequency  Frequency
	Count      int
	DtStart    time.Time
	Interval   int
	ByMonth    []time.Month
	ByMonthDay []int
	ByWeekday  []string
	ByYearDay  []int
	ByWeekNo   []int
}
