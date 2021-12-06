package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gclamigueiro/gokit-example-string-service/internal/entity"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpHandler(uppercaseEndpoint endpoint.Endpoint, countEndpoint endpoint.Endpoint) {

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		countEndpoint,
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request entity.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request entity.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
