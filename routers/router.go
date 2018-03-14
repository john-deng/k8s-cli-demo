// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/kataras/iris"
	"github.com/john-deng/k8s-cli-demo/controllers"
)

var (
	App *iris.Application
	deploymentController controllers.DeploymentController
)

func init() {
	App = iris.New()

	deploymentController = controllers.DeploymentController {}

	deploymentRoutes := App.Party("/deployment", deploymentController.Before)
	{
		// Method POST: http://localhost:8080/deployment/deploy
		deploymentRoutes.Post("/deploy", deploymentController.Deploy)
	}
}
