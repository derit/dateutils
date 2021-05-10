package dateutils

import (
	"time"
)

func Now() time.Time {
	return time.Now()
}

func Format(v string, t time.Time) string {
	return t.Format(v)
}

func ToUnix(t time.Time) int64 {
	return t.Unix()
}

func ToISODate(t time.Time) string {
	return t.Format(time.RFC3339)
}

func BeginingOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())

}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

func ToPtr(v time.Time) *time.Time {
	return &v
}

func PtrToTime(v *time.Time) time.Time {
	if v != nil {
		return *v
	}
	return time.Time{}
}

func TimeUnixMilli(t time.Time) int64 {
	return t.Sub(time.Unix(0, 0)).Milliseconds()
}

func MillisecondsTimeValue(v int64) time.Time {
	return time.Unix(0, v*int64(time.Millisecond)).UTC()
}

func SecondsTime(v int64) time.Time {
	return time.Unix(v, 0).UTC()
}
func ToSecond(t time.Time) int64 {
	return t.UTC().Unix()
}

func DateDiff(start, end time.Time) time.Duration {
	return end.Sub(start)
}

func DateDiffSeconds(start, end time.Time) float64 {
	return end.Sub(start).Seconds()
}

func DateDiffMinutes(start, end time.Time) float64 {
	return end.Sub(start).Minutes()
}

func DateDiffHours(start, end time.Time) float64 {
	return end.Sub(start).Hours()
}

func DateDiffDays(start, end time.Time) float64 {
	return end.Sub(start).Hours() / 24
}

func GetTimeByLocation(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return time.Time{}, err
	}
	t = t.In(loc)
	return t, err
}

func ToUTC(t time.Time) time.Time {
	return t.UTC()
}

func GetLocation(t time.Time) *time.Location {
	return t.Location()
}

func DateEqual(d1, d2 time.Time) bool {
	yy1, mm1, dd1 := d1.Date()
	yy2, mm2, dd2 := d2.Date()
	return yy1 == yy2 && mm1 == mm2 && dd1 == dd2
}

func Equal(d1, d2 time.Time) bool {
	return d1.UnixNano() == d2.UnixNano()
}

func AddNextDay(now time.Time) time.Time {
	return now.AddDate(0, 0, 1)
}

func StrToDate(date, layout string) (time.Time, error) {
	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
