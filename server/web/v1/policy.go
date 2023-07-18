package v1

import (
	"auth/bll"
	"auth/model"
	"auth/server/web/middleware"
	"auth/utils"
	"github.com/gin-gonic/gin"
)

var Policy = &policy{}

func init() {
	RegisterRouter(Policy)
}

type policy struct{}

// Init
func (a *policy) Init(r *gin.RouterGroup) {
	g := r.Group("/policy", middleware.Auth())
	{
		g.POST("/create", a.create)
		g.POST("/update", a.update)
		g.POST("/list", a.list)
		g.POST("/delete", a.delete)
		g.POST("/detail", a.find)
	}
}

// create
func (a *policy) create(c *gin.Context) {
	var (
		in  = &model.PolicyCreateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Policy.Create(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// update
func (a *policy) update(c *gin.Context) {
	var (
		in  = &model.PolicyUpdateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Policy.Update(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// list
func (a *policy) list(c *gin.Context) {
	var (
		in  = &model.PolicyListRequest{}
		out = &model.PolicyListResponse{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Policy.List(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// list
func (a *policy) find(c *gin.Context) {
	var (
		in  = &model.PolicyInfoRequest{}
		out = &model.PolicyInfo{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Policy.Find(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// delete
func (a *policy) delete(c *gin.Context) {
	var (
		in  = &model.PolicyDeleteRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Policy.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
