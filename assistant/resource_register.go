package assistant

import (
	"context"
	"time"

	"auth/config"
	"auth/utils"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const RegisterKeyPrefix = "/phanes/register_resource"

func init() {
	Register(Resource)
}

type resource struct{}

var Resource = &resource{}

func (r *resource) Init() func() {
	var (
		err       error
		cli       *clientv3.Client
		endpoints = make([]string, 0)
	)

	if config.EtcdAddr == "" {
		endpoints = []string{"localhost:2379"}
	} else {
		endpoints = append(endpoints, config.EtcdAddr)
	}

	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second * 3,
	})
	utils.Throw(err)
	utils.Throw(r.register(cli))

	return func() {}
}

func (r *resource) register(cli *clientv3.Client) error {
	var err error
	body := []map[string]interface{}{
		{
			"is_parent": true,
			"name":      "group",
			"type":      "method",
			"path":      "/v1/group",
		},
		{
			"is_parent": false,
			"name":      "group.create",
			"type":      "method",
			"path":      "/v1/group/create",
		},
	}

	if _, err = cli.KV.Put(context.Background(), RegisterKeyPrefix+"/group", utils.ToJsonString(body)); err != nil {
		return err
	}

	return nil
}
