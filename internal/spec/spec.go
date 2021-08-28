package spec

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Spec struct {
	Schema Schema `yaml:"schema"`
}

type Schema map[string]Field

type Field struct {
	Typ    string                 `yaml:"type"`
	Gen    string                 `yaml:"gen"`
	Value  interface{}            `yaml:"value"`
	Option map[string]interface{} `yaml:"option"`
	Schema Schema                 `yaml:"schema,omitempty"`
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
