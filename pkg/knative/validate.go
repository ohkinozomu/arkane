package knative

import (
	"context"
	"errors"

	wr "github.com/ohkinozomu/which-registry"
	api "knative.dev/serving/pkg/apis/serving/v1"
)

func Validate(svc api.Service) error {
	errs := svc.Validate(context.Background())
	if errs != nil {
		return errs
	}

	image := svc.Spec.Template.Spec.Containers[0].Image
	registry, err := wr.Which(image)
	if err != nil {
		return err
	}

	if registry != wr.ECR_PUBLIC && registry != wr.ECR_PRIVATE {
		msg := "[ERROR] App Runner only support ECR image: https://docs.aws.amazon.com/apprunner/latest/dg/service-source-image.html#service-source-image.providers"
		return errors.New(msg)
	}

	if svc.APIVersion != "serving.knative.dev/v1" {
		msg := "[ERROR] Unsupporetd API version"
		return errors.New(msg)
	}

	return nil
}
