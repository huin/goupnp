package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/huin/goupnp/v2alpha/soap"
	"github.com/huin/goupnp/v2alpha/soap/envelope"
)

const serviceType = "fake:service:type"
const actionName = "ActionName"

type Action struct {
	req  ActionArgs
	resp ActionReply
}

var _ soap.Action = &Action{}

func (a *Action) ServiceType() string { return serviceType }
func (a *Action) ActionName() string  { return actionName }
func (a *Action) RefRequest() any     { return &a.req }
func (a *Action) RefResponse() any    { return &a.resp }

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

func TestPerformAction(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	t.Cleanup(cancel)

	service := &fakeSoapServer{
		responses: map[actionKey]*envelope.Action{
			{"/endpointpath", fmt.Sprintf("\"%s#%s\"", serviceType, actionName)}: {
				Args: &ActionReply{Greeting: "Hello, World!"},
			},
		},
	}
	ts := httptest.NewServer(service)
	t.Cleanup(ts.Close)

	c := New(ts.URL + "/endpointpath")

	a := &Action{
		req: ActionArgs{Name: "World"},
	}

	if err := PerformAction(ctx, c, a); err != nil {
		t.Errorf("got error: %v, want success", err)
	} else {
		if got, want := a.resp.Greeting, "Hello, World!"; got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	for _, err := range service.errors {
		t.Errorf("Service error: %v", err)
	}
}
