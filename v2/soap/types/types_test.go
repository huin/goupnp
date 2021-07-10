package types

import (
	"bytes"
	"fmt"
	"math"
	"testing"
	"time"
)

var dummyLoc = time.FixedZone("DummyTZ", 6*3600)

func newFixed14_4Parts(intPart int64, fracPart int16) *Fixed14_4 {
	v, err := Fixed14_4FromParts(intPart, fracPart)
	if err != nil {
		panic(err)
	}
	return &v
}

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
			makeValue: func() SOAPValue { return &Fixed14_4{} },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*Fixed14_4).Fractional == want.(*Fixed14_4).Fractional
			},
			marshalTests: []marshalCase{
				{newFixed14_4Parts(0, 0), "0.0000"},
				{newFixed14_4Parts(1, 2), "1.0002"},
				{newFixed14_4Parts(1, 20), "1.0020"},
				{newFixed14_4Parts(1, 200), "1.0200"},
				{newFixed14_4Parts(1, 2000), "1.2000"},
				{newFixed14_4Parts(-1, -2), "-1.0002"},
				{newFixed14_4Parts(1234, 5678), "1234.5678"},
				{newFixed14_4Parts(-1234, -5678), "-1234.5678"},
				{newFixed14_4Parts(9999_99999_99999, 9999), "99999999999999.9999"},
				{newFixed14_4Parts(-9999_99999_99999, -9999), "-99999999999999.9999"},
			},
			unmarshalErrs: append([]string{
				"", ".", "0.00000000abc", "0.-5",
			}, badNumbers...),
			unmarshalTests: []unmarshalCase{
				{"010", newFixed14_4Parts(10, 0)},
				{"0", newFixed14_4Parts(0, 0)},
				{"0.", newFixed14_4Parts(0, 0)},
				{"0.000005", newFixed14_4Parts(0, 0)},
				{"1.2", newFixed14_4Parts(1, 2000)},
				{"1.20", newFixed14_4Parts(1, 2000)},
				{"1.200", newFixed14_4Parts(1, 2000)},
				{"1.02", newFixed14_4Parts(1, 200)},
				{"1.020", newFixed14_4Parts(1, 200)},
				{"1.002", newFixed14_4Parts(1, 20)},
				{"1.00200005", newFixed14_4Parts(1, 20)},
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
			makeValue: func() SOAPValue { return new(TimeOfDay) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*TimeOfDay).equal(*want.(*TimeOfDay))
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
				"01:02:03+01:23",
				"01:02:03+01:23",
				"01:02:03-01:23",
				"01:02:03-01:23",
			},
		},

		{
			makeValue: func() SOAPValue { return new(TimeOfDayTZ) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*TimeOfDayTZ).equal(*want.(*TimeOfDayTZ))
			},
			marshalTests: []marshalCase{
				{&TimeOfDayTZ{}, "00:00:00"},
				// ISO8601 special case
				{&TimeOfDayTZ{TimeOfDay{24, 0, 0}, TZD{}}, "24:00:00"},
				{&TimeOfDayTZ{TimeOfDay{1, 2, 3}, TZDOffset(0)}, "01:02:03Z"},
				{&TimeOfDayTZ{TimeOfDay{1, 2, 3}, TZDOffset(3600 + 23*60)}, "01:02:03+01:23"},
				{&TimeOfDayTZ{TimeOfDay{1, 2, 3}, TZDOffset(-(3600 + 23*60))}, "01:02:03-01:23"},
			},
			unmarshalTests: []unmarshalCase{
				{"010203+01:23", &TimeOfDayTZ{TimeOfDay{1, 2, 3}, TZDOffset(3600 + 23*60)}},
				{"010203-01:23", &TimeOfDayTZ{TimeOfDay{1, 2, 3}, TZDOffset(-(3600 + 23*60))}},
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
			makeValue: func() SOAPValue { return new(Date) },
			isEqual: func(got, want SOAPValue) bool {
				a, b := got.(*Date), want.(*Date)
				return a.Year == b.Year && a.Month == b.Month && a.Day == b.Day
			},
			marshalTests: []marshalCase{
				{&Date{2013, 10, 8}, "2013-10-08"},
			},
			unmarshalTests: []unmarshalCase{
				{"20131008", &Date{2013, 10, 8}},
			},
			unmarshalErrs: []string{"", "-1"},
		},

		{
			makeValue: func() SOAPValue { return new(DateTime) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*DateTime).equal(*want.(*DateTime))
			},
			marshalTests: []marshalCase{
				{DateTimeFromTime(time.Date(2013, 10, 8, 0, 0, 0, 0, dummyLoc)).ptr(), "2013-10-08T00:00:00"},
				{DateTimeFromTime(time.Date(2013, 10, 8, 10, 30, 50, 0, dummyLoc)).ptr(), "2013-10-08T10:30:50"},
			},
			unmarshalTests: []unmarshalCase{
				{"20131008", DateTimeFromTime(time.Date(2013, 10, 8, 0, 0, 0, 0, dummyLoc)).ptr()},
			},
			unmarshalErrs: []string{
				// Unexpected timezone component.
				"2013-10-08T10:30:50+01:00",
			},
		},

		{
			makeValue: func() SOAPValue { return new(DateTimeTZ) },
			isEqual: func(got, want SOAPValue) bool {
				return got.(*DateTimeTZ).equal(*want.(*DateTimeTZ))
			},
			marshalTests: []marshalCase{
				{DateTimeTZFromTime(time.Date(2013, 10, 8, 0, 0, 0, 0, dummyLoc)).ptr(), "2013-10-08T00:00:00+06:00"},
				{DateTimeTZFromTime(time.Date(2013, 10, 8, 10, 30, 50, 0, dummyLoc)).ptr(), "2013-10-08T10:30:50+06:00"},
				{DateTimeTZFromTime(time.Date(2013, 10, 8, 0, 0, 0, 0, time.UTC)).ptr(), "2013-10-08T00:00:00Z"},
				{DateTimeTZFromTime(time.Date(2013, 10, 8, 10, 30, 50, 0, time.UTC)).ptr(), "2013-10-08T10:30:50Z"},
				{DateTimeTZFromTime(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:23", 3600+23*60))).ptr(), "2013-10-08T10:30:50+01:23"},
				{DateTimeTZFromTime(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:23", -(3600+23*60)))).ptr(), "2013-10-08T10:30:50-01:23"},
			},
			unmarshalTests: []unmarshalCase{
				{"2013-10-08T10:30:50", &DateTimeTZ{Date{2013, 10, 8}, TimeOfDay{10, 30, 50}, TZD{}}},
				{"2013-10-08T10:30:50+00:00", DateTimeTZFromTime(time.Date(2013, 10, 8, 10, 30, 50, 0, time.UTC)).ptr()},
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
						t.Errorf("got unexpected error: %v", err)
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

func TestFixed14_4(t *testing.T) {
	t.Run("Parts", func(t *testing.T) {
		tests := []struct {
			intPart    int64
			fracPart   int16
			fractional int64
		}{
			{0, 0, 0},
			{1, 2, 1_0002},
			{-1, -2, -1_0002},
			{1234, 5678, 1234_5678},
			{-1234, -5678, -1234_5678},
			{9999_99999_99999, 9999, 9999_99999_99999_9999},
			{-9999_99999_99999, -9999, -9999_99999_99999_9999},
		}
		for _, test := range tests {
			test := test
			t.Run(fmt.Sprintf("FromParts(%d,%d)", test.intPart, test.fracPart), func(t *testing.T) {
				got, err := Fixed14_4FromParts(test.intPart, test.fracPart)
				if err != nil {
					t.Errorf("got error %v, want success", err)
				}
				if got.Fractional != test.fractional {
					t.Errorf("got %d, want %d", got.Fractional, test.fractional)
				}
			})
			t.Run(fmt.Sprintf("%d.Parts()", test.fractional), func(t *testing.T) {
				v, err := Fixed14_4FromFractional(test.fractional)
				if err != nil {
					t.Errorf("got error %v, want success", err)
				}
				gotIntPart, gotFracPart := v.Parts()
				if gotIntPart != test.intPart {
					t.Errorf("got %d, want %d", gotIntPart, test.intPart)
				}
				if gotFracPart != test.fracPart {
					t.Errorf("got %d, want %d", gotFracPart, test.fracPart)
				}
			})
		}
	})
	t.Run("Float", func(t *testing.T) {
		tests := []struct {
			flt float64
			fix *Fixed14_4
		}{
			{0, newFixed14_4Parts(0, 0)},
			{1234.5678, newFixed14_4Parts(1234, 5678)},
			{-1234.5678, newFixed14_4Parts(-1234, -5678)},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("To/FromFloat(%v)", test.fix), func(t *testing.T) {
				gotFix, err := Fixed14_4FromFloat(test.flt)
				if err != nil {
					t.Errorf("got error %v, want success", err)
				}
				if gotFix.Fractional != test.fix.Fractional {
					t.Errorf("got %v, want %v", gotFix, test.fix)
				}

				gotFlt := test.fix.Float64()
				if math.Abs(gotFlt-test.flt) > 1e-6 {
					t.Errorf("got %f, want %f", gotFlt, test.flt)
				}
			})
		}

		errTests := []float64{
			1e50,
			-1e50,
			1e14,
			-1e14,
			math.NaN(),
			math.Inf(1),
			math.Inf(-1),
		}
		for _, test := range errTests {
			t.Run(fmt.Sprintf("ErrorFromFloat(%f)", test), func(t *testing.T) {
				got, err := Fixed14_4FromFloat(test)
				if err == nil {
					t.Errorf("got success and %v, want error", got)
				}
			})
		}
	})
}

// methods only used in testing:

func (v TimeOfDay) equal(o TimeOfDay) bool {
	return v.Hour == o.Hour && v.Minute == o.Minute && v.Second == o.Second
}

func (v TimeOfDayTZ) equal(o TimeOfDayTZ) bool {
	return v.TimeOfDay.equal(o.TimeOfDay) && v.TZ.equal(o.TZ)
}

func (d Date) equal(o Date) bool {
	return d.Year == o.Year && d.Month == o.Month && d.Day == o.Day
}

func (dtz DateTime) ptr() *DateTime { return &dtz }

func (dt DateTime) equal(o DateTime) bool {
	return dt.Date.equal(o.Date) && dt.TimeOfDay.equal(o.TimeOfDay)
}

func (dtz DateTimeTZ) ptr() *DateTimeTZ { return &dtz }

func (dtz DateTimeTZ) equal(o DateTimeTZ) bool {
	return dtz.Date.equal(o.Date) &&
		dtz.TimeOfDay.equal(o.TimeOfDay) &&
		dtz.TZ.equal(o.TZ)
}

func (tzd TZD) equal(o TZD) bool {
	return tzd.Offset == o.Offset && tzd.HasTZ == o.HasTZ
}
