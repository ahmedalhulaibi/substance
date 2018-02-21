package gostruct

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func TestGenObjectTypeToStructFunc(t *testing.T) {
	var buff bytes.Buffer
	newGenObjType := substancegen.GenObjectType{Name: "Customer"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:       false,
		IsObjectType: false,
		KeyType:      []string{""},
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

	GenObjectTypeToStructFunc(newGenObjType, &buff)

	var expectedBuff bytes.Buffer
	expectedBuff.WriteString(fmt.Sprintf("\ntype Customer struct {\n"))
	expectedBuff.WriteString(fmt.Sprintf("\tFirstName\tstring\t`json:\"firstName\" `"))
	expectedBuff.WriteString("\n")
	expectedBuff.WriteString(fmt.Sprintf("\tShoppingList\t[]string\t`json:\"shoppingList\" `"))
	expectedBuff.WriteString("\n")
	expectedBuff.WriteString("}\n")

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}

func TestGenObjectPropertyToStringFunc(t *testing.T) {
	var buff bytes.Buffer
	newGenObjType := substancegen.GenObjectType{Name: "Customer"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:       false,
		IsObjectType: false,
		KeyType:      []string{""},
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

	GenObjectPropertyToStringFunc(*newGenObjType.Properties["FirstName"], &buff)
	GenObjectPropertyToStringFunc(*newGenObjType.Properties["ShoppingList"], &buff)

	var expectedBuff bytes.Buffer
	expectedBuff.WriteString(fmt.Sprintf("\tFirstName\tstring\t`json:\"firstName\" `"))
	expectedBuff.WriteString("\n")
	expectedBuff.WriteString(fmt.Sprintf("\tShoppingList\t[]string\t`json:\"shoppingList\" `"))
	expectedBuff.WriteString("\n")

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}

func TestGenObjectTagToStringFunc(t *testing.T) {
	var buff bytes.Buffer
	newGenObjType := substancegen.GenObjectType{Name: "Customer"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:       false,
		IsObjectType: false,
		KeyType:      []string{""},
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

	GenObjectTagToStringFunc(newGenObjType.Properties["FirstName"].Tags, &buff)
	GenObjectTagToStringFunc(newGenObjType.Properties["ShoppingList"].Tags, &buff)

	var expectedBuff bytes.Buffer
	expectedBuff.WriteString(fmt.Sprintf("`json:\"firstName\" `"))
	expectedBuff.WriteString(fmt.Sprintf("`json:\"shoppingList\" `"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}
