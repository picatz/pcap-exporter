package server

import (
	"log"
	"net"
	"net/http"

	"github.com/picatz/pcap-exporter/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start() {
	log.Println("starting server")

	// TODO: implement server options
	go metrics.Collect("", "")

	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)

	ln, err := net.Listen("tcp", "127.0.0.1:9249")
	if err != nil {
		log.Fatal(err)
	}

	http.Serve(ln, mux)
}
