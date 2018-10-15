package httpu

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"net/http"
	"regexp"
)

const (
	defaultMaxMessageBytes = 2048
)

var (
	trailingWhitespaceRx = regexp.MustCompile(" +\r\n")
	crlf                 = []byte("\r\n")
)

// Handler is the interface by which received HTTPU messages are passed to
// handling code.
type Handler interface {
	// ServeMessage is called for each HTTPU message received. peerAddr contains
	// the address that the message was received from.
	ServeMessage(r *http.Request)
}

// HandlerFunc is a function-to-Handler adapter.
type HandlerFunc func(r *http.Request)

// ServeMessage implements Handler.
func (f HandlerFunc) ServeMessage(r *http.Request) {
	f(r)
}

// A Server defines parameters for running an HTTPU server.
type Server struct {
	// UDP address to listen on
	Addr string
	// Should listen for multicast?
	Multicast bool
	// Network interface to listen on for multicast, nil for default multicast
	// interface.
	Interface *net.Interface
	// Handler to invoke.
	Handler Handler
	// Maximum number of bytes to read from a packet, defaultMaxMessageBytes if
	// 0.
	MaxMessageBytes int
}

// ListenAndServe listens on the UDP network address srv.Addr. If srv.Multicast
// is true, then a multicast UDP listener will be used on srv.Interface (or
// default interface if nil).
func (srv *Server) ListenAndServe() error {
	var err error

	var addr *net.UDPAddr
	if addr, err = net.ResolveUDPAddr("udp", srv.Addr); err != nil {
		log.Fatal(err)
	}

	var conn net.PacketConn
	if srv.Multicast {
		conn, err = net.ListenMulticastUDP("udp", srv.Interface, addr)
		if err != nil {
			return err
		}
	} else {
		conn, err = net.ListenUDP("udp", addr)
		if err != nil {
			return err
		}
	}

	return srv.Serve(conn)
}

// Serve messages received on the given packet listener to the srv.Handler.
func (srv *Server) Serve(l net.PacketConn) error {
	maxMessageBytes := defaultMaxMessageBytes
	if srv.MaxMessageBytes != 0 {
		maxMessageBytes = srv.MaxMessageBytes
	}
	for {
		buf := make([]byte, maxMessageBytes)
		n, peerAddr, err := l.ReadFrom(buf)
		if err != nil {
			return err
		}
		buf = buf[:n]

		go func(buf []byte, peerAddr net.Addr) {
			// At least one router's UPnP implementation has added a trailing
			// space after "HTTP/1.1" - trim it.
			buf = trailingWhitespaceRx.ReplaceAllLiteral(buf, crlf)

			req, err := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(buf)))
			if err != nil {
				log.Printf("httpu: Failed to parse request: %v", err)
				return
			}
			req.RemoteAddr = peerAddr.String()
			srv.Handler.ServeMessage(req)
			// No need to call req.Body.Close - underlying reader is
			// bytes.Buffer.
		}(buf, peerAddr)
	}
}

// Serve messages received on the given packet listener to the given handler.
func Serve(l net.PacketConn, handler Handler) error {
	srv := Server{
		Handler:         handler,
		MaxMessageBytes: defaultMaxMessageBytes,
	}
	return srv.Serve(l)
}
