package dateutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNowDate(t *testing.T) {
	n := Now()
	e := time.Now()

	assert.Equal(t, e, n)

}
func TestStrToDate(t *testing.T) {

	e := "2020-10-10"
	d, _ := StrToDate(e, "2006-01-02")
	n := Format("2006-01-02", d)

	assert.Equal(t, e, n)

}
func TestIsoDate(t *testing.T) {

	i := "2020-10-10"
	e := "2020-10-10T00:00:00Z"
	d, _ := StrToDate(i, "2006-01-02")
	n := ToISODate(d)
	assert.Equal(t, e, n)

}
func TestBeginingOfDate(t *testing.T) {

	i := "2020-10-10"

	d, _ := StrToDate(i, "2006-01-02")

	e := time.Date(2020, 10, 10, 0, 0, 0, 0, d.Location())

	n := BeginingOfDay(d)

	assert.Equal(t, e, n)

}

func TestEndOfDate(t *testing.T) {

	i := "2020-10-10"

	d, _ := StrToDate(i, "2006-01-02")

	e := time.Date(2020, 10, 10, 23, 59, 59, 0, d.Location())

	n := EndOfDay(d)

	assert.Equal(t, e, n)

}
func TestPointerDate(t *testing.T) {

	n := ToPtr(Now())
	e := time.Now()

	assert.IsType(t, &e, n)

}
func TestDatePointerToDate(t *testing.T) {

	n := ToPtr(Now())
	v := PtrToTime(n)
	e := time.Now()

	assert.IsType(t, e, v)

}
func TestToMili(t *testing.T) {
	i := "2020-10-10"
	d, _ := StrToDate(i, "2006-01-02")
	e := int64(1602288000000)
	v := TimeUnixMilli(d)
	assert.Equal(t, e, v)

}

func TestMiliToTime(t *testing.T) {
	d := "2020-10-10"
	e, _ := StrToDate(d, "2006-01-02")
	i := int64(1602288000000)
	v := MillisecondsTimeValue(i)
	assert.Equal(t, e, v)
}

func TestToSecond(t *testing.T) {
	d := "2020-10-10"
	i, _ := StrToDate(d, "2006-01-02")

	v := ToSecond(i)
	e := int64(1602288000)
	assert.Equal(t, e, v)
}

func TestSecondToTime(t *testing.T) {
	d := "2020-10-10"
	e, _ := StrToDate(d, "2006-01-02")

	i := int64(1602288000)

	v := SecondsTime(i)
	assert.Equal(t, e, v)
}
func TestDatediff(t *testing.T) {
	i1 := "2020-10-10"
	i2 := "2020-10-11"
	d1, _ := StrToDate(i1, "2006-01-02")
	d2, _ := StrToDate(i2, "2006-01-02")

	v := DateDiff(d1, d2)

	e, _ := time.ParseDuration("24h0m0s")

	assert.Equal(t, e, v)
}
func TestDatediffSecond(t *testing.T) {
	i1 := "2020-10-10"
	i2 := "2020-10-11"
	d1, _ := StrToDate(i1, "2006-01-02")
	d2, _ := StrToDate(i2, "2006-01-02")

	v := DateDiffSeconds(d1, d2)

	e := float64(86400)

	assert.Equal(t, e, v)
}
func TestDatediffMinute(t *testing.T) {
	i1 := "2020-10-10"
	i2 := "2020-10-11"
	d1, _ := StrToDate(i1, "2006-01-02")
	d2, _ := StrToDate(i2, "2006-01-02")

	v := DateDiffMinutes(d1, d2)

	e := float64(1440)

	assert.Equal(t, e, v)
}
func TestDatediffHours(t *testing.T) {
	i1 := "2020-10-10"
	i2 := "2020-10-11"
	d1, _ := StrToDate(i1, "2006-01-02")
	d2, _ := StrToDate(i2, "2006-01-02")

	v := DateDiffHours(d1, d2)

	e := float64(24)

	assert.Equal(t, e, v)
}
func TestDatediffDays(t *testing.T) {
	i1 := "2020-10-10"
	i2 := "2020-10-11"
	d1, _ := StrToDate(i1, "2006-01-02")
	d2, _ := StrToDate(i2, "2006-01-02")

	v := DateDiffDays(d1, d2)

	e := float64(1)

	assert.Equal(t, e, v)
}
func TestToUtc(t *testing.T) {
	i := "2020-10-10"
	d, _ := StrToDate(i, "2006-01-02")

	v := ToUTC(d)

	e := d.UTC()

	assert.Equal(t, e, v)
}

func TestGetLocation(t *testing.T) {
	i := "2020-10-10"
	d, _ := StrToDate(i, "2006-01-02")

	v := GetLocation(d)

	e := d.Location()

	assert.Equal(t, e, v)
}

func TestDateEqual(t *testing.T) {
	i1 := "2020-10-11"
	i2 := "2020-10-11"
	d1, _ := StrToDate(i1, "2006-01-02")
	d2, _ := StrToDate(i2, "2006-01-02")

	v := DateEqual(d1, d2)


	e :=true

	assert.Equal(t, e, v)
}

func TestEqual(t *testing.T) {
	i1 := "2020-10-11"
	i2 := "2020-10-11"
	d1, _ := StrToDate(i1, "2006-01-02")
	d2, _ := StrToDate(i2, "2006-01-02")

	v := DateEqual(d1, d2)

	t.Log(v)
	e :=true

	assert.Equal(t, e, v)
}

func TestAddNextDay(t *testing.T) {
	i := "2020-10-10"
	d, _ := StrToDate(i, "2006-01-02")
	i2 := "2020-10-11"
	e, _ := StrToDate(i2, "2006-01-02")
 
	v := AddNextDay(d) 

	assert.Equal(t, e, v)
} 