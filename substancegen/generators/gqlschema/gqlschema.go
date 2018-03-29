package gqlschema

import (
	"bytes"
	"html/template"
	"log"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

/*OutputGraphqlSchema Returns a buffer containing a GraphQL schema in the standard GraphQL schema syntax*/
func OutputGraphqlSchema(gqlObjectTypes map[string]substancegen.GenObjectType) bytes.Buffer {
	var buff bytes.Buffer

	graphqlSchemaTemplate := "{{range $key, $value := . }}type {{.Name}} {\n {{range .Properties}}\t{{.ScalarName}}: {{if .IsList}}[{{.ScalarType}}]{{else}}{{.ScalarType}}{{end}}{{if .Nullable}}{{else}}!{{end}}\n{{end}}}\n{{end}}"
	tmpl := template.New("graphqlSchema")
	tmpl, err := tmpl.Parse(graphqlSchemaTemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return buff
	}
	//print schema
	err1 := tmpl.Execute(&buff, gqlObjectTypes)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return buff
	}

	return buff
}
