package main

import (
	"os"
	"os/signal"

	"auth/assistant"
	"auth/bll"
	"auth/client"
	"auth/collector"
	log "auth/collector/logger"
	"auth/config"
	"auth/event"
	"auth/server"
	"auth/store"
	sig "go-micro.dev/v4/util/signal"
)

//go:generate phanes gen -register
//go:generate swag init

type InitFunc func() func()

var bootstraps = []InitFunc{}

func main() {
	var (
		cancels = make([]func(), 0)

		// system init func
		bootstraps = []InitFunc{
			config.Init,
			collector.Init,
			event.Init,
			server.Init,
			client.Init,
			store.Init,
			assistant.Init,
			bll.Init,
		}
	)

	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, sig.Shutdown()...)
		<-sigint

		for _, cancel := range cancels {
			cancel()
		}
		close(config.ExitC)
	}()

	for _, fn := range bootstraps {
		cancels = append(cancels, fn())
	}
	log.Info("finished to init all component")

	<-config.ExitC
	log.Info("server shutdown!")
}
