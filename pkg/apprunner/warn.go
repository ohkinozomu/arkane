package apprunner

import "fmt"

func printWarn() {
	msg := `
======================================== WARN ========================================
App Runner doesn't support distributed traffic, so traffic field is ignored.
https://github.com/aws/apprunner-roadmap/issues/40
https://github.com/knative/specs/blob/main/specs/serving/knative-api-specification-1.0.md#spec
======================================================================================	
`
	fmt.Println(msg)
}
