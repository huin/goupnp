package soap

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"
)

func MarshalFixed14_4(v float64) (string, error) {
	if v >= 1e14 || v <= -1e14 {
		return "", fmt.Errorf("soap fixed14.4: value %v out of bounds", v)
	}
	return strconv.FormatFloat(v, 'f', 4, 64), nil
}

func UnmarshalFixed14_4(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	if v >= 1e14 || v <= -1e14 {
		return 0, fmt.Errorf("soap fixed14.4: value %q out of bounds", s)
	}
	return v, nil
}

func MarshalChar(v rune) (string, error) {
	if v == 0 {
		return "", errors.New("soap char: rune 0 is not allowed")
	}
	return string(v), nil
}

func UnmarshalChar(s string) (rune, error) {
	if len(s) == 0 {
		return 0, errors.New("soap char: got empty string")
	}
	r, n := utf8.DecodeRune([]byte(s))
	if n != len(s) {
		return 0, fmt.Errorf("soap char: value %q is not a single rune", s)
	}
	return r, nil
}

func MarshalDate(v time.Time) (string, error) {
	return v.Format("2006-01-02"), nil
}

var dateFmts = []string{"2006-01-02", "20060102"}

func UnmarshalDate(s string) (time.Time, error) {
	for _, f := range dateFmts {
		if t, err := time.ParseInLocation(f, s, time.Local); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("soap date: value %q is not in a recognized date format", s)
}

// TimeOfDay is used in cases where SOAP "time" or "time.tz" is used.
type TimeOfDay struct {
	// Duration of time since midnight.
	FromMidnight time.Duration

	// TimeZone is present only if time.tz is used. It is otherwise ignored.
	TimeZone *time.Location
}

func MarshalTimeOfDay(v TimeOfDay) (string, error) {
	d := int64(v.FromMidnight / time.Second)
	hour := d / 3600
	d = d % 3600
	minute := d / 60
	second := d % 60

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second), nil
}

var timeRegexp = regexp.MustCompile(`^(\d\d)(?::?(\d\d)(?::?(\d\d))?)?$`)

func UnmarshalTimeOfDay(s string) (TimeOfDay, error) {
	parts := timeRegexp.FindStringSubmatch(s)
	if len(parts) < 2 {
		return TimeOfDay{}, fmt.Errorf("soap time: value %q is not in ISO8601 time format", s)
	}
	parts = parts[1:]
	var iParts [3]int64
	for i, pStr := range parts {
		iParts[i], _ = strconv.ParseInt(pStr, 10, 64)
	}
	return TimeOfDay{time.Duration(iParts[0]*3600+iParts[1]*60+iParts[2]) * time.Second, nil}, nil
}
