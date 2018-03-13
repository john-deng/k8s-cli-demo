package main

import (
	_ "github.com/john-deng/k8s-cli-demo/routers"
	"github.com/astaxie/beego"
)

func init() {

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
