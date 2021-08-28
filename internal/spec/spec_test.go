package spec

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoad(t *testing.T) {
	spec, err := ioutil.TempFile("", "spec")
	if err != nil {
		t.Fatalf("failed to create temp file:%v\n", err)
	}
	defer os.Remove(spec.Name())

	tests := []struct {
		name     string
		specBody string
		want     Spec
	}{
		{
			name: "specification gen and no option",
			specBody: `
schema:
    name:
        type: string
        gen: random`,
			want: Spec{
				Schema: map[string]Field{
					"name": {
						Typ: "string",
						Gen: "random",
					},
				},
			},
		},
		{
			name: "specification value",
			specBody: `
schema:
    name:
        type: int
        value: 0`,
			want: Spec{
				Schema: map[string]Field{
					"name": {
						Typ:   "int",
						Value: 0,
					},
				},
			},
		},
		{
			name: "specification option",
			specBody: `
schema:
    name:
        type: float 
        gen: random
        option:
            num: 10`,
			want: Spec{
				Schema: map[string]Field{
					"name": {
						Typ: "float",
						Gen: "random",
						Option: map[string]interface{}{
							"num": 10,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := ioutil.WriteFile(spec.Name(), []byte(test.specBody), 0644); err != nil {
				t.Fatalf("failed to write spec body")
			}
			s, err := Load(spec.Name())
			if err != nil {
				t.Errorf("failed to load spec file")
			}
			if diff := cmp.Diff(s, &test.want); diff != "" {
				t.Errorf("unexpected spec. got: %v, want: %v\n", s, test.want)
			}
		})
	}
}
