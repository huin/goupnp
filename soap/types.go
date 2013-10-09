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

// MarshalDate marshals time.Time to SOAP "date" type. Note that this converts
// to local time, and discards the time-of-day components.
func MarshalDate(v time.Time) (string, error) {
	return v.Local().Format("2006-01-02"), nil
}

var dateFmts = []string{"2006-01-02", "20060102"}

// UnmarshalDate unmarshals time.Time from SOAP "date" type. This outputs the
// date as midnight in the local time zone.
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

	// Set to true if Offset is specified. If false, then the timezone is
	// unspecified (and by ISO8601 - implies some "local" time).
	HasOffset bool

	// Offset is non-zero only if time.tz is used. It is otherwise ignored. If
	// non-zero, then it is regarded as a UTC offset in seconds. Note that the
	// sub-minutes is ignored by the marshal function.
	Offset int16
}

// MarshalTimeOfDay marshals TimeOfDay to the "time" type.
func MarshalTimeOfDay(v TimeOfDay) (string, error) {
	d := int64(v.FromMidnight / time.Second)
	hour := d / 3600
	d = d % 3600
	minute := d / 60
	second := d % 60

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second), nil
}

// UnmarshalTimeOfDay unmarshals TimeOfDay from the "time" type.
func UnmarshalTimeOfDay(s string) (TimeOfDay, error) {
	t, err := UnmarshalTimeOfDayTz(s)
	if err != nil {
		return TimeOfDay{}, err
	} else if t.HasOffset {
		return TimeOfDay{}, fmt.Errorf("soap time: value %q contains unexpected timezone")
	}
	return t, nil
}

// MarshalTimeOfDayTz marshals TimeOfDay to the "time.tz" type.
func MarshalTimeOfDayTz(v TimeOfDay) (string, error) {
	d := int64(v.FromMidnight / time.Second)
	hour := d / 3600
	d = d % 3600
	minute := d / 60
	second := d % 60

	tz := ""
	if v.HasOffset {
		if v.Offset == 0 {
			tz = "Z"
		} else {
			offsetMins := v.Offset / 60
			sign := '+'
			if offsetMins < 1 {
				offsetMins = -offsetMins
				sign = '-'
			}
			tz = fmt.Sprintf("%c%02d:%02d", sign, offsetMins/60, offsetMins%60)
		}
	}

	return fmt.Sprintf("%02d:%02d:%02d%s", hour, minute, second, tz), nil
}

var timeRegexp = regexp.MustCompile(
	`^(\d\d)(?::?(\d\d)(?::?(\d\d))?)?` + // hh[:mm[:ss]]
		`(?:(Z)|([+-])(\d\d)(?::?(\d\d))?)?$`) // Z | ±hh[:mm]

// UnmarshalTimeOfDayTz unmarshals TimeOfDay from the "time.tz" type.
func UnmarshalTimeOfDayTz(s string) (TimeOfDay, error) {
	parts := timeRegexp.FindStringSubmatch(s)
	if parts == nil {
		return TimeOfDay{}, fmt.Errorf("soap time.tz: value %q is not in ISO8601 time format", s)
	}

	// HH:MM:SS parsing.
	parts = parts[1:]
	var iParts [3]int64
	for i, pStr := range parts[:3] {
		iParts[i], _ = strconv.ParseInt(pStr, 10, 64)
	}
	hour, minute, second := iParts[0], iParts[1], iParts[2]

	fromMidnight := time.Duration(hour*3600+minute*60+second) * time.Second

	// ISO8601 special case - values up to 24:00:00 are allowed, so using
	// strictly greater-than for the maximum value.
	if fromMidnight > 24*time.Hour || minute >= 60 || second >= 60 {
		return TimeOfDay{}, fmt.Errorf("soap time.tz: value %q has value(s) out of range", s)
	}

	// Timezone offset parsing.
	hasOffset := false
	var offset int64
	if parts[3] == "Z" {
		hasOffset = true
		offset = 0
	} else if parts[4] != "" {
		hasOffset = true
		hours, _ := strconv.ParseInt(parts[5], 10, 64)
		var mins int64
		if parts[6] != "" {
			mins, _ = strconv.ParseInt(parts[6], 10, 64)
		}
		offset = hours*3600 + mins*60
		if parts[4] == "-" {
			offset = -offset
		}
	}
	return TimeOfDay{
		FromMidnight: time.Duration(hour*3600+minute*60+second) * time.Second,
		HasOffset:    hasOffset,
		Offset:       int16(offset),
	}, nil
}
