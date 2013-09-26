package goupnp

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"
	"time"
)

// TODO: RoundTripper is probably the wrong interface, as there could be
// multiple responses to a request.

type udpRoundTripper struct {
	// If zero, defaults to 3 second deadline (a zero deadline makes no sense).
	Deadline       time.Duration
	MaxWaitSeconds int
}

func (urt udpRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var requestBuf bytes.Buffer
	if err := r.Write(&requestBuf); err != nil {
		return nil, err
	}
	destAddr, err := net.ResolveUDPAddr("udp", r.Host)
	if err != nil {
		return nil, err
	}

	deadline := urt.Deadline
	if urt.Deadline == 0 {
		deadline = 3 * time.Second
	}

	if err = conn.SetDeadline(time.Now().Add(deadline)); err != nil {
		return nil, err
	}

	// Send request.
	if n, err := conn.WriteTo(requestBuf.Bytes(), destAddr); err != nil {
		return nil, err
	} else if n < len(requestBuf.Bytes()) {
		return nil, fmt.Errorf("goupnp: wrote %d bytes rather than full %d in request",
			n, len(requestBuf.Bytes()))
	}

	// Await response.
	responseBytes := make([]byte, 2048)
	n, _, err := conn.ReadFrom(responseBytes)
	if err != nil {
		return nil, err
	}
	responseBytes = responseBytes[:n]

	// Parse response.
	response, err := http.ReadResponse(bufio.NewReader(bytes.NewBuffer(responseBytes)), r)
	if err != nil {
		return nil, err
	}

	return response, err
}
