package main

import (
	"fmt"

	gotoConfig "github.com/moorara/goto/config"
	"github.com/moorara/microservices-demo/services/switch-service/cmd/config"
	"github.com/moorara/microservices-demo/services/switch-service/cmd/server"
	"github.com/moorara/microservices-demo/services/switch-service/cmd/version"
	"github.com/moorara/microservices-demo/services/switch-service/internal/metrics"
	"github.com/moorara/microservices-demo/services/switch-service/pkg/log"
	"github.com/moorara/microservices-demo/services/switch-service/pkg/trace"
)

func main() {
	config := config.New()
	gotoConfig.Pick(&config)

	logger := log.NewLogger(config.ServiceName, "singleton", config.LogLevel)
	metrics := metrics.New(config.ServiceName)

	sampler := trace.NewConstSampler()
	reporter := trace.NewReporter(config.JaegerLogSpans, config.JaegerAgentAddr)
	tracer, tracerCloser := trace.NewTracer(config.ServiceName, sampler, reporter, logger.Logger, metrics.Registry)
	defer tracerCloser.Close()

	server, err := server.New(config, logger, metrics, tracer)
	if err != nil {
		panic(err)
	}

	logger.Info(
		"version", version.Version,
		"revision", version.Revision,
		"branch", version.Branch,
		"goVersion", version.GoVersion,
		"buildTool", version.BuildTool,
		"buildTime", version.BuildTime,
		"message", fmt.Sprintf("%s started.", config.ServiceName),
	)
 
	server.Start()
}
