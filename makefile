

group = example_group_name
project = project_name
version = latest
server = your.docker.registry

gen:
	go generate

run:gen
	go run main.go --registry=etcd --registry_address=127.0.0.1:2389 --config=./script/config.yaml

config:
	etcdctl put "/hello/config/auth"  < ./script/config.json

default:
	@echo ${group}/${project}

build:gen
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o ./bin/${project}

image:build
	docker build -t ${group}/${project}:${version} --platform linux/amd64 --build-arg ARG_PROJECT_NAME=${project} .

push:image
	docker tag ${group}/${project}:${version} ${server}/${group}/${project}:${version}
	docker push ${server}/${group}/${project}:${version}
