package gen

import (
	"encoding/json"
	"fmt"

	"github.com/smiyaguchi/foundry/internal/spec"
)

type Generator interface {
	Generate(spec spec.Field) (interface{}, error)
}

func Convert(spec *spec.Spec) (string, error) {
	o := make(map[string]interface{})
	fmt.Println(spec.Schema)

	for key, value := range spec.Schema {
		gen, err := NewGenerator(value)
		if err != nil {
			return "", fmt.Errorf("failed to new generator: %v\n", err)
		}
		v, err := gen.Generate(value)
		if err != nil {
			return "", fmt.Errorf("failed to generate data: %v\n", err)
		}
		o[key] = v
	}

	b, err := json.Marshal(o)
	if err != nil {
		return "", fmt.Errorf("failed to marshal json: %v\n", err)
	}
	return string(b), nil
}

func NewGenerator(field spec.Field) (Generator, error) {
	switch field.Gen {
	case "random":
		return &genSeed{Num: 10}, nil
	default:
		return nil, fmt.Errorf("not support generator: %s\n", field.Gen)
	}
}
