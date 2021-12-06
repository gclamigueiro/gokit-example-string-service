package endpoint

import (
	"context"

	"github.com/gclamigueiro/gokit-example-string-service/internal/entity"
	"github.com/gclamigueiro/gokit-example-string-service/internal/service"
	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return entity.UppercaseResponse{V: v, Err: err.Error()}, nil
		}
		return entity.UppercaseResponse{V: v, Err: ""}, nil
	}
}

func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.CountRequest)
		v := svc.Count(req.S)
		return entity.CountResponse{V: v}, nil
	}
}
