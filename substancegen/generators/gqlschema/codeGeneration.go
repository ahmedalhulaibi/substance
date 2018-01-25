package gqlgraphqlator

import (
	"bytes"
	"fmt"
	"strings"

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
	}
	fmt.Print(buff.String())
	return buff
}

func (g gql) GenObjectTypeToStringFunc(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	buff.WriteString(fmt.Sprintf("\ntype %s struct {\n", strings.TrimSuffix(gqlObjectType.Name, "s")))
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
	buff.WriteString(fmt.Sprintf("\nfunc (%s) TableName() string {\n\treturn \"%s\"\n}\n", strings.TrimSuffix(gqlObjectType.Name, "s"), gqlObjectType.Name))
}

func (g gql) GenGraphqlGoTypesFunc(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	buff.WriteString(fmt.Sprintf(""))
}
