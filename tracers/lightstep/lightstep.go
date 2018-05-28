package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	lightstep "github.com/lightstep/lightstep-tracer-go"
	opentracing "github.com/opentracing/opentracing-go"
)

const (
	defComponentName = "skipper"
)

func InitTracer(opts []string) (opentracing.Tracer, error) {
	componentName := defComponentName
	var token, host string
	var port int

	for _, o := range opts {
		parts := strings.SplitN(o, "=", 2)
		switch parts[0] {
		case "component-name":
			if len(parts) > 1 {
				componentName = parts[1]
			}
		case "token":
			token = o[6:]
		case "collector":
			subparts := strings.Split(parts[1], ":")
			host = subparts[0]
			if len(subparts) == 1 {
				port = 443
			} else {
				var err error
				aport, err := strconv.Atoi(subparts[1])
				if err != nil {
					return nil, fmt.Errorf("failed to parse %s as int: %s", subparts[1], err)
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
			lightstep.ComponentNameKey: componentName,
		},
	}), nil
}
