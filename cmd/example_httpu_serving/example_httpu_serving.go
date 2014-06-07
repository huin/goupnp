package main

import (
	"log"
	"net"
	"net/http"

	"github.com/huin/goupnp/httpu"
)

func main() {
	eth0, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatal(err)
	}
	srv := httpu.Server{
		Addr:      "239.255.255.250:1900",
		Interface: eth0,
		Handler: httpu.HandlerFunc(func(r *http.Request) {
			log.Printf("Got %s %s message from %v: %v", r.Method, r.URL.Path, r.RemoteAddr, r.Header)
		}),
	}
	err = srv.ListenAndServe()
	log.Printf("Serving failed with error: %v", err)
}
