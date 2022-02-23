package apprunner

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"

	"github.com/briandowns/spinner"
)

func (ar *AppRunner) IsRunning() (bool, error) {
	in := apprunner.DescribeServiceInput{
		ServiceArn: &ar.arn,
	}
	res, err := ar.client.DescribeService(context.Background(), &in)
	if err != nil {
		return false, err
	}
	return res.Service.Status == types.ServiceStatusRunning, nil
}

func (ar *AppRunner) waitForCreated() {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	s.Start()
	// TODO: set timeout
	for {
		created, err := ar.IsRunning()
		if err != nil {
			panic(err)
		}
		if created {
			break
		}
		time.Sleep(time.Second * 5)
	}
	s.Stop()
}

// https://docs.aws.amazon.com/apprunner/latest/api/API_CreateService.html
func (ar *AppRunner) CreateService() error {
	service := ar.Service

	healthCheckConfiguration := types.HealthCheckConfiguration{}

	var instanceConfiguration types.InstanceConfiguration
	serviceAccountName := service.Spec.Template.Spec.ServiceAccountName
	if serviceAccountName != "" {
		instanceConfiguration.InstanceRoleArn = &service.Spec.Template.Spec.ServiceAccountName
	}

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

	res, err := ar.client.CreateService(context.Background(), &params)
	if err != nil {
		return err
	}
	ar.arn = *res.Service.ServiceArn
	fmt.Println("Creating service...")
	ar.waitForCreated()
	fmt.Println("Created.")
	return nil
}
