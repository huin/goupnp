package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/huin/goupnp/v2alpha/soap/envelope"
)

type ActionArgs struct {
	Name string
}
type ActionReply struct {
	Greeting string
}

type actionKey struct {
	endpointURL string
	action      string
}

var _ http.Handler = &fakeSoapServer{}

type fakeSoapServer struct {
	responses map[actionKey]*envelope.Action
	errors    []error
}

func (fss *fakeSoapServer) badRequest(w http.ResponseWriter, err error) {
	fss.errors = append(fss.errors, err)
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(err.Error()))
}

func (fss *fakeSoapServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fss.badRequest(w, fmt.Errorf("want POST, got %q", r.Method))
		return
	}
	actions := r.Header.Values("SOAPACTION")
	if len(actions) != 1 {
		fss.badRequest(w, fmt.Errorf("want exactly 1 SOAPACTION, got %d: %q", len(actions), actions))
		return
	}
	headerAction := actions[0]
	key := actionKey{
		endpointURL: r.URL.Path,
		action:      headerAction,
	}
	response, ok := fss.responses[key]
	if !ok {
		fss.badRequest(w, fmt.Errorf("no response known for %#v", key))
		return
	}

	reqArgs := &ActionArgs{}
	reqAction := envelope.Action{Args: reqArgs}
	if err := envelope.Read(r.Body, &reqAction); err != nil {
		fss.badRequest(w, fmt.Errorf("reading envelope from request: %w", err))
		return
	}
	envelopeAction := fmt.Sprintf("\"%s#%s\"", reqAction.XMLName.Space, reqAction.XMLName.Local)
	if envelopeAction != headerAction {
		fss.badRequest(w, fmt.Errorf("mismatch in header/envelope action: %q/%q", headerAction, envelopeAction))
		return
	}

	w.Header().Add("CONTENT-TYPE", `text/xml; charset="utf-8"`)
	if err := envelope.Write(w, response); err != nil {
		fss.errors = append(fss.errors, fmt.Errorf("writing envelope: %w", err))
	}
}

func TestDo(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	t.Cleanup(cancel)

	service := &fakeSoapServer{
		responses: map[actionKey]*envelope.Action{
			{"/endpointpath", "\"http://example.com/endpointns#Foo\""}: {
				Args: &ActionReply{Greeting: "Hello, World!"},
			},
		},
	}
	ts := httptest.NewServer(service)
	t.Cleanup(ts.Close)

	c := New(ts.URL + "/endpointpath")

	reqAction := envelope.NewAction("http://example.com/endpointns", "Foo",
		&ActionArgs{Name: "World"})
	reply := &ActionReply{}
	replyAction := &envelope.Action{Args: reply}

	if err := c.Do(ctx, reqAction, replyAction); err != nil {
		t.Errorf("got error: %v, want success", err)
	} else {
		if got, want := reply.Greeting, "Hello, World!"; got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	for _, err := range service.errors {
		t.Errorf("Service error: %v", err)
	}
}
