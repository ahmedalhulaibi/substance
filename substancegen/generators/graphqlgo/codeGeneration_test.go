package graphqlgo

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func TestGenObjectGormCreateFunc(t *testing.T) {
	var buff bytes.Buffer
	newGenObjType := substancegen.GenObjectType{Name: "Customer", SourceTableName: "Customers"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:       false,
		IsObjectType: false,
		KeyType:      []string{"PRIMARY KEY"},
		ScalarName:   "FirstName",
		ScalarType:   "string",
		Nullable:     false,
	}
	newGenObjType.Properties["FirstName"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["FirstName"].Tags["json"] = append(newGenObjType.Properties["FirstName"].Tags["json"], "firstName")

	newGenObjType.Properties["ShoppingList"] = &substancegen.GenObjectProperty{
		IsList:       true,
		IsObjectType: false,
		KeyType:      []string{""},
		ScalarName:   "ShoppingList",
		ScalarType:   "string",
		Nullable:     false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")
	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	buff = OutputGraphqlSchema(genObjMap)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("type Customer {\n \tFirstName: string!\n\tShoppingList: [string]!\n}\n"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}