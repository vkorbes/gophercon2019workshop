package service

import (
	"time"

	"github.com/go-kit/kit/log"
)

// WithLogger returns a StringService that will log to the given logger.
func WithLogger(svc StringService, logger log.Logger) StringService {
	return loggingMiddleware{svc, logger}
}

type loggingMiddleware struct {
	next   StringService
	logger log.Logger
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.next.Uppercase(s)
}

func (mw loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.next.Count(s)
}
