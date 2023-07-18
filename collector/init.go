package collector

import (
	"auth/collector/logger"
	"auth/collector/metrics"
	"auth/collector/trace"
)

func Init() func() {
	var (
		cancels = make([]func(), 0)

		inits = []func() func(){
			trace.Init,
			logger.Init,
			metrics.Init,
		}
	)

	for _, init := range inits {
		cancels = append(cancels, init())
	}

	return func() {
		for _, cancel := range cancels {
			cancel()
		}
	}
}
