package bll

import (
	"context"

	"auth/errors"
	"auth/model"
	"auth/store/postgres"
	"auth/utils"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type auth struct {
	e *casbin.Enforcer
}

var Auth = &auth{}

func init() {
	Register(Auth)
}

func (s *auth) init() func() {
	var (
		a   *gormadapter.Adapter
		err error
	)
	if a, err = postgres.Policy.Adapter(); err != nil {
		utils.Throw(err)
	}

	if s.e, err = casbin.NewEnforcer("./script/model.conf", a); err != nil {
		utils.Throw(err)
	}
	utils.Throw(s.e.LoadPolicy())
	return func() {}
}

// Authentication 认证
func (s *auth) Authentication(ctx context.Context, in model.Authentication) error {

	return nil
}

// Authorization 鉴权
func (s *auth) Authorization(ctx context.Context, in *model.Authorization) error {
	var (
		ok  bool
		err error
	)
	if ok, err = s.e.Enforce(in.Sub, in.Obj, in.Act); err != nil {
		return err
	}
	if ok {
		return nil
	}
	return errors.AuthorizationFailed.New("")
}
