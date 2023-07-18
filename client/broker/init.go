package broker

import (
	"go-micro.dev/v4/broker"
	log "auth/collector/logger"
	"auth/config"
)

var PubSub broker.Broker

func Init() func() {
	switch config.Conf.Broker.Type {
	case "rabbitmq":
		PubSub = InitRabbit()
	case "nats":
		PubSub = InitNats()
	}

	return func() {
		if err := PubSub.Disconnect(); err != nil {
			log.Error(err.Error())
		}
	}
}
