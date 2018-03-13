package models

import (
	log "github.com/astaxie/beego/logs"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"fmt"
	"k8s.io/apimachinery/pkg/util/intstr"
)


type App struct{
	Name string	`json:"name"`
	Project string `json:"project"`
	DockerRegistry string `json:"docker_registry"`
	ImageTag string `json:"image_tag"`
	Port int32 `json:"port"`
	Ports []int32 `json:"ports"`
}

type Response struct {
	Message string
	Code int
	Data interface{}
}

func init() {

}


func int32Ptr(i int32) *int32 { return &i }


func Deploy(app *App) (string, error)  {
	deploymentsClient := KubeApi.clientSet.AppsV1beta1().Deployments(app.Project)

	if "" == app.ImageTag {
		app.ImageTag = "latest"
	}
	if "" == app.DockerRegistry {
		app.DockerRegistry = "docker-registry.default.svc:5000"
	}

	deployment := &appsv1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: app.Name,
		},
		Spec: appsv1beta1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": app.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  app.Name,
							Image: app.DockerRegistry + "/" + app.Project + "/" + app.Name + ":" + app.ImageTag,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
							Env: []apiv1.EnvVar{
								{
									Name: "KUBERNETES_SERVICE_HOST",
									Value: "devops.oc.com",
								},
								{
									Name: "KUBERNETES_SERVICE_PORT",
									Value: "8443",
								},
							},
						},
					},
				},
			},
		},
	}

	log.Debug(deployment)

	// Create Deployment
	log.Info("Creating deployment...")
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}

	retVal := fmt.Sprintf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	log.Info(retVal)

	createService(app)

	return retVal, err
}

func createService(app *App) (*apiv1.Service, error)   {
	// create service

	serviceClient := KubeApi.clientSet.CoreV1().Services(app.Project)
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: app.Name,
			Labels: map[string]string{
				"app": app.Name,
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Name:       "http",
					Protocol:   apiv1.ProtocolTCP,
					Port:       8080,
					TargetPort: intstr.IntOrString{IntVal: 8080},
				},
			},
		},
	}

	var svc *apiv1.Service
	var err error
	svc, err = serviceClient.Create(service)
	log.Info(svc, err)

	return svc, err
}