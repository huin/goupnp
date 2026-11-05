package httpu

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"
)

func TestClientReuseAfterTimeout(t *testing.T) {
	server, err := net.ListenPacket("udp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	if err := server.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
		t.Fatal(err)
	}
	serverDone := make(chan error, 1)
	go func() {
		var packet [2048]byte
		// The first request has a 20ms deadline and must time out.
		// Drain it without replying, then answer the reused client request.
		for i := 0; i < 2; i++ {
			_, peer, readErr := server.ReadFrom(packet[:])
			if readErr != nil {
				serverDone <- readErr
				return
			}
			if i == 1 {
				_, writeErr := server.WriteTo([]byte("HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n"), peer)
				serverDone <- writeErr
			}
		}
	}()

	client, err := NewHTTPUClientAddr("127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client.conn = &cancelOnResponseConn{PacketConn: client.conn, cancel: cancel}
	req := &http.Request{Method: "M-SEARCH", Host: server.LocalAddr().String(), URL: &url.URL{Opaque: "*"}}
	responses, err := client.Do(req, 20*time.Millisecond, 1)
	if err != nil || len(responses) != 0 {
		t.Fatalf("first request = %v, %v; want no responses and no error", responses, err)
	}
	// The next request has cancellation but no deadline. It must not inherit
	// the expired socket deadline from Do.
	watchdog := time.AfterFunc(time.Second, cancel)
	defer watchdog.Stop()
	responses, err = client.DoWithContext(req.WithContext(ctx), 1)
	if err != nil || len(responses) != 1 {
		t.Fatalf("reused request = %v, %v; want one response and no error", responses, err)
	}
	responses[0].Body.Close()
	if err := <-serverDone; err != nil {
		t.Fatal(err)
	}
}

type cancelOnResponseConn struct {
	net.PacketConn
	cancel context.CancelFunc
}

func (c *cancelOnResponseConn) ReadFrom(p []byte) (int, net.Addr, error) {
	n, peer, err := c.PacketConn.ReadFrom(p)
	if n > 0 {
		c.cancel()
	}
	return n, peer, err
}

func TestRequestWaitsForCancellationWatcher(t *testing.T) {
	client, err := NewHTTPUClientAddr("127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	conn := &delayedDeadlineConn{
		PacketConn: client.conn,
		reading:    make(chan struct{}),
		setting:    make(chan struct{}),
		release:    make(chan struct{}),
	}
	client.conn = conn
	var release sync.Once
	defer release.Do(func() { close(conn.release) })
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() {
		req := (&http.Request{Host: "127.0.0.1:1900", URL: &url.URL{Opaque: "*"}}).WithContext(ctx)
		_, requestErr := client.DoWithContext(req, 0)
		done <- requestErr
	}()
	select {
	case <-conn.reading:
	case <-time.After(time.Second):
		t.Fatal("request did not start reading")
	}
	cancel()
	select {
	case <-conn.setting:
	case <-time.After(time.Second):
		t.Fatal("cancellation did not set the deadline")
	}
	returned := false
	select {
	case err := <-done:
		t.Errorf("request returned before its cancellation watcher finished: %v", err)
		returned = true
	case <-time.After(50 * time.Millisecond):
	}
	release.Do(func() { close(conn.release) })
	if returned {
		return
	}
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("request did not finish after its cancellation watcher")
	}
}

type delayedDeadlineConn struct {
	net.PacketConn
	reading chan struct{}
	setting chan struct{}
	release chan struct{}
}

func (c *delayedDeadlineConn) ReadFrom(p []byte) (int, net.Addr, error) {
	close(c.reading)
	return c.PacketConn.ReadFrom(p)
}

func (c *delayedDeadlineConn) SetDeadline(deadline time.Time) error {
	err := c.PacketConn.SetDeadline(deadline)
	if !deadline.IsZero() && deadline.Before(time.Now()) {
		close(c.setting)
		<-c.release
	}
	return err
}
