package httpu

import (
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestListenAndServeInvalidAddress(t *testing.T) {
	if os.Getenv("GOUPNP_INVALID_ADDRESS_HELPER") == "1" {
		for _, multicast := range []bool{false, true} {
			server := &Server{Addr: "127.0.0.1:invalid-port", Multicast: multicast}
			if err := server.ListenAndServe(); err == nil {
				t.Fatal("ListenAndServe returned no address error")
			}
		}
		return
	}
	// Run in a subprocess so a regression to log.Fatal cannot exit the
	// entire test suite instead of reporting a failed assertion.
	cmd := exec.Command(os.Args[0], "-test.run=^TestListenAndServeInvalidAddress$")
	cmd.Env = append(os.Environ(), "GOUPNP_INVALID_ADDRESS_HELPER=1")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("invalid address terminated the process: %v\n%s", err, output)
	}
}

func TestListenAndServeBindError(t *testing.T) {
	listener, err := net.ListenPacket("udp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	server := &Server{Addr: listener.LocalAddr().String()}
	if err := server.ListenAndServe(); err == nil {
		t.Fatal("ListenAndServe returned no bind error")
	}
}

func TestServeReceivesMessage(t *testing.T) {
	listener, err := net.ListenPacket("udp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	if err := listener.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	messages := make(chan string, 1)
	server := &Server{Handler: HandlerFunc(func(req *http.Request) { messages <- req.Header.Get("ST") })}
	done := make(chan error, 1)
	go func() { done <- server.Serve(listener) }()
	client, err := net.Dial("udp4", listener.LocalAddr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	if _, err := client.Write([]byte("M-SEARCH * HTTP/1.1\r\nST: test-device\r\n\r\n")); err != nil {
		t.Fatal(err)
	}
	select {
	case message := <-messages:
		if message != "test-device" {
			t.Fatalf("search target = %q; want test-device", message)
		}
	case <-time.After(time.Second):
		t.Fatal("HTTPU server did not receive message")
	}
	listener.Close()
	select {
	case err := <-done:
		if err == nil || !strings.Contains(err.Error(), "closed") {
			t.Fatalf("Serve returned %v; want closed connection error", err)
		}
	case <-time.After(time.Second):
		t.Fatal("Serve did not return after closing listener")
	}
}
