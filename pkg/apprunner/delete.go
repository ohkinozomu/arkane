package apprunner

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/briandowns/spinner"
)

func (ar *AppRunner) waitForDeleted() {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	s.Start()
	// TODO: set timeout
	for {
		exists, err := ar.ServiceExists()
		if err != nil {
			panic(err)
		}
		if !exists {
			break
		}
		time.Sleep(time.Second * 5)
	}
	s.Stop()
}

// https://docs.aws.amazon.com/apprunner/latest/api/API_DeleteService.html
func (ar *AppRunner) DeleteService() error {
	printWarn()
	fmt.Println("Deleting service...")

	params := apprunner.DeleteServiceInput{
		ServiceArn: &ar.arn,
	}

	_, err := ar.client.DeleteService(context.Background(), &params)
	if err != nil {
		return err
	}
	ar.waitForDeleted()
	fmt.Println("Deleted.")
	return nil
}
