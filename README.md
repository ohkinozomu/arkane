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

# IAM

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

# Support API

only `serving.knative.dev/v1`

# Status

PoC