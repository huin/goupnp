package types

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

type isEqual func(got, want SOAPValue) bool

type typeTestCase struct {
	makeValue      func() SOAPValue
	isEqual        isEqual
	marshalTests   []marshalCase
	marshalErrs    []SOAPValue
	unmarshalTests []unmarshalCase
	unmarshalErrs  []string
}

type marshalCase struct {
	input SOAPValue
	want  string
}

type unmarshalCase struct {
	input string
	want  SOAPValue
}

func Test(t *testing.T) {
	// Fake out the local time for the implementation.
	localLoc = time.FixedZone("Fake/Local", 6*3600)
	defer func() {
		localLoc = time.Local
	}()

	badNumbers := []string{"", " ", "abc"}

	typeTestCases := []typeTestCase{
		{
			makeValue: func() SOAPValue { return new(UI1) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*UI1) == *want.(*UI1) },
			marshalTests: []marshalCase{
				{NewUI1(0), "0"},
				{NewUI1(1), "1"},
				{NewUI1(255), "255"},
			},
			unmarshalErrs: append([]string{"-1", "256"}, badNumbers...),
		},

		{
			makeValue: func() SOAPValue { return new(UI2) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*UI2) == *want.(*UI2) },
			marshalTests: []marshalCase{
				{NewUI2(0), "0"},
				{NewUI2(1), "1"},
				{NewUI2(65535), "65535"},
			},
			unmarshalErrs: append([]string{"-1", "65536"}, badNumbers...),
		},

		{
			makeValue: func() SOAPValue { return new(UI4) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*UI4) == *want.(*UI4) },
			marshalTests: []marshalCase{
				{NewUI4(0), "0"},
				{NewUI4(1), "1"},
				{NewUI4(4294967295), "4294967295"},
			},
			unmarshalErrs: append([]string{"-1", "4294967296"}, badNumbers...),
		},

		{
			makeValue: func() SOAPValue { return new(UI8) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*UI8) == *want.(*UI8) },
			marshalTests: []marshalCase{
				{NewUI8(0), "0"},
				{NewUI8(1), "1"},
				{NewUI8(18446744073709551615), "18446744073709551615"},
			},
			unmarshalErrs: append([]string{"-1", "18446744073709551616"}, badNumbers...),
		},

		{
			makeValue: func() SOAPValue { return new(I1) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*I1) == *want.(*I1) },
			marshalTests: []marshalCase{
				{NewI1(0), "0"},
				{NewI1(1), "1"},
				{NewI1(-1), "-1"},
				{NewI1(127), "127"},
				{NewI1(-128), "-128"},
			},
			unmarshalErrs: append([]string{"-129", "128"}, badNumbers...),
		},

		{
			makeValue: func() SOAPValue { return new(I2) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*I2) == *want.(*I2) },
			marshalTests: []marshalCase{
				{NewI2(0), "0"},
				{NewI2(1), "1"},
				{NewI2(-1), "-1"},
				{NewI2(32767), "32767"},
				{NewI2(-32768), "-32768"},
			},
			unmarshalErrs: append([]string{"-32769", "32768"}, badNumbers...),
		},

		{
			makeValue: func() SOAPValue { return new(I4) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*I4) == *want.(*I4) },
			marshalTests: []marshalCase{
				{NewI4(0), "0"},
				{NewI4(1), "1"},
				{NewI4(-1), "-1"},
				{NewI4(2147483647), "2147483647"},
				{NewI4(-2147483648), "-2147483648"},
			},
			unmarshalErrs: append([]string{"-2147483649", "2147483648"}, badNumbers...),
		},

		{
			makeValue: func() SOAPValue { return new(I8) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*I8) == *want.(*I8) },
			marshalTests: []marshalCase{
				{NewI8(0), "0"},
				{NewI8(1), "1"},
				{NewI8(-1), "-1"},
				{NewI8(9223372036854775807), "9223372036854775807"},
				{NewI8(-9223372036854775808), "-9223372036854775808"},
			},
			unmarshalErrs: append([]string{"-9223372036854775809", "9223372036854775808"}, badNumbers...),
		},

		{
			makeValue: func() SOAPValue { return new(FloatFixed14_4) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*FloatFixed14_4) == *want.(*FloatFixed14_4) },
			marshalTests: []marshalCase{
				{NewFloatFixed14_4(0), "0.0000"},
				{NewFloatFixed14_4(1), "1.0000"},
				{NewFloatFixed14_4(1.2346), "1.2346"},
				{NewFloatFixed14_4(-1), "-1.0000"},
				{NewFloatFixed14_4(-1.2346), "-1.2346"},
				{NewFloatFixed14_4(1e13), "10000000000000.0000"},
				{NewFloatFixed14_4(-1e13), "-10000000000000.0000"},
			},
			marshalErrs: []SOAPValue{
				NewFloatFixed14_4(1e14),
				NewFloatFixed14_4(-1e14),
			},
		},

		{
			makeValue: func() SOAPValue { return new(Char) },
			isEqual:   func(got, want SOAPValue) bool { return *got.(*Char) == *want.(*Char) },
			marshalTests: []marshalCase{
				{NewChar('a'), "a"},
				{NewChar('z'), "z"},
				{NewChar('\u1234'), "\u1234"},
			},
			marshalErrs:   []SOAPValue{NewChar(0)},
			unmarshalErrs: []string{"aa", ""},
		},

		{
			makeValue: func() SOAPValue { return new(DateLocal) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*DateLocal).ToTime().Equal(want.(*DateLocal).ToTime())
			},
			marshalTests: []marshalCase{
				{NewDateLocal(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)), "2013-10-08"},
			},
			unmarshalTests: []unmarshalCase{
				{"20131008", NewDateLocal(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc))},
			},
			unmarshalErrs: []string{"", "-1"},
		},

		{
			makeValue: func() SOAPValue { return new(TimeOfDay) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*TimeOfDay).Equal(want.(*TimeOfDay))
			},
			marshalTests: []marshalCase{
				{&TimeOfDay{}, "00:00:00"},
				// ISO8601 special case
				{&TimeOfDay{Hour: 24}, "24:00:00"},
			},
			unmarshalTests: []unmarshalCase{
				{"000000", &TimeOfDay{}},
			},
			unmarshalErrs: []string{
				// Misformatted values:
				"foo 01:02:03", "foo\n01:02:03", "01:02:03 foo", "01:02:03\nfoo", "01:02:03Z",
				"01:02:03+01", "01:02:03+01:23", "01:02:03+0123", "01:02:03-01", "01:02:03-01:23",
				"01:02:03-0123",
				// Values out of range:
				"24:01:00",
				"24:00:01",
				"25:00:00",
				"00:60:00",
				"00:00:60",
				// Unexpected timezone component:
				"01:02:03Z",
				"01:02:03+01",
				"01:02:03+01:23",
				"01:02:03+0123",
				"01:02:03-01",
				"01:02:03-01:23",
				"01:02:03-0123",
			},
		},

		{
			makeValue: func() SOAPValue { return new(TimeOfDayTZ) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*TimeOfDayTZ).Equal(want.(*TimeOfDayTZ))
			},
			marshalTests: []marshalCase{
				{&TimeOfDayTZ{}, "00:00:00"},
				// ISO8601 special case
				{&TimeOfDayTZ{TimeOfDay{24, 0, 0}, false, 0}, "24:00:00"},
				{&TimeOfDayTZ{TimeOfDay{1, 2, 3}, true, 0}, "01:02:03Z"},
				{&TimeOfDayTZ{TimeOfDay{1, 2, 3}, true, 3600 + 23*60}, "01:02:03+01:23"},
				{&TimeOfDayTZ{TimeOfDay{1, 2, 3}, true, -(3600 + 23*60)}, "01:02:03-01:23"},
			},
			unmarshalTests: []unmarshalCase{
				{"000000", &TimeOfDayTZ{}},
				{"01Z", &TimeOfDayTZ{TimeOfDay{1, 0, 0}, true, 0}},
				{"01+01", &TimeOfDayTZ{TimeOfDay{1, 0, 0}, true, 3600}},
				{"01:02:03+01", &TimeOfDayTZ{TimeOfDay{1, 2, 3}, true, 3600}},
				{"01:02:03+0123", &TimeOfDayTZ{TimeOfDay{1, 2, 3}, true, 3600 + 23*60}},
				{"01:02:03-01", &TimeOfDayTZ{TimeOfDay{1, 2, 3}, true, -3600}},
				{"01:02:03-0123", &TimeOfDayTZ{TimeOfDay{1, 2, 3}, true, -(3600 + 23*60)}},
			},
			unmarshalErrs: []string{
				// Misformatted values:
				"foo 01:02:03", "foo\n01:02:03", "01:02:03 foo", "01:02:03\nfoo",
				// Values out of range:
				"24:01:00",
				"24:00:01",
				"25:00:00",
				"00:60:00",
				"00:00:60",
			},
		},

		{
			makeValue: func() SOAPValue { return new(DateLocal) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*DateLocal).ToTime().Equal(want.(*DateLocal).ToTime())
			},
			marshalTests: []marshalCase{
				{NewDateLocal(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)), "2013-10-08"},
			},
			unmarshalTests: []unmarshalCase{
				{"20131008", NewDateLocal(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc))},
			},
			unmarshalErrs: []string{
				// Unexpected time component.
				"2013-10-08T10:30:50",
				// Unexpected timezone component.
				"2013-10-08+01",
			},
		},

		{
			makeValue: func() SOAPValue { return new(DateTimeLocal) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*DateTimeLocal).ToTime().Equal(want.(*DateTimeLocal).ToTime())
			},
			marshalTests: []marshalCase{
				{NewDateTimeLocal(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)), "2013-10-08T00:00:00"},
				{NewDateTimeLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, localLoc)), "2013-10-08T10:30:50"},
			},
			unmarshalTests: []unmarshalCase{
				{"20131008", NewDateTimeLocal(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc))},
			},
			unmarshalErrs: []string{
				// Unexpected timezone component.
				"2013-10-08T10:30:50+01",
			},
		},

		{
			makeValue: func() SOAPValue { return new(DateTimeTZLocal) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*DateTimeTZLocal).ToTime().Equal(want.(*DateTimeTZLocal).ToTime())
			},
			marshalTests: []marshalCase{
				{NewDateTimeTZLocal(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)), "2013-10-08T00:00:00+06:00"},
				{NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, localLoc)), "2013-10-08T10:30:50+06:00"},
				{NewDateTimeTZLocal(time.Date(2013, 10, 8, 0, 0, 0, 0, time.UTC)), "2013-10-08T00:00:00+00:00"},
				{NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, time.UTC)), "2013-10-08T10:30:50+00:00"},
				{NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:23", 3600+23*60))), "2013-10-08T10:30:50+01:23"},
				{NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:23", -(3600+23*60)))), "2013-10-08T10:30:50-01:23"},
			},
			unmarshalTests: []unmarshalCase{
				{"20131008", NewDateTimeTZLocal(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc))},
				{"2013-10-08T10:30:50", NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, localLoc))},
				{"2013-10-08T10:30:50Z", NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, time.UTC))},
				{"2013-10-08T10:30:50+01", NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:00", 3600)))},
				{"2013-10-08T10:30:50+0123", NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:23", 3600+23*60)))},
				{"2013-10-08T10:30:50-01", NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:00", -3600)))},
				{"2013-10-08T10:30:50-0123", NewDateTimeTZLocal(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:23", -(3600+23*60))))},
			},
		},

		{
			makeValue: func() SOAPValue { return new(Boolean) },
			isEqual: func(got, want SOAPValue) bool {
				return *got.(*Boolean) == *want.(*Boolean)
			},
			marshalTests: []marshalCase{
				{NewBoolean(true), "1"},
				{NewBoolean(false), "0"},
			},
			unmarshalTests: []unmarshalCase{
				{"true", NewBoolean(true)},
				{"false", NewBoolean(false)},
				{"yes", NewBoolean(true)},
				{"no", NewBoolean(false)},
			},
			unmarshalErrs: []string{"", "2", "-1"},
		},

		{
			makeValue: func() SOAPValue { return new(BinBase64) },
			isEqual: func(got, want SOAPValue) bool {
				return bytes.Equal(*got.(*BinBase64), *want.(*BinBase64))
			},
			marshalTests: []marshalCase{
				{&BinBase64{}, ""},
				{NewBinBase64([]byte("a")), "YQ=="},
				{NewBinBase64([]byte("Longer String.")), "TG9uZ2VyIFN0cmluZy4="},
				{NewBinBase64([]byte("Longer Aligned.")), "TG9uZ2VyIEFsaWduZWQu"},
			},
		},

		{
			makeValue: func() SOAPValue { return new(BinHex) },
			isEqual: func(got, want SOAPValue) bool {
				return bytes.Equal(*got.(*BinHex), *want.(*BinHex))
			},
			marshalTests: []marshalCase{
				{&BinHex{}, ""},
				{NewBinHex([]byte("a")), "61"},
				{NewBinHex([]byte("Longer String.")), "4c6f6e67657220537472696e672e"},
			},
			unmarshalTests: []unmarshalCase{
				{"4C6F6E67657220537472696E672E", NewBinHex([]byte("Longer String."))},
			},
		},

		{
			makeValue: func() SOAPValue { return new(URI) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*URI).ToURL().String() == want.(*URI).ToURL().String()
			},
			marshalTests: []marshalCase{
				{&URI{Scheme: "http", Host: "example.com", Path: "/path"}, "http://example.com/path"},
			},
		},
	}

	for _, tt := range typeTestCases {
		tt := tt

		// Convert marshalTests into additional round trip equivalent unmarshalTests
		for _, mt := range tt.marshalTests {
			tt.unmarshalTests = append(tt.unmarshalTests, unmarshalCase{
				input: mt.want,
				want:  mt.input,
			})
		}

		t.Run(fmt.Sprintf("%T", tt.makeValue()), func(t *testing.T) {
			for i, mt := range tt.marshalTests {
				mt := mt
				t.Run(fmt.Sprintf("marshalTest#%d_%v", i, mt.input), func(t *testing.T) {
					got, err := mt.input.Marshal()
					if err != nil {
						t.Errorf("got unexpected error: %v", err)
					}
					if got != mt.want {
						t.Errorf("got %q, want: %q", got, mt.want)
					}
				})
			}
			for i, input := range tt.marshalErrs {
				input := input
				t.Run(fmt.Sprintf("marshalErr#%d_%v", i, input), func(t *testing.T) {
					got, err := input.Marshal()
					if err == nil {
						t.Errorf("got %q, want error", got)
					}
				})
			}
			for i, ut := range tt.unmarshalTests {
				ut := ut
				t.Run(fmt.Sprintf("unmarshalTest#%d_%q", i, ut.input), func(t *testing.T) {
					got := tt.makeValue()
					if err := got.Unmarshal(ut.input); err != nil {
						t.Errorf("got error, want success")
					}
					if !tt.isEqual(got, ut.want) {
						t.Errorf("got %v, want %v", got, ut.want)
					}
				})
			}
			for i, input := range tt.unmarshalErrs {
				input := input
				t.Run(fmt.Sprintf("unmarshalErrs#%d_%q", i, input), func(t *testing.T) {
					got := tt.makeValue()
					if err := got.Unmarshal(input); err == nil {
						t.Errorf("got %v, want error", got)
					}
				})
			}
		})
	}
}
