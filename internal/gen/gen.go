package gen

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/smiyaguchi/foundry/internal/spec"
)

type Generator interface {
	Generate(option GenOption) (interface{}, error)
}

type GenOption map[string]interface{}

func Convert(spec *spec.Spec, num int) (string, error) {
	var result string
	for i := 0; i < num; i++ {
		v, err := schemaToValue(spec.Schema)
		if err != nil {
			return "", err
		}
		b, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("failed to marshal json: %v\n", err)
		}
		result += string(b)
	}

	return result, nil
}

func schemaToValue(schema spec.Schema) (map[string]interface{}, error) {
	o := make(map[string]interface{})
	for key, field := range schema {
		if len(field.Schema) != 0 {
			v, err := schemaToValue(field.Schema)
			if err != nil {
				return nil, err
			}
			o[key] = v
			continue
		}

		if field.Value != nil {
			v, err := convertValue(field.Typ, field.Value)
			if err != nil {
				return nil, err
			}
			o[key] = v
			continue
		}
		gen, err := NewGenerator(field)
		if err != nil {
			return nil, fmt.Errorf("failed to new generator: %v\n", err)
		}
		v, err := gen.Generate(field.Option)
		if err != nil {
			return nil, fmt.Errorf("failed to generate data: %v\n", err)
		}
		o[key] = v
	}
	return o, nil
}

func NewGenerator(field spec.Field) (Generator, error) {
	gen := strings.ToLower(field.Gen)
	switch gen {
	case "ipv4":
		return &genIPv4{}, nil
	case "ipv6":
		return &genIPv6{}, nil
	case "random":
		return &genSeed{}, nil
	case "uuid":
		return &genUUID{}, nil
	default:
		return &genDefault{typ: strings.ToLower(field.Typ)}, nil
	}
}

func convertValue(typ string, value interface{}) (interface{}, error) {
	t := strings.ToLower(typ)
	switch t {
	case "string":
		if v, ok := value.(string); ok {
			return v, nil
		}
	case "int":
		if v, ok := value.(int); ok {
			return v, nil
		}
	case "float":
		if v, ok := value.(float64); ok {
			return v, nil
		}
	case "bool":
		if v, ok := value.(bool); ok {
			return v, nil
		}
	default:
		return nil, fmt.Errorf("not support value type: %s\n", value)
	}
	return nil, fmt.Errorf("failed to cast type: %s, value:%v\n", t, value)
}
