# arkane

AWS App Runner with Knative API

# Run

```
$ go run main.go apply --file test/sample.yaml
```

# IAM

## apply

```
apprunner:CreateService
apprunner:ListServices
apprunner:UpdateService
iam:CreateServiceLinkedRole
```

# Support API

only `serving.knative.dev/v1`

# Status

PoC