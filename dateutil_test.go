package dateutil

import (
	"testing"
	"time"
)

func TestRrule(t *testing.T) {
	tests := []struct {
		name string
		rule *RRule
		want []time.Time
	}{
		{
			"Yearly",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 9, 0),
				datetime(1998, 9, 2, 9, 0),
				datetime(1999, 9, 2, 9, 0),
			},
		},
		{
			"Yearly interval",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				Interval:  2,
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 9, 0),
				datetime(1999, 9, 2, 9, 0),
				datetime(2001, 9, 2, 9, 0),
			},
		},
		{
			"Yearly interval large",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				Interval:  100,
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 9, 0),
				datetime(2097, 9, 2, 9, 0),
				datetime(2197, 9, 2, 9, 0),
			},
		},
		{
			"Yearly by month",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByMonth:   []time.Month{time.January, time.March},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 1, 2, 9, 0),
				datetime(1998, 3, 2, 9, 0),
				datetime(1999, 1, 2, 9, 0),
			},
		},
		{
			"Yearly by month day",
			&RRule{
				Frequency:  FreqYearly,
				Count:      3,
				ByMonthDay: []int{1, 3},
				DtStart:    datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 3, 9, 0),
				datetime(1997, 10, 1, 9, 0),
				datetime(1997, 10, 3, 9, 0),
			},
		},
		{
			"Yearly by month and month day",
			&RRule{
				Frequency:  FreqYearly,
				Count:      3,
				ByMonth:    []time.Month{time.January, time.March},
				ByMonthDay: []int{5, 7},
				DtStart:    datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 1, 5, 9, 0),
				datetime(1998, 1, 7, 9, 0),
				datetime(1998, 3, 5, 9, 0),
			},
		},
		{
			"Yearly by weekday",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByWeekday: []string{"TU", "TH"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 9, 0),
				datetime(1997, 9, 4, 9, 0),
				datetime(1997, 9, 9, 9, 0),
			},
		},
		{
			"Yearly by N weekday",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByWeekday: []string{"1TU", "-1TH"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 12, 25, 9, 0),
				datetime(1998, 1, 6, 9, 0),
				datetime(1998, 12, 31, 9, 0),
			},
		},
		{
			"Yearly by N weekday large",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByWeekday: []string{"3TU", "-3TH"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 12, 11, 9, 0),
				datetime(1998, 1, 20, 9, 0),
				datetime(1998, 12, 17, 9, 0),
			},
		},
		{
			"Yearly by month and weekday",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByMonth:   time.Month{time.January, time.March},
				ByWeekday: []string{"TU", "TH"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 1, 1, 9, 0),
				datetime(1998, 1, 6, 9, 0),
				datetime(1998, 1, 8, 9, 0),
			},
		},
		{
			"Yearly by month and N weekday",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByMonth:   time.Month{time.January, time.March},
				ByWeekday: []string{"1TU", "-1TH"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 1, 6, 9, 0),
				datetime(1998, 1, 29, 9, 0),
				datetime(1998, 3, 3, 9, 0),
			},
		},
		{
			"Yearly by month and N weekday large",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByMonth:   time.Month{time.January, time.March},
				ByWeekday: []string{"3TU", "-3TH"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 1, 15, 9, 0),
				datetime(1998, 1, 20, 9, 0),
				datetime(1998, 3, 12, 9, 0),
			},
		},
		{
			"Yearly by month day and weekday",
			&RRule{
				Frequency:  FreqYearly,
				Count:      3,
				ByMonthDay: []int{1, 3},
				ByWeekday:  []string{"TU", "TH"},
				DtStart:    datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 1, 1, 9, 0),
				datetime(1998, 2, 3, 9, 0),
				datetime(1998, 3, 3, 9, 0),
			},
		},
		{
			"YearlyByMonthAndMonthDayAndWeekDay",
			&RRule{
				Frequency:  FreqYearly,
				Count:      3,
				ByMonth:    []time.Month{time.January, time.March},
				ByMonthDay: []int{1, 3},
				ByWeekday:  []string{"TU", "TH"},
				DtStart:    datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 1, 1, 9, 0),
				datetime(1998, 3, 3, 9, 0),
				datetime(2001, 3, 1, 9, 0),
			},
		},
		{
			"YearlyByYearDay",
			&RRule{
				Frequency: FreqYearly,
				Count:     4,
				ByYearDay: []int{1, 100, 200, 365},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 12, 31, 9, 0),
				datetime(1998, 1, 1, 9, 0),
				datetime(1998, 4, 10, 9, 0),
				datetime(1998, 7, 19, 9, 0),
			},
		},
		{
			"YearlyByYearDayNeg",
			&RRule{
				Frequency: FreqYearly,
				Count:     4,
				ByYearDay: []int{-365, -266, -166, -1},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 12, 31, 9, 0),
				datetime(1998, 1, 1, 9, 0),
				datetime(1998, 4, 10, 9, 0),
				datetime(1998, 7, 19, 9, 0),
			},
		},
		{
			"YearlyByMonthAndYearDay",
			&RRule{
				Frequency: FreqYearly,
				Count:     4,
				ByMonth:   []time.Month{time.April, time.July},
				ByYearDay: []int{1, 100, 200, 365},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 4, 10, 9, 0),
				datetime(1998, 7, 19, 9, 0),
				datetime(1999, 4, 10, 9, 0),
				datetime(1999, 7, 19, 9, 0),
			},
		},
		{
			"YearlyByMonthAndYearDayNeg",
			&RRule{
				Frequency: FreqYearly,
				Count:     4,
				ByMonth:   []time.Month{time.April, time.July},
				ByYearDay: []int{-365, -266, -166, -1},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 4, 10, 9, 0),
				datetime(1998, 7, 19, 9, 0),
				datetime(1999, 4, 10, 9, 0),
				datetime(1999, 7, 19, 9, 0),
			},
		},
		{
			"YearlyByWeekNo",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByWeekNo:  []int{20},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 5, 11, 9, 0),
				datetime(1998, 5, 12, 9, 0),
				datetime(1998, 5, 13, 9, 0),
			},
		},
		{
			"YearlyByWeekNoAndWeekDay",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByWeekNo:  []int{1},
				ByWeekday: []string{"MO"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 12, 29, 9, 0),
				datetime(1999, 1, 4, 9, 0),
				datetime(2000, 1, 3, 9, 0),
			},
		},
		{
			"YearlyByWeekNoAndWeekDayLarge",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByWeekNo:  []int{52},
				ByWeekday: []string{"SU"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 12, 28, 9, 0),
				datetime(1998, 12, 27, 9, 0),
				datetime(2000, 1, 2, 9, 0),
			},
		},
		{
			"YearlyByWeekNoAndWeekDayLast",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByWeekNo:  []int{-1},
				ByWeekday: []string{"SU"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 12, 28, 9, 0),
				datetime(1999, 1, 3, 9, 0),
				datetime(2000, 1, 2, 9, 0),
			},
		},
		// {
		// 	"YearlyByEaster",
		// 	&RRule{
		// 		Frequency: FreqYearly,
		// 		Count:     3,
		// 		ByEaster:  0,
		// 		DtStart:   datetime(1997, 9, 2, 9, 0),
		// 	},
		// 	[]time.Time{
		// 		datetime(1998, 4, 12, 9, 0),
		// 		datetime(1999, 4, 4, 9, 0),
		// 		datetime(2000, 4, 23, 9, 0),
		// 	},
		// },
		// {
		// 	"YearlyByEasterPos",
		// 	&RRule{
		// 		Frequency: FreqYearly,
		// 		Count:     3,
		// 		ByEaster:  1,
		// 		DtStart:   datetime(1997, 9, 2, 9, 0),
		// 	},
		// 	[]time.Time{
		// 		datetime(1998, 4, 13, 9, 0),
		// 		datetime(1999, 4, 5, 9, 0),
		// 		datetime(2000, 4, 24, 9, 0),
		// 	},
		// },
		// {
		// 	"testYearlyByEasterNeg",
		// 	&RRule{
		// 		Frequency: FreqYearly,
		// 		Count:     3,
		// 		ByEaster:  -1,
		// 		DtStart:   datetime(1997, 9, 2, 9, 0),
		// 	},
		// 	[]time.Time{
		// 		datetime(1998, 4, 11, 9, 0),
		// 		datetime(1999, 4, 3, 9, 0),
		// 		datetime(2000, 4, 22, 9, 0),
		// 	},
		// },
		{
			"YearlyByWeekNoAndWeekDay53",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByWeekNo:  []int{53},
				ByWeekday: []string{"MO"},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1998, 12, 28, 9, 0),
				datetime(2004, 12, 27, 9, 0),
				datetime(2009, 12, 28, 9, 0),
			},
		},
		{
			"YearlyByHour",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByHour:    []int{6, 18},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 18, 0),
				datetime(1998, 9, 2, 6, 0),
				datetime(1998, 9, 2, 18, 0),
			},
		},
		{
			"YearlyByMinute",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByMinute:  []int{6, 18},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 9, 6),
				datetime(1997, 9, 2, 9, 18),
				datetime(1998, 9, 2, 9, 6),
			},
		},
		{
			"YearlyBySecond",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				BySecond:  []int{6, 18},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 9, 0).Add(6 * time.Second),
				datetime(1997, 9, 2, 9, 0).Add(18 * time.Second),
				datetime(1998, 9, 2, 9, 0).Add(6 * time.Second),
			},
		},
		{
			"YearlyByHourAndMinute",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByHour:    []int{6, 18},
				ByMinute:  []int{6, 18},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 18, 6),
				datetime(1997, 9, 2, 18, 18),
				datetime(1998, 9, 2, 6, 6),
			},
		},
		{
			"YearlyByHourAndSecond",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByHour:    []int{6, 18},
				BySecond:  []int{6, 18},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 18, 0).Add(6 * time.Second),
				datetime(1997, 9, 2, 18, 0).Add(18 * time.Second),
				datetime(1998, 9, 2, 6, 0).Add(6 * time.Second),
			},
		},
		{
			"YearlyByMinuteAndSecond",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByMinute:  []int{6, 18},
				BySecond:  []int{6, 18},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 9, 6).Add(6 * time.Second),
				datetime(1997, 9, 2, 9, 6).Add(18 * time.Second),
				datetime(1998, 9, 2, 6, 18).Add(6 * time.Second),
			},
		},
		{
			"YearlyByHourAndMinuteAndSecond",
			&RRule{
				Frequency: FreqYearly,
				Count:     3,
				ByHour:    []int{6, 18},
				ByMinute:  []int{6, 18},
				BySecond:  []int{6, 18},
				DtStart:   datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 9, 2, 18, 6).Add(6 * time.Second),
				datetime(1997, 9, 2, 18, 6).Add(18 * time.Second),
				datetime(1998, 9, 2, 18, 18).Add(6 * time.Second),
			},
		},
		{
			"YearlyBySetPos",
			&RRule{
				Frequency:  FreqYearly,
				Count:      3,
				ByMonthDay: []int{15},
				ByHour:     []int{6, 18},
				BySetPos:   []int{3, -3},
				DtStart:    datetime(1997, 9, 2, 9, 0),
			},
			[]time.Time{
				datetime(1997, 11, 15, 18, 0),
				datetime(1998, 2, 15, 6, 0),
				datetime(1998, 11, 15, 18, 0),
			},
		},
	}
}

func datetime(year, month, day, hour, minute) time.Time {
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
}
