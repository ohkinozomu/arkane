package apprunner

import (
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
)

func (ar *AppRunner) newSourceConfiguration() types.SourceConfiguration {
	service := ar.svc
	authenticationConfiguration := types.AuthenticationConfiguration{}

	port := strconv.FormatInt(int64(service.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort), 10)
	imageConfiguration := types.ImageConfiguration{
		Port: &port,
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
