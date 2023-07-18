package v1

import (
	"auth/bll"
	"auth/model"
	"auth/server/web/middleware"
	"auth/utils"
	"github.com/gin-gonic/gin"
)

var Role = &role{}

func init() {
	RegisterRouter(Role)
}

type role struct{}

// Init
func (a *role) Init(r *gin.RouterGroup) {
	g := r.Group("/role", middleware.Auth())
	{
		g.POST("/create", a.create)
		g.POST("/update", a.update)
		g.POST("/list", a.list)
		g.POST("/delete", a.delete)
		g.POST("/detail", a.find)
	}
}

// create
func (a *role) create(c *gin.Context) {
	var (
		in  = &model.RoleCreateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Role.Create(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// update
func (a *role) update(c *gin.Context) {
	var (
		in  = &model.RoleUpdateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Role.Update(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// list
func (a *role) list(c *gin.Context) {
	var (
		in  = &model.RoleListRequest{}
		out = &model.RoleListResponse{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Role.List(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// list
func (a *role) find(c *gin.Context) {
	var (
		in  = &model.RoleInfoRequest{}
		out = &model.RoleInfo{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Role.Find(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// delete
func (a *role) delete(c *gin.Context) {
	var (
		in  = &model.RoleDeleteRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Role.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
