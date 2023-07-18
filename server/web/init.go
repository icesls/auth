package web

import (
	"auth/server/web/middleware"
	"auth/server/web/v1"
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/asim/go-micro/plugins/server/http/v4"

	"time"

	"auth/config"
	"auth/utils"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

var (
	webName           string
	webAddr           string
	defaultListenAddr = ":7771"
	srv               server.Server
)

func Init() micro.Option {
	webName = config.Conf.Name + "-http"

	if config.Conf.Http.Listen != "" {
		defaultListenAddr = config.Conf.Http.Listen
	}

	srv = http.NewServer(
		server.Name(webName),
		server.Version(config.Conf.Version),
		server.Address(defaultListenAddr),
		server.RegisterTTL(time.Second*30),
		server.RegisterInterval(time.Second*15),
		server.Registry(etcd.NewRegistry(registry.Addrs(config.EtcdAddr))),
	)

	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// resource routers
	v1Group := router.Group("v1", middleware.OtelMiddleware())
	v1.Init(v1Group)

	utils.Throw(srv.Handle(srv.NewHandler(router)))
	utils.Throw(srv.Start())
	if config.Conf.Traefik.Enabled {
		utils.Throw(config.Register(webName, srv))
	}
	return micro.Server(srv)
}
