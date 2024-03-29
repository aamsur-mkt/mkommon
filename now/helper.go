package now

import "time"

func New(t time.Time) *Now {
	return &Now{t}
}

func NewParse(layout string, value string) *Now {
	if t, e := time.Parse(layout, value); e == nil {
		t = t.In(time.Local)
		return &Now{t}
	}

	return &Now{time.Now()}
}

func BeginningOfMinute() time.Time {
	return New(time.Now()).BeginningOfMinute()
}

func BeginningOfHour() time.Time {
	return New(time.Now()).BeginningOfHour()
}

func BeginningOfDay() time.Time {
	return New(time.Now()).BeginningOfDay()
}

func BeginningOfWeek() time.Time {
	return New(time.Now()).BeginningOfWeek()
}

func BeginningOfMonth() time.Time {
	return New(time.Now()).BeginningOfMonth()
}

func BeginningOfQuarter() time.Time {
	return New(time.Now()).BeginningOfQuarter()
}

func BeginningOfYear() time.Time {
	return New(time.Now()).BeginningOfYear()
}

func EndOfMinute() time.Time {
	return New(time.Now()).EndOfMinute()
}

func EndOfHour() time.Time {
	return New(time.Now()).EndOfHour()
}

func EndOfDay() time.Time {
	return New(time.Now()).EndOfDay()
}

func EndOfWeek() time.Time {
	return New(time.Now()).EndOfWeek()
}

func EndOfMonth() time.Time {
	return New(time.Now()).EndOfMonth()
}

func EndOfQuarter() time.Time {
	return New(time.Now()).EndOfQuarter()
}

func EndOfYear() time.Time {
	return New(time.Now()).EndOfYear()
}

func Monday() time.Time {
	return New(time.Now()).Monday()
}

func Sunday() time.Time {
	return New(time.Now()).Sunday()
}

func EndOfSunday() time.Time {
	return New(time.Now()).EndOfSunday()
}

func Parse(strs ...string) (time.Time, error) {
	return New(time.Now()).Parse(strs...)
}

func MustParse(strs ...string) time.Time {
	return New(time.Now()).MustParse(strs...)
}

func Between(time1, time2 string) bool {
	return New(time.Now()).Between(time1, time2)
}
