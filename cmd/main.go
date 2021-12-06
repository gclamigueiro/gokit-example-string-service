package main

import (
	"net/http"
	"os"

	"github.com/go-kit/log"

	"github.com/gclamigueiro/gokit-example-string-service/internal/endpoint"
	"github.com/gclamigueiro/gokit-example-string-service/internal/handler"
	"github.com/gclamigueiro/gokit-example-string-service/internal/service"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	svc := service.NewStringService()
	svc = service.NewLoggingMiddleware(logger, svc)
	svc = service.NewInstrumentingMiddleware(requestCount, requestLatency, countResult, svc)

	uppercaseEndpoint := endpoint.MakeUppercaseEndpoint(svc)
	uppercaseEndpoint = endpoint.LoggingMiddleware(log.With(logger, "method", "uppercase"))(uppercaseEndpoint)

	countEndpoint := endpoint.MakeCountEndpoint(svc)
	countEndpoint = endpoint.LoggingMiddleware(log.With(logger, "method", "count"))(countEndpoint)

	handler.NewHttpHandler(uppercaseEndpoint, countEndpoint)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)

}
