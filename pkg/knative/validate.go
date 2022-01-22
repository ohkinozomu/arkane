package knative

import (
	"context"

	api "knative.dev/serving/pkg/apis/serving/v1"
)

func Validate(svc api.Service) error {
	errs := svc.Validate(context.Background())
	if errs != nil {
		return errs
	}
	return nil
}
