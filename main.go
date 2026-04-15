package main

import (
	"flag"
	"net/http"

	"github.com/pretty66/websocketproxy"
)

var (
	backend = flag.String("backend", "", "Backend URL for proxying")
	port    = flag.String("port", ":8080", "proxy port")
	uri     = flag.String("uri", "wsproxy", "uri to proxy")
)

func main() {
	flag.Parse()
	wp, err := websocketproxy.NewProxy(*backend, func(r *http.Request) error {
		return nil
	})
	if err != nil {
		panic(err)
	}
	// proxy path
	http.HandleFunc("/"+*uri, wp.Proxy)
	err = http.ListenAndServe(*port, nil)
	if err != nil {
		panic(err)
	}
}
