package controllers

import (
	"github.com/john-deng/k8s-cli-demo/models"
	"github.com/kataras/iris"

)

// Operations about object
type DeploymentController struct {

}

func (c *DeploymentController) Before(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())

	// .Next is required to move forward to the chain of handlers,
	// if missing then it stops the execution at this handler.
	ctx.Next()
}


// @Title Deploy
// @Description deploy application
// @Param	body
// @Success 200 {string}
// @Failure 403 body is empty
// @router / [post]
func (c *DeploymentController) Deploy(ctx iris.Context) {
	var app models.App
	err := ctx.ReadJSON(&app)
	if err != nil {
		ctx.Values().Set("error", "deployment failed, read and parse request body failed. " + err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	// invoke models
	models.Deploy(&app)

	ctx.JSON(app)
}
