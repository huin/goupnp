// Package types defines types that encode values in SOAP requests and responses.
//
// Based on https://openconnectivity.org/upnp-specs/UPnP-arch-DeviceArchitecture-v2.0-20200417.pdf
// pages 58-60.
//
// Date/time formats are based on http://www.w3.org/TR/1998/NOTE-datetime-19980827
package types

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type SOAPValue interface {
	Marshal() (string, error)
	Unmarshal(s string) error
}

type UI1 uint8

var _ SOAPValue = new(UI1)

func NewUI1(v uint8) *UI1 {
	v2 := UI1(v)
	return &v2
}

func (v *UI1) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UI1) Marshal() (string, error) {
	return v.String(), nil
}

func (v *UI1) Unmarshal(s string) error {
	v2, err := strconv.ParseUint(s, 10, 8)
	*v = UI1(v2)
	return err
}

type UI2 uint16

var _ SOAPValue = new(UI2)

func NewUI2(v uint16) *UI2 {
	v2 := UI2(v)
	return &v2
}

func (v *UI2) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UI2) Marshal() (string, error) {
	return v.String(), nil
}

func (v *UI2) Unmarshal(s string) error {
	v2, err := strconv.ParseUint(s, 10, 16)
	*v = UI2(v2)
	return err
}

type UI4 uint32

var _ SOAPValue = new(UI4)

func NewUI4(v uint32) *UI4 {
	v2 := UI4(v)
	return &v2
}

func (v *UI4) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UI4) Marshal() (string, error) {
	return v.String(), nil
}

func (v *UI4) Unmarshal(s string) error {
	v2, err := strconv.ParseUint(s, 10, 32)
	*v = UI4(v2)
	return err
}

type UI8 uint64

var _ SOAPValue = new(UI8)

func NewUI8(v uint64) *UI8 {
	v2 := UI8(v)
	return &v2
}

func (v *UI8) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UI8) Marshal() (string, error) {
	return v.String(), nil
}

func (v *UI8) Unmarshal(s string) error {
	v2, err := strconv.ParseUint(s, 10, 64)
	*v = UI8(v2)
	return err
}

type I1 int8

var _ SOAPValue = new(I1)

func NewI1(v int8) *I1 {
	v2 := I1(v)
	return &v2
}

func (v *I1) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *I1) Marshal() (string, error) {
	return v.String(), nil
}

func (v *I1) Unmarshal(s string) error {
	v2, err := strconv.ParseInt(s, 10, 8)
	*v = I1(v2)
	return err
}

type I2 int16

var _ SOAPValue = new(I2)

func NewI2(v int16) *I2 {
	v2 := I2(v)
	return &v2
}

func (v *I2) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *I2) Marshal() (string, error) {
	return v.String(), nil
}

func (v *I2) Unmarshal(s string) error {
	v2, err := strconv.ParseInt(s, 10, 16)
	*v = I2(v2)
	return err
}

type I4 int32

var _ SOAPValue = new(I4)

func NewI4(v int32) *I4 {
	v2 := I4(v)
	return &v2
}

func (v *I4) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *I4) Marshal() (string, error) {
	return v.String(), nil
}

func (v *I4) Unmarshal(s string) error {
	v2, err := strconv.ParseInt(s, 10, 32)
	*v = I4(v2)
	return err
}

type I8 int64

var _ SOAPValue = new(I8)

func NewI8(v int64) *I8 {
	v2 := I8(v)
	return &v2
}

func (v *I8) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *I8) Marshal() (string, error) {
	return v.String(), nil
}

func (v *I8) Unmarshal(s string) error {
	v2, err := strconv.ParseInt(s, 10, 64)
	*v = I8(v2)
	return err
}

type R4 float32

var _ SOAPValue = new(R4)

func NewR4(v float32) *R4 {
	v2 := R4(v)
	return &v2
}

func (v *R4) String() string {
	return strconv.FormatFloat(float64(*v), 'G', -1, 32)
}

func (v *R4) Marshal() (string, error) {
	return v.String(), nil
}

func (v *R4) Unmarshal(s string) error {
	v2, err := strconv.ParseFloat(s, 32)
	*v = R4(v2)
	return err
}

type R8 float64

var _ SOAPValue = new(R8)

func NewR8(v float64) *R8 {
	v2 := R8(v)
	return &v2
}

func (v *R8) String() string {
	return strconv.FormatFloat(float64(*v), 'G', -1, 64)
}

func (v *R8) Marshal() (string, error) {
	return v.String(), nil
}

func (v *R8) Unmarshal(s string) error {
	v2, err := strconv.ParseFloat(s, 64)
	*v = R8(v2)
	return err
}

const Fixed14_4Denominator = 1e4
const Fixed14_4MaxInteger = 1e14 - 1
const Fixed14_4MaxFractional = 1e18 - 1

// Fixed14_4 represents a fixed point number with up to 14 decimal digits
// before the decimal point (integer part), and up to 4 decimal digits
// after the decimal point (fractional part).
//
// Corresponds to the SOAP "fixed.14.4" type.
//
// This is a struct to avoid accidentally using the value directly as an
// integer.
type Fixed14_4 struct {
	// Fractional divided by 1e4 is the fixed point value. Take care setting
	// this directly, it should only contain values in the range (-1e18, 1e18).
	Fractional int64
}

var _ SOAPValue = &Fixed14_4{}

// Fixed14_4FromParts creates a Fixed14_4 from components.
// Bounds:
//   * Both intPart and fracPart must have the same sign.
//   * -1e14 < intPart < 1e14
//   * -1e4 < fracPart < 1e4
func Fixed14_4FromParts(intPart int64, fracPart int16) (Fixed14_4, error) {
	var v Fixed14_4
	err := v.SetParts(intPart, fracPart)
	return v, err
}

// SetFromParts sets the value based on the integer component and the fractional component.
// Bounds:
//   * Both intPart and fracPart must have the same sign.
//   * -1e14 < intPart < 1e14
//   * -1e4 < fracPart < 1e4
func (v *Fixed14_4) SetParts(intPart int64, fracPart int16) error {
	if (intPart < 0) != (fracPart < 0) {
		return fmt.Errorf("want intPart and fracPart with same sign, got %d and %d",
			intPart, fracPart)
	}
	if intPart < -Fixed14_4MaxInteger || intPart > Fixed14_4MaxInteger {
		return fmt.Errorf("want intPart in range (-1e14,1e14), got %d", intPart)
	}
	if fracPart < -Fixed14_4Denominator || fracPart > Fixed14_4Denominator {
		return fmt.Errorf("want fracPart in range (-1e4,1e4), got %d", fracPart)
	}
	v.Fractional = intPart*Fixed14_4Denominator + int64(fracPart)
	return nil
}

// Returns the integer part and fractional part of the fixed point number.
func (v Fixed14_4) Parts() (int64, int16) {
	return v.Fractional / Fixed14_4Denominator, int16(v.Fractional % Fixed14_4Denominator)
}

// Fixed14_4FromFractional creates a Fixed14_4 from an integer, where the
// parameter divided by 1e4 is the fixed point value.
func Fixed14_4FromFractional(fracValue int64) (Fixed14_4, error) {
	var v Fixed14_4
	err := v.SetFractional(fracValue)
	return v, err
}

// SetFromFractional sets the value of the fixed point number, where fracValue
// divided by 1e4 is the fixed point value. Unlike setting v.Fractional
// directly, this checks the value.
func (v *Fixed14_4) SetFractional(fracValue int64) error {
	if fracValue < -Fixed14_4MaxFractional || fracValue > Fixed14_4MaxFractional {
		return fmt.Errorf("want intPart in range (-1e18,1e18), got %d", fracValue)
	}
	v.Fractional = fracValue
	return nil
}

// Fixed14_4FromFloat creates a Fixed14_4 from a float64. Returns error if the
// float is outside the range.
func Fixed14_4FromFloat(f float64) (Fixed14_4, error) {
	i := int64(f * Fixed14_4Denominator)
	return Fixed14_4FromFractional(i)
}

// SetFloat64 sets the value of the fixed point number from a float64. Returns
// error if the float is outside the range.
func (v *Fixed14_4) SetFloat64(f float64) error {
	i := int64(f * Fixed14_4Denominator)
	return v.SetFractional(i)
}

func (v Fixed14_4) Float64() float64 {
	return float64(v.Fractional) / Fixed14_4Denominator
}

func (v Fixed14_4) String() string {
	intPart, fracPart := v.Parts()
	if fracPart < 0 {
		fracPart = -fracPart
	}
	return fmt.Sprintf("%d.%04d", intPart, fracPart)
}

func (v *Fixed14_4) Marshal() (string, error) {
	return v.String(), nil
}

func (v *Fixed14_4) Unmarshal(s string) error {
	parts := strings.SplitN(s, ".", 2)
	intPart, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return err
	}

	var fracPart int64
	if len(parts) >= 2 && len(parts[1]) > 0 {
		fracStr := parts[1]

		for _, r := range fracStr {
			if r < '0' || r > '9' {
				return fmt.Errorf("found non-digit in fractional component of %q", s)
			}
		}

		// Take only the 4 most significant digits of the fractional component.
		fracStr = fracStr[:min(len(fracStr), 4)]

		fracPart, err = strconv.ParseInt(fracStr, 10, 16)
		if err != nil {
			return err
		}
		if fracPart < 0 {
			// This shouldn't happen by virtue of earlier digit-only check.
			return fmt.Errorf("got negative fractional component in %q", s)
		}

		switch len(fracStr) {
		case 1:
			fracPart *= 1000
		case 2:
			fracPart *= 100
		case 3:
			fracPart *= 10
		case 4:
			fracPart *= 1
		}
		if intPart < 0 {
			fracPart = -fracPart
		}
	}
	v.SetParts(intPart, int16(fracPart))
	return nil
}

// Char maps rune to SOAP "char" type.
type Char rune

var _ SOAPValue = new(Char)

func NewChar(v rune) *Char {
	v2 := Char(v)
	return &v2
}

func (v *Char) String() string {
	return string(*v)
}

func (v *Char) Marshal() (string, error) {
	if *v == 0 {
		return "", errors.New("soap char: rune 0 is not allowed")
	}
	return v.String(), nil
}

func (v *Char) Unmarshal(s string) error {
	if len(s) == 0 {
		return errors.New("soap char: got empty string")
	}
	v2, n := utf8.DecodeRune([]byte(s))
	if n != len(s) {
		return fmt.Errorf("soap char: value %q is not a single rune", s)
	}
	*v = Char(v2)
	return nil
}

type String string

var _ SOAPValue = new(String)

func NewString(v string) *String {
	v2 := String(v)
	return &v2
}

func (v *String) Marshal() (string, error) {
	return string(*v), nil
}

func (v *String) Unmarshal(s string) error {
	*v = String(s)
	return nil
}

func parseInt(s string, err *error) int {
	v, parseErr := strconv.ParseInt(s, 10, 64)
	if parseErr != nil {
		*err = parseErr
	}
	return int(v)
}

var dateRegexps = []*regexp.Regexp{
	// yyyy[-mm[-dd]]
	regexp.MustCompile(`^(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?$`),
	// yyyy[mm[dd]]
	regexp.MustCompile(`^(\d{4})(?:(\d{2})(?:(\d{2}))?)?$`),
}

type prefixRemainder struct {
	prefix    string
	remainder string
}

// prefixUntilAny returns a prefix of the leading string prior to any
// characters in `chars`, and the remainder. If no character from `chars` is
// present in `s`, then returns `s` as `prefix`, and empty remainder.
//
// prefixUntilAny("123/abc", "/") => {"123", "/abc"}
// prefixUntilAny("123", "/") => {"123", ""}
func prefixUntilAny(s string, chars string) prefixRemainder {
	i := strings.IndexAny(s, chars)
	if i == -1 {
		return prefixRemainder{prefix: s, remainder: ""}
	}
	return prefixRemainder{prefix: s[:i], remainder: s[i:]}
}

// TimeOfDay is used in cases where SOAP "time" or "time.tz" is used.
// It contains non-timezone aware components.
type TimeOfDay struct {
	Hour   int8
	Minute int8
	Second int8
}

var _ SOAPValue = &TimeOfDay{}

func TimeOfDayFromTime(t time.Time) TimeOfDay {
	h, m, s := t.Clock()
	return TimeOfDay{int8(h), int8(m), int8(s)}
}

func (tod *TimeOfDay) SetFromTime(t time.Time) {
	h, m, s := t.Clock()
	tod.Hour = int8(h)
	tod.Minute = int8(m)
	tod.Second = int8(s)
}

// Sets components based on duration since midnight.
func (tod *TimeOfDay) SetFromDuration(d time.Duration) error {
	if d < 0 || d > 24*time.Hour {
		return fmt.Errorf("out of range of SOAP time type: %v", d)
	}
	tod.Hour = int8(d / time.Hour)
	d = d % time.Hour
	tod.Minute = int8(d / time.Minute)
	d = d % time.Minute
	tod.Second = int8(d / time.Second)
	return nil
}

// Returns duration since midnight.
func (tod TimeOfDay) ToDuration() time.Duration {
	return time.Duration(tod.Hour)*time.Hour +
		time.Duration(tod.Minute)*time.Minute +
		time.Duration(tod.Second)*time.Second
}

func (tod TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", tod.Hour, tod.Minute, tod.Second)
}

// IsValid returns true iff v is positive and <= 24 hours.
// It allows equal to 24 hours as a special case as 24:00:00 is an allowed
// value by the SOAP type.
func (tod TimeOfDay) CheckValid() error {
	if (tod.Hour < 0 || tod.Minute < 0 || tod.Second < 0) ||
		(tod.Hour == 24 && (tod.Minute > 0 || tod.Second > 0)) ||
		tod.Hour > 24 || tod.Minute >= 60 || tod.Second >= 60 {
		return fmt.Errorf("soap time: value %v has components(s) out of range", tod)
	}
	return nil
}

// clear removes data from v, setting to default values.
func (tod *TimeOfDay) clear() {
	tod.Hour = 0
	tod.Minute = 0
	tod.Second = 0
}

func (tod *TimeOfDay) Marshal() (string, error) {
	if err := tod.CheckValid(); err != nil {
		return "", err
	}
	return tod.String(), nil
}

var timeRegexps = []*regexp.Regexp{
	// hh:mm:ss
	regexp.MustCompile(`^(\d{2})(\d{2})(\d{2})$`),
	// hhmmss
	regexp.MustCompile(`^(\d{2}):(\d{2}):(\d{2})$`),
}

func (tod *TimeOfDay) Unmarshal(s string) error {
	tod.clear()
	var parts []string
	for _, re := range timeRegexps {
		parts = re.FindStringSubmatch(s)
		if parts != nil {
			break
		}
	}
	if parts == nil {
		return fmt.Errorf("value %q is not in ISO8601 time format", s)
	}

	var err error
	tod.Hour = int8(parseInt(parts[1], &err))
	tod.Minute = int8(parseInt(parts[2], &err))
	tod.Second = int8(parseInt(parts[3], &err))
	if err != nil {
		return fmt.Errorf("value %q is not in ISO8601 time format: %v", s, err)
	}

	return tod.CheckValid()
}

// TimeOfDayTZ maps to the SOAP "time.tz" type.
type TimeOfDayTZ struct {
	// Components of the time of day.
	TimeOfDay TimeOfDay

	// Timezone designator.
	TZ TZD
}

var _ SOAPValue = &TimeOfDayTZ{}

func (todz TimeOfDayTZ) String() string {
	return fmt.Sprintf("%v%v", todz.TimeOfDay, todz.TZ)
}

// clear removes data from v, setting to default values.
func (todz *TimeOfDayTZ) clear() {
	todz.TimeOfDay.clear()
	todz.TZ.clear()
}

func (todz *TimeOfDayTZ) Marshal() (string, error) {
	return todz.String(), nil
}

func (todz *TimeOfDayTZ) Unmarshal(s string) error {
	todz.clear()
	parts := prefixUntilAny(s, "Z+-")
	if err := todz.TimeOfDay.Unmarshal(parts.prefix); err != nil {
		return err
	}
	return todz.TZ.unmarshal(parts.remainder)
}

// Date maps to the SOAP "date" type. Marshaling and Unmarshalling does *not*
// check if components are in range.
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

var _ SOAPValue = &Date{}

func DateFromTime(t time.Time) Date {
	var d Date
	d.SetFromTime(t)
	return d
}

func (d *Date) SetFromTime(t time.Time) {
	d.Year = t.Year()
	d.Month = t.Month()
	d.Day = t.Day()
}

// ToTime returns a time.Time from the date components, at midnight, and using
// the given location.
func (d Date) ToTime(loc *time.Location) time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

// CheckValid returns an error if the date components are out of range.
func (d Date) CheckValid() error {
	y, m, day := d.ToTime(time.UTC).Date()
	if y != d.Year || m != d.Month || day != d.Day {
		return fmt.Errorf("SOAP date component(s) out of range in %v", d)
	}
	return nil
}

// clear removes data from d, setting to default (invalid/zero) values.
func (d *Date) clear() {
	d.Year = 0
	d.Month = 0
	d.Day = 0
}

func (d *Date) Marshal() (string, error) {
	return d.String(), nil
}

func (d *Date) Unmarshal(s string) error {
	d.clear()
	var parts []string
	for _, re := range dateRegexps {
		parts = re.FindStringSubmatch(s)
		if parts != nil {
			break
		}
	}
	if parts == nil {
		return fmt.Errorf("error parsing date: value %q is not in a recognized ISO8601 date format", s)
	}

	var err error
	d.Year = parseInt(parts[1], &err)
	d.Month = time.January
	d.Day = 1
	if len(parts[2]) != 0 {
		d.Month = time.Month(parseInt(parts[2], &err))
		if len(parts[3]) != 0 {
			d.Day = parseInt(parts[3], &err)
		}
	}

	if err != nil {
		return fmt.Errorf("error parsing date %q: %v", s, err)
	}

	return nil
}

// DateTime maps to SOAP "dateTime" type.
type DateTime struct {
	Date      Date
	TimeOfDay TimeOfDay
}

var _ SOAPValue = &DateTime{}

func DateTimeFromTime(v time.Time) DateTime {
	dt := DateTime{}
	dt.Date.SetFromTime(v)
	dt.TimeOfDay.SetFromTime(v)
	return dt
}

func (dt DateTime) String() string {
	return fmt.Sprintf("%vT%v", dt.Date, dt.TimeOfDay)
}

func (dt DateTime) ToTime(loc *time.Location) time.Time {
	return time.Date(dt.Date.Year, dt.Date.Month, dt.Date.Day,
		int(dt.TimeOfDay.Hour), int(dt.TimeOfDay.Minute), int(dt.TimeOfDay.Second), 0,
		loc)
}

// clear removes data from dt, setting to default values.
func (dt *DateTime) clear() {
	dt.Date.clear()
	dt.TimeOfDay.clear()
}

func (dt *DateTime) Marshal() (string, error) {
	return dt.String(), nil
}

func (dt *DateTime) Unmarshal(s string) error {
	dt.clear()
	parts := prefixUntilAny(s, "T")
	if err := dt.Date.Unmarshal(parts.prefix); err != nil {
		return err
	}

	if parts.remainder == "" {
		return nil
	}

	if parts.remainder[0] != 'T' {
		return fmt.Errorf("missing 'T' time separator in dateTime %q", s)
	}

	return dt.TimeOfDay.Unmarshal(parts.remainder[1:])
}

// DateTime maps to SOAP type "dateTime.tz".
type DateTimeTZ struct {
	Date      Date
	TimeOfDay TimeOfDay
	TZ        TZD
}

var _ SOAPValue = &DateTimeTZ{}

func DateTimeTZFromTime(t time.Time) DateTimeTZ {
	return DateTimeTZ{
		Date:      DateFromTime(t),
		TimeOfDay: TimeOfDayFromTime(t),
		TZ:        TZDFromTime(t),
	}
}

func (dtz DateTimeTZ) String() string {
	return dtz.Date.String() + "T" + dtz.TimeOfDay.String() + dtz.TZ.String()
}

// Time converts `dtz` to time.Time, using defaultLoc as the default location if
// `dtz` contains no offset information.
func (dtz DateTimeTZ) Time(defaultLoc *time.Location) time.Time {
	return time.Date(
		dtz.Date.Year,
		dtz.Date.Month,
		dtz.Date.Day,
		int(dtz.TimeOfDay.Hour),
		int(dtz.TimeOfDay.Minute),
		int(dtz.TimeOfDay.Second),
		0,
		dtz.TZ.Location(defaultLoc),
	)
}

// clear removes data from dtz, setting to default values.
func (dtz *DateTimeTZ) clear() {
	dtz.Date.clear()
	dtz.TimeOfDay.clear()
	dtz.TZ.clear()
}

func (dtz *DateTimeTZ) Marshal() (string, error) {
	return dtz.String(), nil
}

func (dtz *DateTimeTZ) Unmarshal(s string) error {
	dtz.clear()
	dateParts := prefixUntilAny(s, "T")
	if err := dtz.Date.Unmarshal(dateParts.prefix); err != nil {
		return err
	}

	if dateParts.remainder == "" {
		return nil
	}

	// Trim the leading "T" between date and time.
	remainder := dateParts.remainder[1:]
	timeParts := prefixUntilAny(remainder, "Z+-")
	if err := dtz.TimeOfDay.Unmarshal(timeParts.prefix); err != nil {
		return err
	}

	if timeParts.remainder == "" {
		return nil
	}

	return dtz.TZ.unmarshal(timeParts.remainder)
}

// TZD is a timezone designator. Not a full SOAP time in itself, but used as
// part of timezone-aware date/time types that are.
type TZD struct {
	// Offset is the timezone offset in seconds. Note that the SOAP encoding
	// only encodes precisions up to minutes, this is in seconds for
	// interoperability with time.Time.
	Offset int
	// HasTZ specifies if a timezone offset is specified or not.
	HasTZ bool
}

func TZDFromTime(t time.Time) TZD {
	_, offset := t.Zone()
	return TZD{
		Offset: offset,
		HasTZ:  true,
	}
}

func TZDOffset(secondsOffset int) TZD {
	return TZD{
		Offset: secondsOffset,
		HasTZ:  true,
	}
}

// Location returns an appropriate *time.Location (time.UTC or time.FixedZone),
// or defaultLoc if `v` contains no offset information.
func (tzd TZD) Location(defaultLoc *time.Location) *time.Location {
	if !tzd.HasTZ {
		return defaultLoc
	}
	if tzd.Offset == 0 {
		return time.UTC
	}
	return time.FixedZone(tzd.String(), tzd.Offset)
}

func (tzd TZD) String() string {
	if !tzd.HasTZ {
		return ""
	}

	if tzd.Offset == 0 {
		return "Z"
	}

	offsetMins := tzd.Offset / 60
	sign := '+'
	if offsetMins < 1 {
		offsetMins = -offsetMins
		sign = '-'
	}
	h, m := offsetMins/60, offsetMins%60
	return fmt.Sprintf("%c%02d:%02d", sign, h, m)
}

// clear removes offset information from `v`.
func (tzd *TZD) clear() {
	tzd.Offset = 0
	tzd.HasTZ = false
}

// (+|-)(hh):(mm)
var timezoneRegexp = regexp.MustCompile(`^([+-])(\d{2}):(\d{2})$`)

func (tzd *TZD) unmarshal(s string) error {
	tzd.clear()
	if s == "" {
		return nil
	}
	tzd.HasTZ = true
	if s == "Z" {
		return nil
	}
	parts := timezoneRegexp.FindStringSubmatch(s)
	if parts == nil {
		return fmt.Errorf("value %q is not in ISO8601 timezone format", s)
	}

	var err error
	tzd.Offset = parseInt(parts[2], &err) * 3600
	tzd.Offset += parseInt(parts[3], &err) * 60
	if parts[1] == "-" {
		tzd.Offset = -tzd.Offset
	}

	if err != nil {
		err = fmt.Errorf("value %q is not in ISO8601 timezone format: %v", s, err)
	}

	return nil
}

type Boolean bool

var _ SOAPValue = new(Boolean)

func NewBoolean(v bool) *Boolean {
	v2 := Boolean(v)
	return &v2
}

func (v *Boolean) String() string {
	if *v {
		return "true"
	}
	return "false"
}

func (v *Boolean) Marshal() (string, error) {
	if *v {
		return "1", nil
	}
	return "0", nil
}

func (v *Boolean) Unmarshal(s string) error {
	switch s {
	case "0", "false", "no":
		*v = false
	case "1", "true", "yes":
		*v = true
	default:
		return fmt.Errorf("soap boolean: %q is not a valid boolean value", s)
	}
	return nil
}

// BinBase64 maps []byte to SOAP "bin.base64" type.
type BinBase64 []byte

var _ SOAPValue = new(BinBase64)

func NewBinBase64(v []byte) *BinBase64 {
	v2 := BinBase64(v)
	return &v2
}

func (v *BinBase64) String() string {
	return base64.StdEncoding.EncodeToString(*v)
}

func (v *BinBase64) Marshal() (string, error) {
	return v.String(), nil
}

func (v *BinBase64) Unmarshal(s string) error {
	v2, err := base64.StdEncoding.DecodeString(s)
	*v = v2
	return err
}

// BinHex maps []byte to SOAP "bin.hex" type.
type BinHex []byte

var _ SOAPValue = new(BinHex)

func NewBinHex(v []byte) *BinHex {
	v2 := BinHex(v)
	return &v2
}

func (v *BinHex) String() string {
	return hex.EncodeToString(*v)
}

func (v *BinHex) Marshal() (string, error) {
	return v.String(), nil
}

func (v *BinHex) Unmarshal(s string) error {
	v2, err := hex.DecodeString(s)
	*v = v2
	return err
}

// URI maps *url.URL to SOAP "uri" type.
type URI url.URL

var _ SOAPValue = new(URI)

func (v *URI) String() string {
	return v.ToURL().String()
}

func (v *URI) ToURL() *url.URL {
	return (*url.URL)(v)
}

func (v *URI) Marshal() (string, error) {
	return (*url.URL)(v).String(), nil
}

func (v *URI) Unmarshal(s string) error {
	v2, err := url.Parse(s)
	if err != nil {
		return err
	}
	*v = URI(*v2)
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
