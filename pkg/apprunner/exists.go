package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

func (ar *AppRunner) ServiceExists() (bool, error) {
	params := apprunner.ListServicesInput{}
	out, err := ar.client.ListServices(context.Background(), &params)
	if err != nil {
		return false, err
	}
	for _, s := range out.ServiceSummaryList {
		if *s.ServiceName == ar.Service.ObjectMeta.Name {
			ar.arn = *s.ServiceArn
			return true, nil
		}
	}
	return false, nil
}
