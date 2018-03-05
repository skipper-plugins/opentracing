package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	lightstep "github.com/lightstep/lightstep-tracer-go"
	opentracing "github.com/opentracing/opentracing-go"
)

func InitTracer(opts []string) (opentracing.Tracer, error) {
	var token, host string
	var port int
	for _, o := range opts {
		switch {
		case strings.HasPrefix(o, "token="):
			token = o[6:]
		case strings.HasPrefix(o, "collector="):
			parts := strings.Split(o[10:], ":")
			host = parts[0]
			if len(parts) == 1 {
				port = 443
			} else {
				var err error
				aport, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("failed to parse %s as int: %s", parts[1], err)
				}
				port = int(aport)
			}
		}

	}
	if token == "" {
		return nil, errors.New("missing token= option")
	}
	if host == "" {
		host = lightstep.DefaultGRPCCollectorHost
		port = 443
	}
	return lightstep.NewTracer(lightstep.Options{
		AccessToken: token,
		Collector: lightstep.Endpoint{
			Host: host,
			Port: port,
		},
		UseGRPC: true,
		Tags: map[string]interface{}{
			lightstep.ComponentNameKey: "skipper",
		},
	}), nil
}
