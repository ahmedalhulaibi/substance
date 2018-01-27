package graphqlgo

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func (g gql) OutputCodeFunc(gqlObjectTypes map[string]substancegen.GenObjectType) bytes.Buffer {
	var buff bytes.Buffer
	//print schema
	for _, value := range gqlObjectTypes {
		for _, propVal := range value.Properties {
			propVal.Tags["json"] = append(propVal.Tags["json"], propVal.ScalarName)
		}
		g.GenObjectTypeToStringFunc(value, &buff)
		g.GenGormObjectTableNameOverrideFunc(value, &buff)
		g.GenGraphqlGoTypeFunc(value, &buff)
	}
	fmt.Print(buff.String())
	return buff
}

func (g gql) GenObjectTypeToStringFunc(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	gqlObjectTypeNameSingular := strings.TrimSuffix(gqlObjectType.Name, "s")
	buff.WriteString(fmt.Sprintf("\ntype %s struct {\n", gqlObjectTypeNameSingular))
	for _, property := range gqlObjectType.Properties {
		g.GenObjectPropertyToStringFunc(property, buff)
	}
	buff.WriteString("}\n")
}

func (g gql) GenObjectPropertyToStringFunc(gqlObjectProperty substancegen.GenObjectProperty, buff *bytes.Buffer) {
	if gqlObjectProperty.IsList {
		buff.WriteString(fmt.Sprintf("\t%s\t[]%s\t", gqlObjectProperty.ScalarName, gqlObjectProperty.ScalarType))
	} else {
		buff.WriteString(fmt.Sprintf("\t%s\t%s\t", gqlObjectProperty.ScalarName, gqlObjectProperty.ScalarType))
	}
	g.GenObjectTagToStringFunc(gqlObjectProperty.Tags, buff)
	buff.WriteString("\n")
}

func (g gql) GenObjectTagToStringFunc(genObjectTags substancegen.GenObjectTag, buff *bytes.Buffer) {
	buff.WriteString("`")
	for key, tags := range genObjectTags {
		buff.WriteString(fmt.Sprintf("%s:\"", key))
		for _, tag := range tags {
			buff.WriteString(fmt.Sprintf("%s", tag))
		}
		buff.WriteString("\" ")
	}
	buff.WriteString("`")
}

func (g gql) GenGormObjectTableNameOverrideFunc(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	gqlObjectTypeNameSingular := strings.TrimSuffix(gqlObjectType.Name, "s")
	buff.WriteString(fmt.Sprintf("\nfunc (%s) TableName() string {\n\treturn \"%s\"\n}\n", gqlObjectTypeNameSingular, gqlObjectType.Name))
}

func (g gql) GenGraphqlGoTypeFunc(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	a := []rune(strings.TrimSuffix(gqlObjectType.Name, "s"))
	a[0] = unicode.ToLower(a[0])
	gqlObjectTypeNameLowCamel := string(a)
	gqlObjectTypeNameSingular := strings.TrimSuffix(gqlObjectType.Name, "s")
	buff.WriteString(fmt.Sprintf("\nvar %sType = graphql.NewObject(\n\tgraphql.ObjectConfig{\n\t\tName: \"%s\",\n\t\tFields: graphql.Fields{\n\t\t\t", gqlObjectTypeNameLowCamel, gqlObjectTypeNameSingular))

	for _, property := range gqlObjectType.Properties {
		g.GenGraphqlGoTypePropertyFunc(property, buff)
	}

	buff.WriteString(fmt.Sprintf("\n\t\t},\n\t},\n)\n"))
}

func (g gql) GenGraphqlGoTypePropertyFunc(gqlObjectProperty substancegen.GenObjectProperty, buff *bytes.Buffer) {
	var gqlPropertyTypeName string
	if gqlObjectProperty.IsObjectType {
		a := []rune(strings.TrimSuffix(gqlObjectProperty.ScalarName, "s"))
		a[0] = unicode.ToLower(a[0])
		gqlPropertyTypeName = fmt.Sprintf("%sType", string(a))
	} else {
		gqlPropertyTypeName = g.GetGraphqlDataType(gqlObjectProperty.ScalarType)
	}

	buff.WriteString(fmt.Sprintf("\n\t\t\t\"%s\": &graphql.Field{\n\t\t\t\tType: %s,\n\t\t\t},", gqlObjectProperty.ScalarName, gqlPropertyTypeName))
}

func (g gql) GetGraphqlDataType(goDataType string) string {
	graphqlTypes := make(map[string]string)
	graphqlTypes["int"] = "graphql.Int"
	graphqlTypes["int8"] = "graphql.Int"
	graphqlTypes["int16"] = "graphql.Int"
	graphqlTypes["int32"] = "graphql.Int"
	graphqlTypes["int64"] = "graphql.Int"
	graphqlTypes["uint"] = "graphql.Int"
	graphqlTypes["uint8"] = "graphql.Int"
	graphqlTypes["uint16"] = "graphql.Int"
	graphqlTypes["uint32"] = "graphql.Int"
	graphqlTypes["uint64"] = "graphql.Int"
	graphqlTypes["byte"] = "graphql.Int"
	graphqlTypes["rune"] = "graphql.Int"
	graphqlTypes["bool"] = "graphql.Boolean"
	graphqlTypes["string"] = "graphql.String"
	graphqlTypes["float32"] = "graphql.Float"
	graphqlTypes["float64"] = "graphql.Float"
	return graphqlTypes[goDataType]
}
