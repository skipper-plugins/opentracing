package main

import (
	instana "github.com/instana/golang-sensor"
	opentracing "github.com/opentracing/opentracing-go"
)

func InitTracer(_ []string) (opentracing.Tracer, error) {
	return instana.NewTracerWithOptions(&instana.Options{
		Service:  "skipper",
		LogLevel: instana.Error,
	}), nil
}
