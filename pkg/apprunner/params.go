package apprunner

import (
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	v1 "k8s.io/api/core/v1"
)

func k8sEnvToAppRunnerEnv(env []v1.EnvVar) map[string]string {
	arEnv := map[string]string{}
	for _, e := range env {
		arEnv[e.Name] = e.Value
	}
	return arEnv
}

func (ar *AppRunner) newSourceConfiguration() types.SourceConfiguration {
	service := ar.Service
	authenticationConfiguration := types.AuthenticationConfiguration{}

	port := strconv.FormatInt(int64(service.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort), 10)
	imageConfiguration := types.ImageConfiguration{
		Port:                        &port,
		RuntimeEnvironmentVariables: k8sEnvToAppRunnerEnv(service.Spec.Template.Spec.Containers[0].Env),
	}

	imageRepository := types.ImageRepository{
		ImageIdentifier:     &service.Spec.Template.Spec.Containers[0].Image,
		ImageRepositoryType: "ECR_PUBLIC",
		ImageConfiguration:  &imageConfiguration,
	}

	autoDeploymentsEnabled := false
	sourceConfiguration := types.SourceConfiguration{
		AuthenticationConfiguration: &authenticationConfiguration,
		AutoDeploymentsEnabled:      &autoDeploymentsEnabled,
		ImageRepository:             &imageRepository,
	}
	return sourceConfiguration
}
