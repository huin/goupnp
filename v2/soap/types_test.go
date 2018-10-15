package soap

import (
	"bytes"
	"math"
	"net/url"
	"testing"
	"time"
)

type convTest interface {
	marshal() (string, error)
	unmarshal(string) (interface{}, error)
	equal(result interface{}) bool
}

// duper is an interface that convTest values may optionally also implement to
// generate another convTest for a value in an otherwise identical testCase.
type duper interface {
	dupe(tag string) []convTest
}

type testCase struct {
	value            convTest
	str              string
	wantMarshalErr   bool
	wantUnmarshalErr bool
	noMarshal        bool
	noUnMarshal      bool
	tag              string
}

type ui1Test uint8

func (v ui1Test) marshal() (string, error) {
	return MarshalUI1(uint8(v))
}
func (v ui1Test) unmarshal(s string) (interface{}, error) {
	return UnmarshalUI1(s)
}
func (v ui1Test) equal(result interface{}) bool {
	return uint8(v) == result.(uint8)
}
func (v ui1Test) dupe(tag string) []convTest {
	if tag == "dupe" {
		return []convTest{
			ui2Test(v),
			ui4Test(v),
		}
	}
	return nil
}

type ui2Test uint16

func (v ui2Test) marshal() (string, error) {
	return MarshalUI2(uint16(v))
}
func (v ui2Test) unmarshal(s string) (interface{}, error) {
	return UnmarshalUI2(s)
}
func (v ui2Test) equal(result interface{}) bool {
	return uint16(v) == result.(uint16)
}

type ui4Test uint32

func (v ui4Test) marshal() (string, error) {
	return MarshalUI4(uint32(v))
}
func (v ui4Test) unmarshal(s string) (interface{}, error) {
	return UnmarshalUI4(s)
}
func (v ui4Test) equal(result interface{}) bool {
	return uint32(v) == result.(uint32)
}

type i1Test int8

func (v i1Test) marshal() (string, error) {
	return MarshalI1(int8(v))
}
func (v i1Test) unmarshal(s string) (interface{}, error) {
	return UnmarshalI1(s)
}
func (v i1Test) equal(result interface{}) bool {
	return int8(v) == result.(int8)
}
func (v i1Test) dupe(tag string) []convTest {
	if tag == "dupe" {
		return []convTest{
			i2Test(v),
			i4Test(v),
		}
	}
	return nil
}

type i2Test int16

func (v i2Test) marshal() (string, error) {
	return MarshalI2(int16(v))
}
func (v i2Test) unmarshal(s string) (interface{}, error) {
	return UnmarshalI2(s)
}
func (v i2Test) equal(result interface{}) bool {
	return int16(v) == result.(int16)
}

type i4Test int32

func (v i4Test) marshal() (string, error) {
	return MarshalI4(int32(v))
}
func (v i4Test) unmarshal(s string) (interface{}, error) {
	return UnmarshalI4(s)
}
func (v i4Test) equal(result interface{}) bool {
	return int32(v) == result.(int32)
}

type intTest int64

func (v intTest) marshal() (string, error) {
	return MarshalInt(int64(v))
}
func (v intTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalInt(s)
}
func (v intTest) equal(result interface{}) bool {
	return int64(v) == result.(int64)
}

type fixed14_4Test float64

func (v fixed14_4Test) marshal() (string, error) {
	return MarshalFixed14_4(float64(v))
}
func (v fixed14_4Test) unmarshal(s string) (interface{}, error) {
	return UnmarshalFixed14_4(s)
}
func (v fixed14_4Test) equal(result interface{}) bool {
	return math.Abs(float64(v)-result.(float64)) < 0.001
}

type charTest rune

func (v charTest) marshal() (string, error) {
	return MarshalChar(rune(v))
}
func (v charTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalChar(s)
}
func (v charTest) equal(result interface{}) bool {
	return rune(v) == result.(rune)
}

type dateTest struct{ time.Time }

func (v dateTest) marshal() (string, error) {
	return MarshalDate(time.Time(v.Time))
}
func (v dateTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalDate(s)
}
func (v dateTest) equal(result interface{}) bool {
	return v.Time.Equal(result.(time.Time))
}
func (v dateTest) dupe(tag string) []convTest {
	if tag != "no:dateTime" {
		return []convTest{dateTimeTest{v.Time}}
	}
	return nil
}

type timeOfDayTest struct {
	TimeOfDay
}

func (v timeOfDayTest) marshal() (string, error) {
	return MarshalTimeOfDay(v.TimeOfDay)
}
func (v timeOfDayTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalTimeOfDay(s)
}
func (v timeOfDayTest) equal(result interface{}) bool {
	return v.TimeOfDay == result.(TimeOfDay)
}
func (v timeOfDayTest) dupe(tag string) []convTest {
	if tag != "no:time.tz" {
		return []convTest{timeOfDayTzTest{v.TimeOfDay}}
	}
	return nil
}

type timeOfDayTzTest struct {
	TimeOfDay
}

func (v timeOfDayTzTest) marshal() (string, error) {
	return MarshalTimeOfDayTz(v.TimeOfDay)
}
func (v timeOfDayTzTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalTimeOfDayTz(s)
}
func (v timeOfDayTzTest) equal(result interface{}) bool {
	return v.TimeOfDay == result.(TimeOfDay)
}

type dateTimeTest struct{ time.Time }

func (v dateTimeTest) marshal() (string, error) {
	return MarshalDateTime(time.Time(v.Time))
}
func (v dateTimeTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalDateTime(s)
}
func (v dateTimeTest) equal(result interface{}) bool {
	return v.Time.Equal(result.(time.Time))
}
func (v dateTimeTest) dupe(tag string) []convTest {
	if tag != "no:dateTime.tz" {
		return []convTest{dateTimeTzTest{v.Time}}
	}
	return nil
}

type dateTimeTzTest struct{ time.Time }

func (v dateTimeTzTest) marshal() (string, error) {
	return MarshalDateTimeTz(time.Time(v.Time))
}
func (v dateTimeTzTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalDateTimeTz(s)
}
func (v dateTimeTzTest) equal(result interface{}) bool {
	return v.Time.Equal(result.(time.Time))
}

type booleanTest bool

func (v booleanTest) marshal() (string, error) {
	return MarshalBoolean(bool(v))
}
func (v booleanTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalBoolean(s)
}
func (v booleanTest) equal(result interface{}) bool {
	return bool(v) == result.(bool)
}

type binBase64Test []byte

func (v binBase64Test) marshal() (string, error) {
	return MarshalBinBase64([]byte(v))
}
func (v binBase64Test) unmarshal(s string) (interface{}, error) {
	return UnmarshalBinBase64(s)
}
func (v binBase64Test) equal(result interface{}) bool {
	return bytes.Equal([]byte(v), result.([]byte))
}

type binHexTest []byte

func (v binHexTest) marshal() (string, error) {
	return MarshalBinHex([]byte(v))
}
func (v binHexTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalBinHex(s)
}
func (v binHexTest) equal(result interface{}) bool {
	return bytes.Equal([]byte(v), result.([]byte))
}

type uriTest struct{ URL *url.URL }

func (v uriTest) marshal() (string, error) {
	return MarshalURI(v.URL)
}
func (v uriTest) unmarshal(s string) (interface{}, error) {
	return UnmarshalURI(s)
}
func (v uriTest) equal(result interface{}) bool {
	return v.URL.String() == result.(*url.URL).String()
}

func Test(t *testing.T) {
	const time010203 time.Duration = (1*3600 + 2*60 + 3) * time.Second
	const time0102 time.Duration = (1*3600 + 2*60) * time.Second
	const time01 time.Duration = (1 * 3600) * time.Second
	const time235959 time.Duration = (23*3600 + 59*60 + 59) * time.Second

	// Fake out the local time for the implementation.
	localLoc = time.FixedZone("Fake/Local", 6*3600)
	defer func() {
		localLoc = time.Local
	}()

	tests := []testCase{
		// ui1
		{str: "", value: ui1Test(0), wantUnmarshalErr: true, noMarshal: true, tag: "dupe"},
		{str: " ", value: ui1Test(0), wantUnmarshalErr: true, noMarshal: true, tag: "dupe"},
		{str: "abc", value: ui1Test(0), wantUnmarshalErr: true, noMarshal: true, tag: "dupe"},
		{str: "-1", value: ui1Test(0), wantUnmarshalErr: true, noMarshal: true, tag: "dupe"},
		{str: "0", value: ui1Test(0), tag: "dupe"},
		{str: "1", value: ui1Test(1), tag: "dupe"},
		{str: "255", value: ui1Test(255), tag: "dupe"},
		{str: "256", value: ui1Test(0), wantUnmarshalErr: true, noMarshal: true},

		// ui2
		{str: "65535", value: ui2Test(65535)},
		{str: "65536", value: ui2Test(0), wantUnmarshalErr: true, noMarshal: true},

		// ui4
		{str: "4294967295", value: ui4Test(4294967295)},
		{str: "4294967296", value: ui4Test(0), wantUnmarshalErr: true, noMarshal: true},

		// i1
		{str: "", value: i1Test(0), wantUnmarshalErr: true, noMarshal: true, tag: "dupe"},
		{str: " ", value: i1Test(0), wantUnmarshalErr: true, noMarshal: true, tag: "dupe"},
		{str: "abc", value: i1Test(0), wantUnmarshalErr: true, noMarshal: true, tag: "dupe"},
		{str: "0", value: i1Test(0), tag: "dupe"},
		{str: "-1", value: i1Test(-1), tag: "dupe"},
		{str: "127", value: i1Test(127), tag: "dupe"},
		{str: "-128", value: i1Test(-128), tag: "dupe"},
		{str: "128", value: i1Test(0), wantUnmarshalErr: true, noMarshal: true},
		{str: "-129", value: i1Test(0), wantUnmarshalErr: true, noMarshal: true},

		// i2
		{str: "32767", value: i2Test(32767)},
		{str: "-32768", value: i2Test(-32768)},
		{str: "32768", value: i2Test(0), wantUnmarshalErr: true, noMarshal: true},
		{str: "-32769", value: i2Test(0), wantUnmarshalErr: true, noMarshal: true},

		// i4
		{str: "2147483647", value: i4Test(2147483647)},
		{str: "-2147483648", value: i4Test(-2147483648)},
		{str: "2147483648", value: i4Test(0), wantUnmarshalErr: true, noMarshal: true},
		{str: "-2147483649", value: i4Test(0), wantUnmarshalErr: true, noMarshal: true},

		// int
		{str: "9223372036854775807", value: intTest(9223372036854775807)},
		{str: "-9223372036854775808", value: intTest(-9223372036854775808)},
		{str: "9223372036854775808", value: intTest(0), wantUnmarshalErr: true, noMarshal: true},
		{str: "-9223372036854775809", value: intTest(0), wantUnmarshalErr: true, noMarshal: true},

		// fixed.14.4
		{str: "0.0000", value: fixed14_4Test(0)},
		{str: "1.0000", value: fixed14_4Test(1)},
		{str: "1.2346", value: fixed14_4Test(1.23456)},
		{str: "-1.0000", value: fixed14_4Test(-1)},
		{str: "-1.2346", value: fixed14_4Test(-1.23456)},
		{str: "10000000000000.0000", value: fixed14_4Test(1e13)},
		{str: "100000000000000.0000", value: fixed14_4Test(1e14), wantMarshalErr: true, wantUnmarshalErr: true},
		{str: "-10000000000000.0000", value: fixed14_4Test(-1e13)},
		{str: "-100000000000000.0000", value: fixed14_4Test(-1e14), wantMarshalErr: true, wantUnmarshalErr: true},

		// char
		{str: "a", value: charTest('a')},
		{str: "z", value: charTest('z')},
		{str: "\u1234", value: charTest(0x1234)},
		{str: "aa", value: charTest(0), wantMarshalErr: true, wantUnmarshalErr: true},
		{str: "", value: charTest(0), wantMarshalErr: true, wantUnmarshalErr: true},

		// date
		{str: "2013-10-08", value: dateTest{time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)}, tag: "no:dateTime"},
		{str: "20131008", value: dateTest{time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)}, noMarshal: true, tag: "no:dateTime"},
		{str: "2013-10-08T10:30:50", value: dateTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:dateTime"},
		{str: "2013-10-08T10:30:50Z", value: dateTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:dateTime"},
		{str: "", value: dateTest{}, wantMarshalErr: true, wantUnmarshalErr: true, noMarshal: true},
		{str: "-1", value: dateTest{}, wantUnmarshalErr: true, noMarshal: true},

		// time
		{str: "00:00:00", value: timeOfDayTest{TimeOfDay{FromMidnight: 0}}},
		{str: "000000", value: timeOfDayTest{TimeOfDay{FromMidnight: 0}}, noMarshal: true},
		{str: "24:00:00", value: timeOfDayTest{TimeOfDay{FromMidnight: 24 * time.Hour}}, noMarshal: true}, // ISO8601 special case
		{str: "24:01:00", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "24:00:01", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "25:00:00", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "00:60:00", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "00:00:60", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "01:02:03", value: timeOfDayTest{TimeOfDay{FromMidnight: time010203}}},
		{str: "010203", value: timeOfDayTest{TimeOfDay{FromMidnight: time010203}}, noMarshal: true},
		{str: "23:59:59", value: timeOfDayTest{TimeOfDay{FromMidnight: time235959}}},
		{str: "235959", value: timeOfDayTest{TimeOfDay{FromMidnight: time235959}}, noMarshal: true},
		{str: "01:02", value: timeOfDayTest{TimeOfDay{FromMidnight: time0102}}, noMarshal: true},
		{str: "0102", value: timeOfDayTest{TimeOfDay{FromMidnight: time0102}}, noMarshal: true},
		{str: "01", value: timeOfDayTest{TimeOfDay{FromMidnight: time01}}, noMarshal: true},
		{str: "foo 01:02:03", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "foo\n01:02:03", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "01:02:03 foo", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "01:02:03\nfoo", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "01:02:03Z", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "01:02:03+01", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "01:02:03+01:23", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "01:02:03+0123", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "01:02:03-01", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "01:02:03-01:23", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},
		{str: "01:02:03-0123", value: timeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:time.tz"},

		// time.tz
		{str: "24:00:01", value: timeOfDayTzTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "01Z", value: timeOfDayTzTest{TimeOfDay{time01, true, 0}}, noMarshal: true},
		{str: "01:02:03Z", value: timeOfDayTzTest{TimeOfDay{time010203, true, 0}}},
		{str: "01+01", value: timeOfDayTzTest{TimeOfDay{time01, true, 3600}}, noMarshal: true},
		{str: "01:02:03+01", value: timeOfDayTzTest{TimeOfDay{time010203, true, 3600}}, noMarshal: true},
		{str: "01:02:03+01:23", value: timeOfDayTzTest{TimeOfDay{time010203, true, 3600 + 23*60}}},
		{str: "01:02:03+0123", value: timeOfDayTzTest{TimeOfDay{time010203, true, 3600 + 23*60}}, noMarshal: true},
		{str: "01:02:03-01", value: timeOfDayTzTest{TimeOfDay{time010203, true, -3600}}, noMarshal: true},
		{str: "01:02:03-01:23", value: timeOfDayTzTest{TimeOfDay{time010203, true, -(3600 + 23*60)}}},
		{str: "01:02:03-0123", value: timeOfDayTzTest{TimeOfDay{time010203, true, -(3600 + 23*60)}}, noMarshal: true},

		// dateTime
		{str: "2013-10-08T00:00:00", value: dateTimeTest{time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)}, tag: "no:dateTime.tz"},
		{str: "20131008", value: dateTimeTest{time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)}, noMarshal: true},
		{str: "2013-10-08T10:30:50", value: dateTimeTest{time.Date(2013, 10, 8, 10, 30, 50, 0, localLoc)}, tag: "no:dateTime.tz"},
		{str: "2013-10-08T10:30:50T", value: dateTimeTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "2013-10-08T10:30:50+01", value: dateTimeTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:dateTime.tz"},
		{str: "2013-10-08T10:30:50+01:23", value: dateTimeTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:dateTime.tz"},
		{str: "2013-10-08T10:30:50+0123", value: dateTimeTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:dateTime.tz"},
		{str: "2013-10-08T10:30:50-01", value: dateTimeTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:dateTime.tz"},
		{str: "2013-10-08T10:30:50-01:23", value: dateTimeTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:dateTime.tz"},
		{str: "2013-10-08T10:30:50-0123", value: dateTimeTest{}, wantUnmarshalErr: true, noMarshal: true, tag: "no:dateTime.tz"},

		// dateTime.tz
		{str: "2013-10-08T10:30:50", value: dateTimeTzTest{time.Date(2013, 10, 8, 10, 30, 50, 0, localLoc)}, noMarshal: true},
		{str: "2013-10-08T10:30:50+01", value: dateTimeTzTest{time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:00", 3600))}, noMarshal: true},
		{str: "2013-10-08T10:30:50+01:23", value: dateTimeTzTest{time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:23", 3600+23*60))}},
		{str: "2013-10-08T10:30:50+0123", value: dateTimeTzTest{time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:23", 3600+23*60))}, noMarshal: true},
		{str: "2013-10-08T10:30:50-01", value: dateTimeTzTest{time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:00", -3600))}, noMarshal: true},
		{str: "2013-10-08T10:30:50-01:23", value: dateTimeTzTest{time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:23", -(3600+23*60)))}},
		{str: "2013-10-08T10:30:50-0123", value: dateTimeTzTest{time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:23", -(3600+23*60)))}, noMarshal: true},

		// boolean
		{str: "0", value: booleanTest(false)},
		{str: "1", value: booleanTest(true)},
		{str: "false", value: booleanTest(false), noMarshal: true},
		{str: "true", value: booleanTest(true), noMarshal: true},
		{str: "no", value: booleanTest(false), noMarshal: true},
		{str: "yes", value: booleanTest(true), noMarshal: true},
		{str: "", value: booleanTest(false), noMarshal: true, wantUnmarshalErr: true},
		{str: "other", value: booleanTest(false), noMarshal: true, wantUnmarshalErr: true},
		{str: "2", value: booleanTest(false), noMarshal: true, wantUnmarshalErr: true},
		{str: "-1", value: booleanTest(false), noMarshal: true, wantUnmarshalErr: true},

		// bin.base64
		{str: "", value: binBase64Test{}},
		{str: "YQ==", value: binBase64Test("a")},
		{str: "TG9uZ2VyIFN0cmluZy4=", value: binBase64Test("Longer String.")},
		{str: "TG9uZ2VyIEFsaWduZWQu", value: binBase64Test("Longer Aligned.")},

		// bin.hex
		{str: "", value: binHexTest{}},
		{str: "61", value: binHexTest("a")},
		{str: "4c6f6e67657220537472696e672e", value: binHexTest("Longer String.")},
		{str: "4C6F6E67657220537472696E672E", value: binHexTest("Longer String."), noMarshal: true},

		// uri
		{str: "http://example.com/path", value: uriTest{&url.URL{Scheme: "http", Host: "example.com", Path: "/path"}}},
	}

	// Generate extra test cases from convTests that implement duper.
	var extras []testCase
	for i := range tests {
		if duper, ok := tests[i].value.(duper); ok {
			dupes := duper.dupe(tests[i].tag)
			for _, duped := range dupes {
				dupedCase := testCase(tests[i])
				dupedCase.value = duped
				extras = append(extras, dupedCase)
			}
		}
	}
	tests = append(tests, extras...)

	for _, test := range tests {
		if !test.noMarshal {
			resultStr, err := test.value.marshal()
			if err != nil && !test.wantMarshalErr {
				t.Errorf(
					"For %T marshal %v, want %q, got error: %v",
					test.value, test.value, test.str, err,
				)
			} else if err == nil && test.wantMarshalErr {
				t.Errorf(
					"For %T marshal %v, want error, got %q",
					test.value, test.value, resultStr,
				)
			} else if err == nil && resultStr != test.str {
				t.Errorf(
					"For %T marshal %v, want %q, got %q",
					test.value, test.value, test.str, resultStr,
				)
			}
		}

		if !test.noUnMarshal {
			resultValue, err := test.value.unmarshal(test.str)
			if err != nil && !test.wantUnmarshalErr {
				t.Errorf(
					"For %T unmarshal %q, want %v, got error: %v",
					test.value, test.str, test.value, err,
				)
			} else if err == nil && test.wantUnmarshalErr {
				t.Errorf(
					"For %T unmarshal %q, want error, got %v",
					test.value, test.str, resultValue,
				)
			} else if err == nil && !test.value.equal(resultValue) {
				t.Errorf(
					"For %T unmarshal %q, want %v, got %v",
					test.value, test.str, test.value, resultValue,
				)
			}
		}
	}
}
