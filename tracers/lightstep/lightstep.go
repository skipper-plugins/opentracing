package main

import (
	"errors"
	"strings"

	lightstep "github.com/lightstep/lightstep-tracer-go"
	opentracing "github.com/opentracing/opentracing-go"
)

func InitTracer(opts []string) (opentracing.Tracer, error) {
	var token string
	for _, o := range opts {
		if strings.HasPrefix(o, "token=") {
			token = o[6:]
		}
	}
	if token == "" {
		return nil, errors.New("missing token= option")
	}
	return lightstep.NewTracer(lightstep.Options{
		AccessToken: token,
		Tags: map[string]interface{}{
			lightstep.ComponentNameKey: "skipper",
		},
	}), nil
}
