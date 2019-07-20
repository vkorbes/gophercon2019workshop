package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/ellenkorbes/gophercon2019workshop/francesc/gokit/stringsvc/handler"
	"github.com/ellenkorbes/gophercon2019workshop/francesc/gokit/stringsvc/service"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := flag.Int("p", 8080, "port on which the server should listen")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)

	svc := service.New()
	svc = service.WithLogger(svc, logger)
	svc = service.WithMetrics(svc)

	http.Handle("/uppercase", handler.Uppercase(svc))
	http.Handle("/count", handler.Count(svc))
	http.Handle("/metrics", promhttp.Handler())

	address := fmt.Sprintf(":%d", *port)
	logger.Log("msg", "HTTP", "addr", address)
	logger.Log("err", http.ListenAndServe(address, nil))
}
