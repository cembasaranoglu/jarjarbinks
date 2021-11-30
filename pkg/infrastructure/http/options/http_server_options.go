package options

import "time"

type HttpServerOption struct {
	Name    string
	Metric  *HttpServerMetricOption
}

type HttpServerMetricOption struct{
	ExcludedEndpoints []string
}

type HttpServerStartOption struct {
	Port                            int
	GracefullyShutdown              bool
	GracefullyShutdownTimeoutPeriod time.Duration
}
