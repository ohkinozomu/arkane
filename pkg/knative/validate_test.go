package knative

import (
	"errors"
	"testing"
)

type test struct {
	file        string
	validateErr error
}

func TestValidate(t *testing.T) {
	tests := []test{
		test{
			file:        "../../test/sample.yaml",
			validateErr: nil,
		},
		test{
			file:        "../../test/multiple_containers.yaml",
			validateErr: errors.New("[ERROR] App Runner doesn't support multiple containers: https://github.com/aws/apprunner-roadmap/issues/71"),
		},
	}

	for _, test := range tests {
		svc, err := Parse(test.file)
		if err != nil {
			t.Fatalf("failed test: %v", err)
		}

		err = Validate(svc)
		if err != nil && err.Error() != test.validateErr.Error() {
			t.Fatalf("failed test: %v", err)
		}
	}
}
