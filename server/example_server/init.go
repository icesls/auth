package example_server

import (
	"go-micro.dev/v4"
	"auth/lib/example_server"
)

func Init() micro.Option {
	return micro.Server(example_server.NewServer())
}
