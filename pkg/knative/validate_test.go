package knative

import (
	"testing"
)

func TestValidate(t *testing.T) {
	svc, err := Parse("../../test/sample.yaml")
	if err != nil {
		t.Fatalf("failed test: %v", err)
	}

	err = Validate(svc)
	if err != nil {
		t.Fatalf("failed test: %v", err)
	}
}
