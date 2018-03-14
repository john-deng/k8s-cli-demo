#!/usr/bin/env bash

GOOS=linux go build -o app main.go

docker build --tag=docker-registry-default.apps.oc.com/demo-dev/k8s-cli-demo:latest .

docker push docker-registry-default.apps.oc.com/demo-dev/k8s-cli-demo:latest