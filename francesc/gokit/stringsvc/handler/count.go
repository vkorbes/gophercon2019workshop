package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/ellenkorbes/gophercon2019workshop/francesc/gokit/stringsvc/service"
)

// Count returns a JSON on HTTP handler for Count.
func Count(svc service.StringService) http.Handler {
	return httptransport.NewServer(
		countEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func countEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{V: v}, nil
	}
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req countRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
