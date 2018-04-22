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

/*GenPackageImports writes a predefined package and import statement to a buffer*/
func (g Gql) GenPackageImports(dbType string, buff *bytes.Buffer) {
	buff.WriteString("package main\nimport (\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"log\"\n\t\"net/http\"\n\t\"github.com/graphql-go/graphql\"\n\t\"github.com/graphql-go/handler\"")

	if importVal, exists := g.GraphqlDbTypeImports[dbType]; exists {
		buff.WriteString(importVal)
	}
	buff.WriteString("\n)")
}

/*GenerateGraphqlGoTypesFunc takes a map of gen objects and outputs graphql-go types to a buffer*/
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

/*GenGraphqlGoMainFunc generates the main function (entrypoint) for the graphql-go server*/
func GenGraphqlGoMainFunc(dbType string, connectionString string, gqlObjectTypes map[string]substancegen.GenObjectType, buff *bytes.Buffer) {
	var sampleQuery bytes.Buffer
	GenGraphqlGoSampleQuery(gqlObjectTypes, &sampleQuery)

	// buff.WriteString(fmt.Sprintf("\nvar DB *gorm.DB\n\n"))
	// buff.WriteString(fmt.Sprintf("\nfunc main() {\n\n\tDB, _ = gorm.Open(\"%s\",\"%s\")\n\tdefer DB.Close()\n\n\t", dbType, connectionString))

	// buff.WriteString(fmt.Sprintf("\n\tfmt.Println(\"Test with Get\t: curl -g 'http://localhost:8080/graphql?query={%s}'\")", sampleQuery.String()))

	// buff.WriteString(graphqlGoMainFunc)

	// buff.WriteString("\n}\n")

	mainData := struct {
		DbType           string
		ConnectionString string
		SampleQuery      string
	}{
		dbType,
		connectionString,
		sampleQuery.String(),
	}
	tmpl := template.New("graphqlGoMainFunc")
	tmpl, err := tmpl.Parse(graphqlGoMainFunc)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	//print schema
	err1 := tmpl.Execute(buff, mainData)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
	}
}

/*GenGraphqlGoFieldsFunc generates a basic graphql-go queries
to retrieve the first element of each object type (and its associations) from a database*/
func GenGraphqlGoFieldsFunc(gqlObjectTypes map[string]substancegen.GenObjectType, buff *bytes.Buffer) {
	tmpl := template.New("graphqlFields")
	tmpl, err := tmpl.Parse(graphqlGoFieldsTemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	//print schema
	err1 := tmpl.Execute(buff, gqlObjectTypes)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
	}
}

/*GenGraphqlGoSampleQuery generates a sample graphql query based on the given objects*/
func GenGraphqlGoSampleQuery(gqlObjectTypes map[string]substancegen.GenObjectType, buff *bytes.Buffer) {
	tmpl := template.New("graphqlQuery")
	tmpl, err := tmpl.Parse(graphqlQueryTemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	//print schema
	err1 := tmpl.Execute(buff, gqlObjectTypes)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}

	bufferString := buff.String()
	bufferString = strings.Replace(bufferString, " ", "", -1)
	buff.Reset()
	buff.WriteString(bufferString)
}

var GoNumericAliasTypeMap map[string]string

/*InitGoNumericAliasTypeMap initializes gqlPlugin data for go alias numeric type mapping
This is currently used in the graphql-go generation this is required as graphql-go implements graphql.Int as an int and not int32, int64, uint32, etc*/
func InitGoNumericAliasTypeMap() {
	GoNumericAliasTypeMap = make(map[string]string, 16)
	GoNumericAliasTypeMap["int"] = "int"
	GoNumericAliasTypeMap["int8"] = "int"
	GoNumericAliasTypeMap["int16"] = "int"
	GoNumericAliasTypeMap["int32"] = "int"
	GoNumericAliasTypeMap["int64"] = "int"
	GoNumericAliasTypeMap["uint"] = "int"
	GoNumericAliasTypeMap["uint8"] = "int"
	GoNumericAliasTypeMap["uint16"] = "int"
	GoNumericAliasTypeMap["uint32"] = "int"
	GoNumericAliasTypeMap["uint64"] = "int"
	GoNumericAliasTypeMap["float32"] = "float32"
	GoNumericAliasTypeMap["float64"] = "float32"
}

/*GetGoNumericAliasType returns the alias numeric type for another specific numeric type
For example given:
 int32 return int
 int64 return int
 int   return int
This is currently used in the graphql-go generation this is required as graphql-go implements graphql.Int as an int and not int32, int64, uint32, etc */
func GetGoNumericAliasType(goType string) string {
	if GoNumericAliasTypeMap == nil {
		InitGoNumericAliasTypeMap()
	}
	if val, ok := GoNumericAliasTypeMap[goType]; ok {
		return val
	}
	return goType
}

/*GenGraphqlGoFieldsFunc generates a basic graphql-go queries
to retrieve the first element of each object type (and its associations) from a database*/
func GenGraphqlGoFieldsGetFunc(gqlObjectTypes map[string]substancegen.GenObjectType, buff *bytes.Buffer) {
	funcMap := template.FuncMap{
		"goType": GetGoNumericAliasType,
	}
	tmpl := template.New("graphqlFieldsGet").Funcs(funcMap)

	tmpl, err := tmpl.Parse(graphqlGoQueryFieldsGetTemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	//print schema
	err1 := tmpl.Execute(buff, gqlObjectTypes)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
	}
}
