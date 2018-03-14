package main

import (
	"github.com/kataras/iris"
	"github.com/john-deng/k8s-cli-demo/routers"
)

func main() {
	// Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
	routers.App.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)
}
