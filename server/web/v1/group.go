package v1

import (
	"auth/bll"
	"auth/model"
	"auth/server/web/middleware"
	"auth/utils"
	"github.com/gin-gonic/gin"
)

var Group = &group{}

func init() {
	RegisterRouter(Group)
}

type group struct{}

// Init
// #[resource resource]
// #[register_parent("group", "method", "v1/group")]
// #[register_child("create", "method", "v1/group/create")]
// #[resource("create", "method", "v1/group/create")]
func (a *group) Init(r *gin.RouterGroup) {
	g := r.Group("/group", middleware.Auth())
	{
		g.POST("/create", a.create)
		g.POST("/update", a.update)
		g.POST("/list", a.list)
		g.POST("/delete", a.delete)
		g.POST("/detail", a.find)
	}
}

// create
func (a *group) create(c *gin.Context) {
	var (
		in  = &model.GroupCreateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Group.Create(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// update
func (a *group) update(c *gin.Context) {
	var (
		in  = &model.GroupUpdateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Group.Update(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// list
func (a *group) list(c *gin.Context) {
	var (
		in  = &model.GroupListRequest{}
		out = &model.GroupListResponse{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Group.List(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// list
func (a *group) find(c *gin.Context) {
	var (
		in  = &model.GroupInfoRequest{}
		out = &model.GroupInfo{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Group.Find(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// delete
func (a *group) delete(c *gin.Context) {
	var (
		in  = &model.GroupDeleteRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Group.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
