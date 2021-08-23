package spec

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Spec struct {
	Schema interface{}
}

func Load(filename string) (*Spec, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	spec := Spec{}
	err = yaml.Unmarshal(buf, &spec)
	if err != nil {
		return nil, err
	}
	return &spec, nil
}
