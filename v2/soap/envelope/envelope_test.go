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
		t.Errorf("EncodeEnvelope want success, got err=%v", err)
	}

	recvAction := &Action{Args: &Args{}}

	err = Read(buf, recvAction)
	if err != nil {
		t.Errorf("Reading envelope want success, got err=%v", err)
	}

	if !reflect.DeepEqual(sendAction, recvAction) {
		t.Errorf("want recvAction=%+v, got %+v", sendAction, recvAction)
	}
}
