package v1

import (
	"auth/bll"
	"auth/model"
	"auth/server/web/middleware"
	"auth/utils"
	"github.com/gin-gonic/gin"
)

var Resource = &resource{}

func init() {
	RegisterRouter(Resource)
}

type resource struct{}

// Init
func (a *resource) Init(r *gin.RouterGroup) {
	g := r.Group("/resource", middleware.Auth())
	{
		g.POST("/create", a.create)
		g.POST("/update", a.update)
		g.POST("/list", a.list)
		g.POST("/delete", a.delete)
		g.POST("/detail", a.find)
	}
}

// create
func (a *resource) create(c *gin.Context) {
	var (
		in  = &model.ResourceCreateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Resource.Create(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// update
func (a *resource) update(c *gin.Context) {
	var (
		in  = &model.ResourceUpdateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Resource.Update(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// list
func (a *resource) list(c *gin.Context) {
	var (
		in  = &model.ResourceListRequest{}
		out = &model.ResourceListResponse{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Resource.List(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// list
func (a *resource) find(c *gin.Context) {
	var (
		in  = &model.ResourceInfoRequest{}
		out = &model.ResourceInfo{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Resource.Find(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// delete
func (a *resource) delete(c *gin.Context) {
	var (
		in  = &model.ResourceDeleteRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Resource.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
