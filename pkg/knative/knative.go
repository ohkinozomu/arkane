package knative

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	api "knative.dev/serving/pkg/apis/serving/v1"
)

func Parse(filePath string) (api.Service, error) {
	var service api.Service
	f, err := os.Open(filePath)
	if err != nil {
		return service, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return service, err
	}

	err = yaml.Unmarshal(b, &service)
	if err != nil {
		return service, err
	}
	return service, nil
}
