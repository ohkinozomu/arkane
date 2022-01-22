package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/ohkinozomu/arkane/pkg/knative"
	api "knative.dev/serving/pkg/apis/serving/v1"
)

type AppRunner struct {
	client  *apprunner.Client
	Service api.Service
	arn     string
}

func newClient() (*apprunner.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	client := apprunner.NewFromConfig(cfg)
	return client, nil
}

func New(svc api.Service) (AppRunner, error) {
	var ar AppRunner
	client, err := newClient()
	if err != nil {
		return ar, err
	}

	err = knative.Validate(svc)
	if err != nil {
		return ar, err
	}

	ar.client = client
	ar.Service = svc
	return ar, nil
}
