package gen

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/smiyaguchi/foundry/internal/spec"
)

type Generator interface {
	Generate(option GenOption) (interface{}, error)
}

type GenOption map[string]interface{}

func Convert(spec *spec.Spec) (string, error) {
	o := make(map[string]interface{})

	for key, field := range spec.Schema {
		if field.Value != "" {
			v, err := convertValue(field.Typ, field.Value)
			if err != nil {
				return "", err
			}
			o[key] = v
			continue
		}
		gen, err := NewGenerator(field)
		if err != nil {
			return "", fmt.Errorf("failed to new generator: %v\n", err)
		}
		v, err := gen.Generate(field.Option)
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
	gen := strings.ToLower(field.Gen)
	switch gen {
	case "random":
		return &genSeed{}, nil
	default:
		return &genDefault{typ: strings.ToLower(field.Typ)}, nil
	}
}

func convertValue(typ, value string) (interface{}, error) {
	t := strings.ToLower(typ)
	switch t {
	case "string":
		return value, nil
	case "int":
		i, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return i, nil
	case "float":
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		return f, nil
	case "bool":
		b, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		return b, nil
	default:
		return nil, fmt.Errorf("not support value type: %s\n", value)
	}
}
