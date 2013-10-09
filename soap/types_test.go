package soap

import (
	"math"
	"testing"
	"time"
)

type convTest interface {
	Marshal() (string, error)
	Unmarshal(string) (interface{}, error)
	Equal(result interface{}) bool
}

type testCase struct {
	value            convTest
	str              string
	wantMarshalErr   bool
	wantUnmarshalErr bool
	noMarshal        bool
	noUnMarshal      bool
}

type Fixed14_4Test float64

func (v Fixed14_4Test) Marshal() (string, error) {
	return MarshalFixed14_4(float64(v))
}
func (v Fixed14_4Test) Unmarshal(s string) (interface{}, error) {
	return UnmarshalFixed14_4(s)
}
func (v Fixed14_4Test) Equal(result interface{}) bool {
	return math.Abs(float64(v)-result.(float64)) < 0.001
}

type CharTest rune

func (v CharTest) Marshal() (string, error) {
	return MarshalChar(rune(v))
}
func (v CharTest) Unmarshal(s string) (interface{}, error) {
	return UnmarshalChar(s)
}
func (v CharTest) Equal(result interface{}) bool {
	return rune(v) == result.(rune)
}

type DateTest struct{ time.Time }

func (v DateTest) Marshal() (string, error) {
	return MarshalDate(time.Time(v.Time))
}
func (v DateTest) Unmarshal(s string) (interface{}, error) {
	return UnmarshalDate(s)
}
func (v DateTest) Equal(result interface{}) bool {
	return v.Time.Equal(result.(time.Time))
}

type TimeOfDayTest struct {
	TimeOfDay
}

func (v TimeOfDayTest) Marshal() (string, error) {
	return MarshalTimeOfDay(v.TimeOfDay)
}
func (v TimeOfDayTest) Unmarshal(s string) (interface{}, error) {
	return UnmarshalTimeOfDay(s)
}
func (v TimeOfDayTest) Equal(result interface{}) bool {
	return v.TimeOfDay == result.(TimeOfDay)
}

func Test(t *testing.T) {
	tests := []testCase{
		// Fixed14_4
		{str: "0.0000", value: Fixed14_4Test(0)},
		{str: "1.0000", value: Fixed14_4Test(1)},
		{str: "1.2346", value: Fixed14_4Test(1.23456)},
		{str: "-1.0000", value: Fixed14_4Test(-1)},
		{str: "-1.2346", value: Fixed14_4Test(-1.23456)},
		{str: "10000000000000.0000", value: Fixed14_4Test(1e13)},
		{str: "100000000000000.0000", value: Fixed14_4Test(1e14), wantMarshalErr: true, wantUnmarshalErr: true},
		{str: "-10000000000000.0000", value: Fixed14_4Test(-1e13)},
		{str: "-100000000000000.0000", value: Fixed14_4Test(-1e14), wantMarshalErr: true, wantUnmarshalErr: true},

		// Char
		{str: "a", value: CharTest('a')},
		{str: "z", value: CharTest('z')},
		{str: "\u1234", value: CharTest(0x1234)},
		{str: "aa", value: CharTest(0), wantMarshalErr: true, wantUnmarshalErr: true},
		{str: "", value: CharTest(0), wantMarshalErr: true, wantUnmarshalErr: true},

		// Date
		{str: "2013-10-08", value: DateTest{time.Date(2013, 10, 8, 0, 0, 0, 0, time.Local)}},
		{str: "20131008", value: DateTest{time.Date(2013, 10, 8, 0, 0, 0, 0, time.Local)}, noMarshal: true},
		{str: "2013-10-08T10:30:50", value: DateTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "", value: DateTest{}, wantMarshalErr: true, wantUnmarshalErr: true, noMarshal: true},
		{str: "-1", value: DateTest{}, wantUnmarshalErr: true, noMarshal: true},

		// Time
		{str: "00:00:00", value: TimeOfDayTest{TimeOfDay{0, nil}}},
		{str: "000000", value: TimeOfDayTest{TimeOfDay{0, nil}}, noMarshal: true},
		{str: "01:02:03", value: TimeOfDayTest{TimeOfDay{(1*3600 + 2*60 + 3) * time.Second, nil}}},
		{str: "010203", value: TimeOfDayTest{TimeOfDay{(1*3600 + 2*60 + 3) * time.Second, nil}}, noMarshal: true},
		{str: "23:59:59", value: TimeOfDayTest{TimeOfDay{(23*3600 + 59*60 + 59) * time.Second, nil}}},
		{str: "235959", value: TimeOfDayTest{TimeOfDay{(23*3600 + 59*60 + 59) * time.Second, nil}}, noMarshal: true},
		{str: "01:02", value: TimeOfDayTest{TimeOfDay{(1*3600 + 2*60) * time.Second, nil}}, noMarshal: true},
		{str: "0102", value: TimeOfDayTest{TimeOfDay{(1*3600 + 2*60) * time.Second, nil}}, noMarshal: true},
		{str: "01", value: TimeOfDayTest{TimeOfDay{(1 * 3600) * time.Second, nil}}, noMarshal: true},
		{str: "01", value: TimeOfDayTest{TimeOfDay{(1 * 3600) * time.Second, nil}}, noMarshal: true},
		{str: "foo 01:02:03", value: TimeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "foo\n01:02:03", value: TimeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "01:02:03 foo", value: TimeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true},
		{str: "01:02:03\nfoo", value: TimeOfDayTest{}, wantUnmarshalErr: true, noMarshal: true},
	}

	for _, test := range tests {
		if test.noMarshal {
		} else if resultStr, err := test.value.Marshal(); err != nil && !test.wantMarshalErr {
			t.Errorf("For %s, want %q, got error: %v", test.value, test.str, err)
		} else if err == nil && test.wantMarshalErr {
			t.Errorf("For %s, want error, got %q", test.value, resultStr)
		} else if err == nil && resultStr != test.str {
			t.Errorf("For %s, want %q, got %q", test.value, test.str, resultStr)
		}

		if test.noUnMarshal {
		} else if resultValue, err := test.value.Unmarshal(test.str); err != nil && !test.wantUnmarshalErr {
			t.Errorf("For %q, want %v, got error: %v", test.str, test.value, err)
		} else if err == nil && test.wantUnmarshalErr {
			t.Errorf("For %q, want error, got %v", test.str, resultValue)
		} else if err == nil && !test.value.Equal(resultValue) {
			t.Errorf("For %q, want %v, got %v", test.str, test.value, resultValue)
		}
	}
}
