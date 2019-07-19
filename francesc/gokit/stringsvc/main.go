package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	port := flag.Int("p", 8080, "port on which the server should listen")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)

	svc := stringService{}

	var uppercase endpoint.Endpoint
	uppercase = makeUppercaseEndpoint(svc)
	uppercase = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)
	http.Handle("/uppercase", httptransport.NewServer(
		uppercase,
		decodeUppercaseRequest,
		encodeResponse,
	))

	var count endpoint.Endpoint
	count = makeCountEndpoint(svc)
	count = loggingMiddleware(log.With(logger, "method", "count"))(count)
	http.Handle("/count", httptransport.NewServer(
		count,
		decodeCountRequest,
		encodeResponse,
	))

	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
