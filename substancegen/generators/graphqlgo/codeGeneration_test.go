package graphqlgo

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func TestOutputGraphqlSchemaFunc(t *testing.T) {
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
	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	buff = OutputGraphqlSchema(genObjMap)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("type Customer {\n \tFirstName: string!\n\tShoppingList: [string]!\n}\n"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}
func TestGenGraphqlGoSampleQueryFunc(t *testing.T) {
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
	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	buff = GenGraphqlGoSampleQuery(genObjMap)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("Customer{FirstName,ShoppingList,},"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}
func TestGenGraphqlGoFieldsFunc(t *testing.T) {
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
		IsObjectType:    true,
		KeyType:         []string{""},
		ScalarName:      "ShoppingList",
		ScalarNameUpper: "ShoppingList",
		ScalarType:      "string",
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")
	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	buff = GenGraphqlGoFieldsFunc(genObjMap)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(`
	var Fields = graphql.Fields{ 
		"Customer": &graphql.Field{
			Type: customerType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				CustomerObj := Customer{}
				DB.First(&CustomerObj)
				ShoppingListObj := []string{}
				DB.Model(&CustomerObj).Association("ShoppingList").Find(&ShoppingListObj)
				CustomerObj.ShoppingList = append(CustomerObj.ShoppingList, ShoppingListObj...)
				return CustomerObj, nil
			},
		},
}
`)

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}
func TestGenGraphqlGoMainFunc(t *testing.T) {
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
		IsObjectType:    true,
		KeyType:         []string{""},
		ScalarName:      "ShoppingList",
		ScalarNameUpper: "ShoppingList",
		ScalarType:      "string",
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")
	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	GenGraphqlGoMainFunc("test", "testConnString", genObjMap, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(`
var DB *gorm.DB


func main() {

	DB, _ = gorm.Open("test","testConnString")
	defer DB.Close()

	
	fmt.Println("Test with Get	: curl -g 'http://localhost:8080/graphql?query={Customer{FirstName,},}'")
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: Fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	gHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", gHandler)

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
`)

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%v'\n\nReceived\n\n'%v'\n\n", expectedBuff, buff)
		expectedBuffBytes := expectedBuff.Bytes()
		buffBytes := buff.Bytes()
		for i, _ := range expectedBuffBytes {
			if expectedBuffBytes[i] != buffBytes[i] {
				fmt.Println(i)
			}
		}
	}
}
