package soap

import (
	"errors"
	"fmt"
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
		if t, err := time.Parse(f, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("soap date: value %q is not in a recognized date format", s)
}
