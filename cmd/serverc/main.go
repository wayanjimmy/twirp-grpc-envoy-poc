package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/pkg/pkgc"
	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcc"
	"github.com/twitchtv/twirp"
)

func main() {
	var cw ConnectionWatcher
	serviceHandler := svcc.NewCServiceServer(pkgc.NewSvcC(),
		twirp.WithServerPathPrefix("/rz"))
	httpServer := &http.Server{
		Addr:         ":3003",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		//TLSConfig:    tlsConfig(),
		ConnState: cw.OnStateChange,
	}
	httpServer.Handler = serviceHandler
	go func() {
		//err := httpServer.ListenAndServeTLS("../../server.crt", "../../server.key")
		err := httpServer.ListenAndServe()
		if err != nil {
			os.Exit(-1)
		}
	}()

	log.Println("Server C started at port :3003")
	select {}
}

type ConnectionWatcher struct {
	n int64
}

// OnStateChange records open connections in response to connection
// state changes. Set net/http Server.ConnState to this method
// as value.
func (cw *ConnectionWatcher) OnStateChange(conn net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		atomic.AddInt64(&cw.n, 1)
	case http.StateHijacked, http.StateClosed:
		atomic.AddInt64(&cw.n, -1)
	}
	log.Printf("connection count : %d \n", int(atomic.LoadInt64(&cw.n)))
}

// Count returns the number of connections at the time
// the call.
func (cw *ConnectionWatcher) Count() int {
	return int(atomic.LoadInt64(&cw.n))

}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("../../server.crt")
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile("../../server.key")
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}
}
