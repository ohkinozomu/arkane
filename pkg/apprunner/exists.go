package apprunner

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

func (ar *AppRunner) ServiceExists() (bool, error) {
	name := ar.svc.ObjectMeta.Name

	fmt.Printf("%v already exists\n", name)
	params := apprunner.ListServicesInput{}
	out, err := ar.client.ListServices(context.Background(), &params)
	if err != nil {
		return false, err
	}
	for _, s := range out.ServiceSummaryList {
		if *s.ServiceName == ar.svc.ObjectMeta.Name {
			ar.arn = *s.ServiceArn
			return true, nil
		}
	}
	return false, nil
}
