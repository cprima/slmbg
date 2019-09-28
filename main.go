package main

import (
	"net/http"

	"github.com/cprior/slmbg/cmd"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
}

func main() {

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()
	cmd.Execute()
}
