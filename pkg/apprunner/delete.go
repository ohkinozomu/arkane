package apprunner

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

// https://docs.aws.amazon.com/apprunner/latest/api/API_DeleteService.html
func (ar *AppRunner) DeleteService() error {
	fmt.Println("Deleting service...")

	params := apprunner.DeleteServiceInput{
		ServiceArn: &ar.arn,
	}

	_, err := ar.client.DeleteService(context.Background(), &params)
	if err != nil {
		return err
	}
	fmt.Println("Deleted.")
	return nil
}
