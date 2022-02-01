# arkane

AWS App Runner with Knative API

# Run

## Apply

```
$ go run main.go apply -f test/sample.yaml
```

## Delete

```
$ go run main.go delete -f test/sample.yaml
```

# Required IAM policy

## Apply

```
apprunner:CreateService
apprunner:UpdateService
apprunner:ListServices
apprunner:DeleteService
apprunner:DescribeService
iam:CreateServiceLinkedRole
```

# Required Go version

`^1.16.0`

https://github.com/knative/serving/pull/12051

# Supported API

only `serving.knative.dev/v1`

# Project Status

PoC

# Difference between Knative API and App Runner

- App Runner doesn't support distributed traffic, so `traffic` field is ignored.
https://github.com/aws/apprunner-roadmap/issues/40
https://github.com/knative/specs/blob/main/specs/serving/knative-api-specification-1.0.md#spec

- `namespace` field is ignored.