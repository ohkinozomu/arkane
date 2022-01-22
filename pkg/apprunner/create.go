package apprunner

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
)

// https://docs.aws.amazon.com/apprunner/latest/api/API_CreateService.html
func (ar *AppRunner) CreateService() error {
	printWarn()
	fmt.Println("Creating service...")

	service := ar.Service

	healthCheckConfiguration := types.HealthCheckConfiguration{}
	instanceConfiguration := types.InstanceConfiguration{}

	sourceConfiguration := ar.newSourceConfiguration()

	params := apprunner.CreateServiceInput{
		HealthCheckConfiguration: &healthCheckConfiguration,
		InstanceConfiguration:    &instanceConfiguration,

		ServiceName: &service.ObjectMeta.Name,

		SourceConfiguration: &sourceConfiguration,

		// AutoScalingConfigurationArn:
		// EncryptionConfiguration:
		// Tags:
	}

	_, err := ar.client.CreateService(context.Background(), &params)
	if err != nil {
		return err
	}
	fmt.Println("Created.")
	return nil
}
