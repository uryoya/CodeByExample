package lang

import (
	"strings"
	"testing"
	"text/template"
	"uryoya/code_by_example/test"
)

func TestTextTemplate(t *testing.T) {
	type Inventory struct {
		Material string
		Count    uint
	}

	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	var s strings.Builder
	err = tmpl.Execute(&s, sweaters)
	if err != nil {
		panic(err)
	}
	test.Expect(t, s.String() == "17 items are made of wool")
}
