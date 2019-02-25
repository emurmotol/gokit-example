// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/marcpar/HextoRGB/hex_to_rgb/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	TransformEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.HexToRgbService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{TransformEndpoint: MakeTransformEndpoint(s)}
	for _, m := range mdw["Transform"] {
		eps.TransformEndpoint = m(eps.TransformEndpoint)
	}
	return eps
}
