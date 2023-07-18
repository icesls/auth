package bll

import (
	"context"
	"encoding/json"
	"time"

	log "auth/collector/logger"
	"auth/config"
	"auth/model"
	"auth/model/entity"
	"auth/model/mapping"
	"auth/store"
	"auth/store/postgres"
	"auth/utils"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

const WatchResourcePrefix = "/phanes/register_resource"

type resource struct {
	iResource store.IResource
	cli       *clientv3.Client
}

var Resource = &resource{
	iResource: postgres.Resource,
}

func init() {
	Register(Resource)
}

func (a *resource) init() func() {
	var (
		err       error
		endpoints = make([]string, 0)
	)

	if config.EtcdAddr == "" {
		endpoints = []string{"localhost:2379"}
	} else {
		endpoints = append(endpoints, config.EtcdAddr)
	}

	a.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second * 3,
	})
	utils.Throw(err)

	if err = a.registerFromEtcd(); err != nil {
		log.Error("register resource error", zap.String("err", err.Error()))
	}
	go a.watch()

	return func() {
		a.cli.Close()
	}
}

// Create
func (a *resource) Create(ctx context.Context, in *model.ResourceCreateRequest) error {
	var (
		err error
	)
	c := buildResource(in)
	_, err = a.iResource.Create(ctx, c)
	return err
}

// Update
func (a *resource) Update(ctx context.Context, in *model.ResourceUpdateRequest) error {
	var (
		dict = make(map[string]interface{})
	)

	if in.Icon != nil {
		dict["icon"] = in.Icon
	}

	if in.ParentId != nil {
		dict["parent_id"] = in.ParentId
	}

	if in.CreatedAt != nil {
		dict["created_at"] = in.CreatedAt
	}

	if in.UpdatedAt != nil {
		dict["updated_at"] = in.UpdatedAt
	}

	// do other update here
	updateAt := time.Now().Unix()
	in.UpdatedAt = &updateAt
	return a.iResource.Update(ctx, in.Id, dict)
}

// Delete
func (a *resource) Delete(ctx context.Context, in *model.ResourceDeleteRequest) error {
	return a.iResource.Delete(ctx, in.Id)
}

// List
func (a *resource) List(ctx context.Context, in *model.ResourceListRequest) (*model.ResourceListResponse, error) {
	var (
		err   error
		total int
		list  []*entity.Resource
		out   = &model.ResourceListResponse{}
	)

	if total, list, err = a.iResource.List(ctx, in); err != nil {
		return nil, err
	}

	out.Total = total
	out.List = mapping.ResourcesEntityToDto(list)

	return out, nil
}

// Find
func (a *resource) Find(ctx context.Context, in *model.ResourceInfoRequest) (*model.ResourceInfo, error) {
	var (
		err  error
		data *entity.Resource
		out  = &model.ResourceInfo{}
	)

	if data, err = a.iResource.Find(ctx, in); err != nil {
		return nil, err
	}

	out = mapping.ResourceEntityToDto(data)
	return out, nil
}

func (a *resource) registerFromEtcd() error {
	var (
		err      error
		ctx      = context.Background()
		datas    []*model.ResourceCreateRequest
		response *clientv3.GetResponse
	)

	if response, err = a.cli.KV.Get(ctx, WatchResourcePrefix, clientv3.WithPrefix()); err != nil {
		return err
	}

	for _, kv := range response.Kvs {
		if err = json.Unmarshal(kv.Value, &datas); err != nil {
			return err
		}
	}

	return a.register(datas...)
}

func (a *resource) watch() {
	var (
		err   error
		datas = make([]*model.ResourceCreateRequest, 0)
	)
	watchChan := a.cli.Watcher.Watch(context.Background(), WatchResourcePrefix, clientv3.WithPrefix())
	for resp := range watchChan {
		for _, event := range resp.Events {
			if err = json.Unmarshal(event.Kv.Value, &datas); err != nil {
				log.Error("watch resource error", zap.String("err", err.Error()))
				return
			}
		}
		a.register(datas...)
	}
}

func (a *resource) register(resources ...*model.ResourceCreateRequest) error {
	var (
		err      error
		ctx      = context.Background()
		parentId int64
		entities = make([]*entity.Resource, 0)
	)

	for _, d := range resources {
		r := buildResource(d)

		if d.IsParent {
			parentId, err = a.iResource.Create(ctx, r)
			if err != nil {
				return err
			}
		} else {
			entities = append(entities, r)
		}
	}

	if parentId != 0 {
		for _, r := range entities {
			r.ParentId = parentId
		}
	}
	return a.iResource.CreateMany(ctx, entities...)
}

// buildResource build entity
func buildResource(in *model.ResourceCreateRequest) *entity.Resource {
	now := time.Now()
	ety := &entity.Resource{

		Name: in.Name,

		Type: in.Type,

		Path: in.Path,

		CreatedAt: now,

		UpdatedAt: now,
	}

	if in.Icon != nil {
		ety.Icon = *in.Icon
	}

	if in.ParentId != nil {
		ety.ParentId = *in.ParentId
	}

	return ety
}
