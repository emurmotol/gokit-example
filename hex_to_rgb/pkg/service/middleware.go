package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(HexToRgbService) HexToRgbService

type loggingMiddleware struct {
	logger log.Logger
	next   HexToRgbService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a HexToRgbService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next HexToRgbService) HexToRgbService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Transform(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Transform", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Transform(ctx, s)
}
