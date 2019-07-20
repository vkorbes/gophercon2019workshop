package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/ellenkorbes/gophercon2019workshop/francesc/gokit/stringsvc/service"
)

// Uppercase returns a JSON on HTTP handler for Uppercase.
func Uppercase(svc service.StringService) http.Handler {
	return httptransport.NewServer(
		uppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)
}

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

func uppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{Err: err.Error()}, nil
		}
		return uppercaseResponse{V: v}, nil
	}
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
