package apprunner

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

// https://docs.aws.amazon.com/apprunner/latest/api/API_UpdateService.html
func (ar *AppRunner) UpdateService() error {
	printWarn()
	fmt.Println("Updating service...")

	sourceConfiguration := ar.newSourceConfiguration()

	params := apprunner.UpdateServiceInput{
		ServiceArn:          &ar.arn,
		SourceConfiguration: &sourceConfiguration,
	}

	_, err := ar.client.UpdateService(context.Background(), &params)
	if err != nil {
		return err
	}
	fmt.Println("Updated.")
	return nil
}
