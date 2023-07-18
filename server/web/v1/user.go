package v1

import (
	"auth/bll"
	"auth/model"
	"auth/server/web/middleware"
	"auth/utils"
	"github.com/gin-gonic/gin"
)

var User = &user{}

func init() {
	RegisterRouter(User)
}

type user struct{}

// Init
func (a *user) Init(r *gin.RouterGroup) {
	g := r.Group("/user", middleware.Auth())
	{
		g.POST("/create", a.create)
		g.POST("/update", a.update)
		g.POST("/list", a.list)
		g.POST("/delete", a.delete)
		g.POST("/detail", a.find)
	}
}

// create
func (a *user) create(c *gin.Context) {
	var (
		in  = &model.UserCreateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.User.Create(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// update
func (a *user) update(c *gin.Context) {
	var (
		in  = &model.UserUpdateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.User.Update(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// list
func (a *user) list(c *gin.Context) {
	var (
		in  = &model.UserListRequest{}
		out = &model.UserListResponse{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.User.List(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// list
func (a *user) find(c *gin.Context) {
	var (
		in  = &model.UserInfoRequest{}
		out = &model.UserInfo{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.User.Find(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// delete
func (a *user) delete(c *gin.Context) {
	var (
		in  = &model.UserDeleteRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.User.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
