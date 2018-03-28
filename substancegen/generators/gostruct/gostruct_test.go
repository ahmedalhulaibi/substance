package gostruct

import (
	"bytes"
	"testing"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func TestGenObjectTypeToStructFunc(t *testing.T) {
	var buff bytes.Buffer
	newGenObjType := substancegen.GenObjectType{Name: "Customer", LowerName: "customer", SourceTableName: "Customers"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "FirstName",
		ScalarNameUpper: "FirstName",
		ScalarType:      "string",
		Nullable:        false,
	}
	newGenObjType.Properties["FirstName"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["FirstName"].Tags["json"] = append(newGenObjType.Properties["FirstName"].Tags["json"], "firstName")

	newGenObjType.Properties["ShoppingList"] = &substancegen.GenObjectProperty{
		IsList:          true,
		IsObjectType:    false,
		KeyType:         []string{""},
		ScalarName:      "ShoppingList",
		ScalarNameUpper: "ShoppingList",
		ScalarType:      "string",
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")
	gqlObjects := make(map[string]substancegen.GenObjectType)
	gqlObjects["Customer"] = newGenObjType
	GenObjectTypeToStructFunc(gqlObjects, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString("\ntype Customer struct { \n\tFirstName\tstring\t`json:\"firstName\" `\n\tShoppingList\t[]string\t`json:\"shoppingList\" `\n}\n")

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%v'\n\nReceived\n\n'%v'\n\n", expectedBuff, buff)
	}
}
