package graphqlgo

import (
	"bytes"
	"fmt"
	"unicode"

	"github.com/jinzhu/inflection"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func (g Gql) OutputCodeFunc(dbType string, connectionString string, gqlObjectTypes map[string]substancegen.GenObjectType) bytes.Buffer {
	var buff bytes.Buffer

	g.GenPackageImports(dbType, &buff)
	//print schema
	g.AddJSONTagsToProperties(gqlObjectTypes)
	for _, value := range gqlObjectTypes {
		g.GenObjectTypeToStringFunc(value, &buff)
		g.GenGormObjectTableNameOverrideFunc(value, &buff)
		g.GenGraphqlGoTypeFunc(value, &buff)
	}
	buff.WriteString(GraphqlGoExecuteQueryFunc)
	g.GenGraphqlGoMainFunc(dbType, connectionString, gqlObjectTypes, &buff)
	return buff
}

func (g Gql) AddJSONTagsToProperties(gqlObjectTypes map[string]substancegen.GenObjectType) {

	for _, value := range gqlObjectTypes {
		for _, propVal := range value.Properties {
			propVal.Tags["json"] = append(propVal.Tags["json"], propVal.ScalarName)
		}
	}
}

func (g Gql) GenPackageImports(dbType string, buff *bytes.Buffer) {
	buff.WriteString("package main\nimport (\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"log\"\n\t\"net/http\"\n\t\"github.com/graphql-go/graphql\"")

	if importVal, exists := g.GraphqlDbTypeImports[dbType]; exists {
		buff.WriteString(importVal)
	}
	buff.WriteString("\n)")
}

func (g Gql) GenObjectTypeToStringFunc(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	gqlObjectTypeNameSingular := inflection.Singular(gqlObjectType.Name)
	buff.WriteString(fmt.Sprintf("\ntype %s struct {\n", gqlObjectTypeNameSingular))
	for _, property := range gqlObjectType.Properties {
		g.GenObjectPropertyToStringFunc(property, buff)
	}
	buff.WriteString("}\n")
}

func (g Gql) GenObjectPropertyToStringFunc(gqlObjectProperty substancegen.GenObjectProperty, buff *bytes.Buffer) {

	a := []rune(gqlObjectProperty.ScalarName)
	a[0] = unicode.ToUpper(a[0])
	gqlObjectPropertyNameUpper := string(a)
	if gqlObjectProperty.IsList {
		buff.WriteString(fmt.Sprintf("\t%s\t[]%s\t", gqlObjectPropertyNameUpper, gqlObjectProperty.ScalarType))
	} else {
		buff.WriteString(fmt.Sprintf("\t%s\t%s\t", gqlObjectPropertyNameUpper, gqlObjectProperty.ScalarType))
	}
	g.GenObjectTagToStringFunc(gqlObjectProperty.Tags, buff)
	buff.WriteString("\n")
}

func (g Gql) GenObjectTagToStringFunc(genObjectTags substancegen.GenObjectTag, buff *bytes.Buffer) {
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

func (g Gql) GenGormObjectTableNameOverrideFunc(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	gqlObjectTypeNameSingular := inflection.Singular(gqlObjectType.Name)
	buff.WriteString(fmt.Sprintf("\nfunc (%s) TableName() string {\n\treturn \"%s\"\n}\n", gqlObjectTypeNameSingular, gqlObjectType.Name))
}

func (g Gql) GenGraphqlGoTypeFunc(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	a := []rune(inflection.Singular(gqlObjectType.Name))
	a[0] = unicode.ToLower(a[0])
	gqlObjectTypeNameLowCamel := string(a)
	gqlObjectTypeNameSingular := inflection.Singular(gqlObjectType.Name)
	buff.WriteString(fmt.Sprintf("\nvar %sType = graphql.NewObject(\n\tgraphql.ObjectConfig{\n\t\tName: \"%s\",\n\t\tFields: graphql.Fields{\n\t\t\t", gqlObjectTypeNameLowCamel, gqlObjectTypeNameSingular))

	for _, property := range gqlObjectType.Properties {
		g.GenGraphqlGoTypePropertyFunc(property, buff)
	}

	buff.WriteString(fmt.Sprintf("\n\t\t},\n\t},\n)\n"))
}

func (g Gql) GenGraphqlGoTypePropertyFunc(gqlObjectProperty substancegen.GenObjectProperty, buff *bytes.Buffer) {
	gqlPropertyTypeName := g.ResolveGraphqlGoFieldType(gqlObjectProperty)
	buff.WriteString(fmt.Sprintf("\n\t\t\t\"%s\": &graphql.Field{\n\t\t\t\tType: %s,\n\t\t\t},", gqlObjectProperty.ScalarName, gqlPropertyTypeName))
}

func (g Gql) ResolveGraphqlGoFieldType(gqlObjectProperty substancegen.GenObjectProperty) string {
	var gqlPropertyTypeName string

	if gqlObjectProperty.IsObjectType {
		a := []rune(inflection.Singular(gqlObjectProperty.ScalarName))
		a[0] = unicode.ToLower(a[0])
		gqlPropertyTypeName = fmt.Sprintf("%sType", string(a))
	} else {
		gqlPropertyTypeName = g.GraphqlDataTypes[gqlObjectProperty.ScalarType]
	}

	if gqlObjectProperty.IsList {
		gqlPropertyTypeName = fmt.Sprintf("graphql.NewList(%s)", gqlPropertyTypeName)
	}

	if !gqlObjectProperty.Nullable {
		gqlPropertyTypeName = fmt.Sprintf("graphql.NewNonNull(%s)", gqlPropertyTypeName)
	}

	return gqlPropertyTypeName
}

func (g Gql) GenGraphqlGoMainFunc(dbType string, connectionString string, gqlObjectTypes map[string]substancegen.GenObjectType, buff *bytes.Buffer) {
	buff.WriteString(fmt.Sprintf("\nvar DB *gorm.DB\n\n"))
	buff.WriteString(fmt.Sprintf("\nfunc main() {\n\n\tDB, err := gorm.Open(\"%s\",\"%s\")\n\tdefer DB.Close()\n\n\t", dbType, connectionString))
	sampleQuery := g.GenGraphqlGoSampleQuery(gqlObjectTypes)
	buff.WriteString(fmt.Sprintf("\n\tfmt.Println(\"Test with Get\t: curl -g 'http://localhost:8080/graphql?query={%s}'\")", sampleQuery.String()))

	buff.WriteString("\n\tfields := graphql.Fields{")
	for _, value := range gqlObjectTypes {
		g.GenGraphqlGoQueryFieldsFunc(value, buff)
	}
	buff.WriteString("\n\t\t}")
	buff.WriteString(GraphqlGoMainConfig)

	buff.WriteString("\n}\n")
}

func (g Gql) GenGraphqlGoQueryFieldsFunc(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	gqlObjectTypeNameSingular := inflection.Singular(gqlObjectType.Name)
	a := []rune(inflection.Singular(gqlObjectType.Name))
	a[0] = unicode.ToLower(a[0])
	gqlObjectTypeNameLowCamel := string(a)
	buff.WriteString(fmt.Sprintf("\n\t\t\"%s\": &graphql.Field{\n\t\t\tType: %sType,", gqlObjectTypeNameSingular, gqlObjectTypeNameLowCamel))
	buff.WriteString(fmt.Sprintf("\n\t\t\tResolve: func(p graphql.ResolveParams) (interface{}, error) {"))
	buff.WriteString(fmt.Sprintf("\n\t\t\t\t%s := %s{}", gqlObjectTypeNameLowCamel, gqlObjectTypeNameSingular))
	buff.WriteString(fmt.Sprintf("\n\t\t\t\tDB.First(&%s)", gqlObjectTypeNameLowCamel))

	for _, propVal := range gqlObjectType.Properties {
		if propVal.IsObjectType {
			a := []rune(propVal.ScalarName)
			a[0] = unicode.ToLower(a[0])
			propValNameLowCamel := string(a)
			b := []rune(propVal.ScalarName)
			b[0] = unicode.ToUpper(b[0])
			propValNameUpperCamel := string(b)
			if propVal.IsList {
				buff.WriteString(fmt.Sprintf("\n\t\t\t\t%s := []%s{}", propValNameLowCamel, propVal.ScalarType))

				buff.WriteString(fmt.Sprintf("\n\t\t\t\tDB.Model(&%s).Association(\"%s\").Find(&%s)", gqlObjectTypeNameLowCamel, propVal.ScalarName, propValNameLowCamel))

				buff.WriteString(fmt.Sprintf("\n\t\t\t\t%s.%s = append(%s.%s, %s...)", gqlObjectTypeNameLowCamel, propValNameUpperCamel, gqlObjectTypeNameLowCamel, propValNameUpperCamel, propValNameLowCamel))
			} else {
				buff.WriteString(fmt.Sprintf("\n\t\t\t\t%s := %s{}", propValNameLowCamel, propVal.ScalarType))

				buff.WriteString(fmt.Sprintf("\n\t\t\t\tDB.Model(&%s).Association(\"%s\").Find(&%s)", gqlObjectTypeNameLowCamel, propVal.ScalarName, propValNameLowCamel))

				buff.WriteString(fmt.Sprintf("\n\t\t\t\t%s.%s = %s", gqlObjectTypeNameLowCamel, propValNameUpperCamel, propValNameLowCamel))
			}
		}
	}
	buff.WriteString(fmt.Sprintf("\n\t\t\t\treturn %s, nil", gqlObjectTypeNameLowCamel))
	buff.WriteString("\n\t\t\t},")
	buff.WriteString("\n\t\t},")
}

func (g Gql) GenGraphqlGoSampleQuery(gqlObjectTypes map[string]substancegen.GenObjectType) bytes.Buffer {
	var buff bytes.Buffer
	for _, gqlObjectType := range gqlObjectTypes {
		g.GenGraphlGoSampleObjectQuery(gqlObjectTypes, gqlObjectType, &buff)
	}
	return buff
}

func (g Gql) GenGraphlGoSampleObjectQuery(gqlObjectTypes map[string]substancegen.GenObjectType, gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	gqlObjectTypeNameSingular := inflection.Singular(gqlObjectType.Name)
	buff.WriteString(fmt.Sprintf("%s{", gqlObjectTypeNameSingular))
	for _, propVal := range gqlObjectType.Properties {
		if !propVal.IsObjectType {
			buff.WriteString(fmt.Sprintf("%s,", propVal.ScalarName))
		}
	}
	buff.WriteString("},")
}

func (g Gql) OutputGraphqlSchema(gqlObjectTypes map[string]substancegen.GenObjectType) bytes.Buffer {
	var buff bytes.Buffer
	//print schema
	for _, value := range gqlObjectTypes {
		buff.WriteString(fmt.Sprintf("type %s {\n", value.Name))
		for _, propVal := range value.Properties {
			nullSymbol := "!"
			if propVal.Nullable {
				nullSymbol = ""
			}
			if propVal.IsList {
				buff.WriteString(fmt.Sprintf("\t %s: [%s]%s\n", propVal.ScalarName, propVal.ScalarType, nullSymbol))
			} else {
				buff.WriteString(fmt.Sprintf("\t %s: %s%s\n", propVal.ScalarName, propVal.ScalarType, nullSymbol))
			}
		}
		buff.WriteString(fmt.Sprintf("}\n"))
	}
	return buff
}

func (g Gql) GenObjectGormCrud(gqlObjectType substancegen.GenObjectType, buff *bytes.Buffer) {
	gqlObjectTypeNameSingular := inflection.Singular(gqlObjectType.Name)
	var primaryKeyColumn string
	for index, propVal := range gqlObjectType.Properties {
		if stringInSlice("p", propVal.KeyType) || stringInSlice("PRIMARY KEY", propVal.KeyType) {
			primaryKeyColumn = index
			break
		}
	}

	buff.WriteString(fmt.Sprintf("\n\nfunc Create%s (db *gorm.DB, new%s %s) {\n\tdb.Create(&new%s)\n}",
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular))

	buff.WriteString(fmt.Sprintf("\n\nfunc Get%s (db *gorm.DB, query%s %s, result%s *%s) {\n\tdb.Where(&query%s).First(result%s)\n}",
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular))

	buff.WriteString(fmt.Sprintf("\n\nfunc Update%s (db *gorm.DB, old%s %s, new%s %s, result%s *%s) {\n\tvar oldResult%s %s\n\tdb.Where(&old%s).First(&oldResult%s)\n\tif oldResult%s.%s == new%s.%s {\n\t\toldResult%s = new%s\n\t\tdb.Save(oldResult%s)\n\t}\n\tGet%s(db, new%s, result%s)\n}",
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		primaryKeyColumn,
		gqlObjectTypeNameSingular,
		primaryKeyColumn,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular))

	buff.WriteString(fmt.Sprintf("\n\nfunc Delete%s (db *gorm.DB, old%s %s) {\n\tdb.Delete(&old%s)\n}",
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular,
		gqlObjectTypeNameSingular))
}
