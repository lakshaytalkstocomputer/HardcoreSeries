package main

import (
	"time"
	"github.com/go-kit/kit/log"
)

type loggingMiddleWare struct {
	logger log.Logger
	next StringService
}

func (mw loggingMiddleWare) Uppercase(s string) (output string, err error){
	defer func(begin time.Time){
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw loggingMiddleWare) Count(s string)(n int){
	defer func(begin time.Time){
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n",n,
			"took",time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Count(s)
	return
}
