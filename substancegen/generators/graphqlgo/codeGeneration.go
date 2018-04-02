package graphqlgo

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
	"unicode"

	"github.com/jinzhu/inflection"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func (g Gql) GenPackageImports(dbType string, buff *bytes.Buffer) {
	buff.WriteString("package main\nimport (\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"log\"\n\t\"net/http\"\n\t\"github.com/graphql-go/graphql\"\n\t\"github.com/graphql-go/handler\"")

	if importVal, exists := g.GraphqlDbTypeImports[dbType]; exists {
		buff.WriteString(importVal)
	}
	buff.WriteString("\n)")
}

func (g Gql) GenerateGraphqlGoTypesFunc(gqlObjectTypes map[string]substancegen.GenObjectType, buff *bytes.Buffer) {
	for _, value := range gqlObjectTypes {
		for _, propVal := range value.Properties {
			if propVal.IsObjectType {
				a := []rune(inflection.Singular(propVal.ScalarName))
				a[0] = unicode.ToLower(a[0])
				propVal.AltScalarType["graphql-go"] = fmt.Sprintf("%sType", string(a))
			} else {
				propVal.AltScalarType["graphql-go"] = g.GraphqlDataTypes[propVal.ScalarType]
			}

			if propVal.IsList {
				propVal.AltScalarType["graphql-go"] = fmt.Sprintf("graphql.NewList(%s)", propVal.AltScalarType["graphql-go"])
			}

			if !propVal.Nullable {
				propVal.AltScalarType["graphql-go"] = fmt.Sprintf("graphql.NewNonNull(%s)", propVal.AltScalarType["graphql-go"])
			}
		}
	}
	tmpl := template.New("graphqlTypes")
	tmpl, err := tmpl.Parse(graphqlTypesTemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	err1 := tmpl.Execute(buff, gqlObjectTypes)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}
}

func GenGraphqlGoMainFunc(dbType string, connectionString string, gqlObjectTypes map[string]substancegen.GenObjectType, buff *bytes.Buffer) {
	buff.WriteString(fmt.Sprintf("\nvar DB *gorm.DB\n\n"))
	buff.WriteString(fmt.Sprintf("\nfunc main() {\n\n\tDB, _ = gorm.Open(\"%s\",\"%s\")\n\tdefer DB.Close()\n\n\t", dbType, connectionString))
	sampleQuery := GenGraphqlGoSampleQuery(gqlObjectTypes)
	buff.WriteString(fmt.Sprintf("\n\tfmt.Println(\"Test with Get\t: curl -g 'http://localhost:8080/graphql?query={%s}'\")", sampleQuery.String()))

	buff.WriteString(GraphqlGoMainConfig)

	buff.WriteString("\n}\n")
}

func GenGraphqlGoFieldsFunc(gqlObjectTypes map[string]substancegen.GenObjectType) bytes.Buffer {
	var buff bytes.Buffer

	tmpl := template.New("graphqlFields")
	tmpl, err := tmpl.Parse(graphqlGoFieldsTemplate)
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

func GenGraphqlGoSampleQuery(gqlObjectTypes map[string]substancegen.GenObjectType) bytes.Buffer {
	var buff bytes.Buffer

	tmpl := template.New("graphqlQuery")
	tmpl, err := tmpl.Parse(graphqlQueryTemplate)
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

	bufferString := buff.String()
	bufferString = strings.Replace(bufferString, " ", "", -1)
	buff.Reset()
	buff.WriteString(bufferString)
	return buff
}
