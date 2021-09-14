package envelope

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestWriteRead(t *testing.T) {
	type Args struct {
		Foo string `xml:"foo"`
		Bar string `xml:"bar"`
	}

	sendAction := &Action{
		XMLName: xml.Name{Space: "http://example.com/namespace", Local: "MyAction"},
		Args: &Args{
			Foo: "foo-1",
			Bar: "bar-2",
		},
	}

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
