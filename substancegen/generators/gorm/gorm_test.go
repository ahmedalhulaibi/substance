package gorm

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func TestGenGormObjectTableNameOverrideFunc(t *testing.T) {
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

	GenGormObjectTableNameOverrideFunc(newGenObjType, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("\nfunc (%s) TableName() string {\n\treturn \"%s\"\n}\n", "Customer", newGenObjType.SourceTableName))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}

func TestGenObjectGormCreateFunc(t *testing.T) {
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

	GenObjectGormCreateFunc(newGenObjType, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("\n\nfunc CreateCustomer (db *gorm.DB, newCustomer Customer) []error {\n\treturn db.Create(&newCustomer).GetErrors()\n}"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}

func TestGenObjectGormReadFunc(t *testing.T) {
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

	GenObjectGormReadFunc(newGenObjType, &buff)

	var expectedBuff bytes.Buffer

	//expectedBuff.WriteString(fmt.Sprintf("\n\nfunc GetCustomer (db *gorm.DB, queryCustomer Customer, resultCustomer *Customer) []error {\n\treturn db.Where(&queryCustomer).First(resultCustomer).GetErrors()\n}"))
	expectedBuff.WriteString(fmt.Sprintf(`

func GetCustomer (db *gorm.DB, queryCustomer Customer, resultCustomer *Customer) []error {
	return db.Where(&queryCustomer).First(resultCustomer).GetErrors()
}`))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}
func TestGenObjectGormReadAllFunc(t *testing.T) {
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

	GenObjectGormReadAllFunc(newGenObjType, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("\n\nfunc GetAllCustomer (db *gorm.DB, queryCustomer Customer, resultCustomer *[]Customer) []error {\n\treturn db.Where(&queryCustomer).Find(resultCustomer).GetErrors()\n}"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}

func TestGenObjectGormUpdateFunc(t *testing.T) {
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

	GenObjectGormUpdateFunc(newGenObjType, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("\n\nfunc UpdateCustomer (db *gorm.DB, oldCustomer Customer, newCustomer Customer, resultCustomer *Customer) []error {\n\tvar oldResultCustomer Customer\n\terr := db.Where(&oldCustomer).First(&oldResultCustomer).GetErrors()\n\tif oldResultCustomer.FirstName == newCustomer.FirstName {\n\t\toldResultCustomer = newCustomer\n\t\terr = append(err,db.Save(oldResultCustomer).GetErrors()...)\n\t}\n\terr = append(err,GetCustomer(db, newCustomer, resultCustomer)...)\n\treturn err\n}"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}

func TestGenObjectGormDeleteFunc(t *testing.T) {
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

	GenObjectGormDeleteFunc(newGenObjType, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("\n\nfunc DeleteCustomer (db *gorm.DB, oldCustomer Customer) []error {\n\treturn db.Delete(&oldCustomer).GetErrors()\n}"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}

func TestGenObjectGormCrud(t *testing.T) {
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

	GenObjectGormCrud(newGenObjType, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("\n\nfunc CreateCustomer (db *gorm.DB, newCustomer Customer) []error {\n\treturn db.Create(&newCustomer).GetErrors()\n}"))

	expectedBuff.WriteString(fmt.Sprintf("\n\nfunc GetCustomer (db *gorm.DB, queryCustomer Customer, resultCustomer *Customer) []error {\n\treturn db.Where(&queryCustomer).First(resultCustomer).GetErrors()\n}"))

	expectedBuff.WriteString(fmt.Sprintf("\n\nfunc GetAllCustomer (db *gorm.DB, queryCustomer Customer, resultCustomer *[]Customer) []error {\n\treturn db.Where(&queryCustomer).Find(resultCustomer).GetErrors()\n}"))

	expectedBuff.WriteString(fmt.Sprintf("\n\nfunc UpdateCustomer (db *gorm.DB, oldCustomer Customer, newCustomer Customer, resultCustomer *Customer) []error {\n\tvar oldResultCustomer Customer\n\terr := db.Where(&oldCustomer).First(&oldResultCustomer).GetErrors()\n\tif oldResultCustomer.FirstName == newCustomer.FirstName {\n\t\toldResultCustomer = newCustomer\n\t\terr = append(err,db.Save(oldResultCustomer).GetErrors()...)\n\t}\n\terr = append(err,GetCustomer(db, newCustomer, resultCustomer)...)\n\treturn err\n}"))

	expectedBuff.WriteString(fmt.Sprintf("\n\nfunc DeleteCustomer (db *gorm.DB, oldCustomer Customer) []error {\n\treturn db.Delete(&oldCustomer).GetErrors()\n}"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}
