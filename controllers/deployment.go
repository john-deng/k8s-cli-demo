package controllers

import (
	"github.com/john-deng/k8s-cli-demo/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type DeploymentController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (c *DeploymentController) Post() {
	var app models.App
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &app); err == nil {
		if _, err := models.Deploy(&app); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = app
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
