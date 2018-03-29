package gqlschema

var graphqlSchemaTemplate = `{{range $key, $value := . }}type {{.Name}} { {{range .Properties}}
	{{.ScalarName}}: {{if .IsList}}[{{index .AltScalarType "gqlschema"}}]{{else}}{{index .AltScalarType "gqlschema"}}{{end}}{{if .Nullable}}{{else}}!{{end}}{{end}}
}
{{end}}`
