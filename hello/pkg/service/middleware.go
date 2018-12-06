package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(HelloService) HelloService

type loggingMiddleware struct {
	logger log.Logger
	next   HelloService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next HelloService) HelloService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SayHello(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "SayHello", "s", s, "rs", rs, "err", err)
	}()
	return l.next.SayHello(ctx, s)
}
