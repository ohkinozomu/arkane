# arkane

AWS App Runner with Knative API

# Run

## Apply

```
$ arkn apply -f test/sample.yaml
```

## Delete

```
$ arkn delete -f test/sample.yaml
```

# Required IAM policy

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

https://github.com/knative/specs/blob/main/specs/serving/knative-api-specification-1.0.md#spec

- App Runner doesn't support distributed traffic, so `traffic` field is ignored.
https://github.com/aws/apprunner-roadmap/issues/40

- App Runner doesn't support side-cars (AKA multiple containers).
https://github.com/aws/apprunner-roadmap/issues/71

- `namespace`, `volumes`, `timeoutSeconds`, `containerConcurrency` field is ignored.