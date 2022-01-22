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
apprunner:ListServices
apprunner:UpdateService
iam:CreateServiceLinkedRole
```

## Delete

```
apprunner:ListServices
apprunner:DeleteService
```

# Supported API

only `serving.knative.dev/v1`

# Projecvt Status

PoC

# Difference between Knative API and App Runner

App Runner doesn't support distributed traffic, so `traffic` field is ignored.
https://github.com/aws/apprunner-roadmap/issues/40
https://github.com/knative/specs/blob/main/specs/serving/knative-api-specification-1.0.md#spec