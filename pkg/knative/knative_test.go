package knative

import (
	"testing"
)

func TestParse(t *testing.T) {
	_, err := Parse("../../test/sample.yaml")
	if err != nil {
		t.Fatal("failed test")
	}
}
