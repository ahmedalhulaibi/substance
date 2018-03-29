package gqlschema

import (
	"bytes"
	"html/template"
	"log"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

var graphqlDataTypes map[string]string

func init() {
	graphqlDataTypes = make(map[string]string)
	graphqlDataTypes["int"] = "Int"
	graphqlDataTypes["int8"] = "Int"
	graphqlDataTypes["int16"] = "Int"
	graphqlDataTypes["int32"] = "Int"
	graphqlDataTypes["int64"] = "Int"
	graphqlDataTypes["uint"] = "Int"
	graphqlDataTypes["uint8"] = "Int"
	graphqlDataTypes["uint16"] = "Int"
	graphqlDataTypes["uint32"] = "Int"
	graphqlDataTypes["uint64"] = "Int"
	graphqlDataTypes["byte"] = "Int"
	graphqlDataTypes["rune"] = "Int"
	graphqlDataTypes["bool"] = "Boolean"
	graphqlDataTypes["string"] = "String"
	graphqlDataTypes["float32"] = "Float"
	graphqlDataTypes["float64"] = "Float"
}

/*OutputGraphqlSchema Returns a buffer containing a GraphQL schema in the standard GraphQL schema syntax*/
func OutputGraphqlSchema(gqlObjectTypes map[string]substancegen.GenObjectType) bytes.Buffer {
	var buff bytes.Buffer

	for _, object := range gqlObjectTypes {
		for _, propVal := range object.Properties {
			if propVal.IsObjectType {
				propVal.AltScalarType["gqlschema"] = propVal.ScalarNameUpper
			}
			propVal.AltScalarType["gqlschema"] = graphqlDataTypes[propVal.ScalarType]
		}
	}

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
