package spec

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Spec struct {
	Schema map[string]Field `yaml:"schema"`
}

type Field struct {
	Typ    string                 `yaml:"type"`
	Gen    string                 `yaml:"gen"`
	Option map[string]interface{} `yaml:"option"`
	Schema map[string]Field       `yaml:"schema,omitempty"`
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
