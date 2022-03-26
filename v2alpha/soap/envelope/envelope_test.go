package envelope

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestWriteRead(t *testing.T) {
	type Args struct {
		Foo string `xml:"foo"`
		Bar string `xml:"bar"`
	}

	sendAction := NewAction("http://example.com/namespace", "MyAction", &Args{
		Foo: "foo-1",
		Bar: "bar-2",
	})

	buf := &bytes.Buffer{}
	err := Write(buf, sendAction)
	if err != nil {
		t.Errorf("Write want success, got err=%v", err)
	}

	recvAction := &Action{Args: &Args{}}

	err = Read(buf, recvAction)
	if err != nil {
		t.Errorf("Read want success, got err=%v", err)
	}

	if !reflect.DeepEqual(sendAction, recvAction) {
		t.Errorf("want recvAction=%+v, got %+v", sendAction, recvAction)
	}
}

func TestReadFault(t *testing.T) {
	env := []byte(xml.Header + `
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"
s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
<s:Body>
<s:Fault>
<faultcode>dummy code</faultcode>
<faultstring>dummy string</faultstring>
<faultactor>dummy actor</faultactor>
<detail>dummy detail</detail>
</s:Fault>
</s:Body>
</s:Envelope>
`)

	type args struct{}

	err := Read(bytes.NewBuffer(env), &Action{Args: &args{}})
	if err == nil {
		t.Fatal("want err != nil, got nil")
	}

	gotFault, ok := err.(*Fault)
	if !ok {
		t.Fatalf("want *Fault, got %T", err)
	}

	wantFault := &Fault{
		Code:   "dummy code",
		String: "dummy string",
		Actor:  "dummy actor",
		Detail: FaultDetail{Raw: []byte("dummy detail")},
	}
	if !reflect.DeepEqual(wantFault, gotFault) {
		t.Errorf("want %+v, got %+v", wantFault, gotFault)
	}
}

func TestFault(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		wantIs bool
	}{
		{"plain fault", &Fault{Code: "code"}, true},
		{"wrapped fault", fmt.Errorf("wrapper: %w", &Fault{Code: "code"}), true},
		{"other error", io.EOF, false},
	}

	for _, test := range tests {
		test := test // copy for closure
		t.Run(test.name, func(t *testing.T) {
			if got, want := errors.Is(test.err, ErrFault), test.wantIs; got != want {
				t.Errorf("got errors.Is(%v, ErrFault)=>%t, want %t", test.err, got, want)
			}
			if !test.wantIs {
				return
			}

			var fault *Fault
			if got, want := errors.As(test.err, &fault), true; got != want {
				t.Fatalf("got errors.As(%v, ...)=>%t, want %t", test.err, got, want)
			}

			if got, want := fault.Code, "code"; got != want {
				t.Errorf("got fault.Code=%q, want %q", got, want)
			}
		})
	}
}
