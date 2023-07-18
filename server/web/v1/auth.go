package v1

import (
	"auth/bll"
	"auth/model"
	"auth/server/web/middleware"
	"auth/utils"

	"github.com/gin-gonic/gin"
)

var Auth = &auth{}

func init() {
	RegisterRouter(Auth)
}

type auth struct{}

// Init
func (s *auth) Init(r *gin.RouterGroup) {
	g := r.Group("/auth", middleware.Auth())
	{
		g.POST("/authorization", s.authorization)

	}
}

func (s *auth) authorization(c *gin.Context) {
	var (
		in  = &model.Authorization{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Auth.Authorization(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
