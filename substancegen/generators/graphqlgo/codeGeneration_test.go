package graphqlgo

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func TestGenGraphqlGoSampleQueryFunc(t *testing.T) {
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
	var buff bytes.Buffer
	GenGraphqlGoSampleQuery(genObjMap, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(fmt.Sprintf("GetCustomer{FirstName,ShoppingList,},"))

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}
func TestGenGraphqlGoFieldsFunc(t *testing.T) {
	newGenObjType := substancegen.GenObjectType{Name: "Customer", LowerName: "customer", SourceTableName: "Customers"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "FirstName",
		ScalarNameUpper: "FirstName",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
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
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")
	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	var buff bytes.Buffer
	gqlPlugin := Gql{}
	InitGraphqlDataTypes(&gqlPlugin)
	gqlPlugin.PopulateAltScalarType(genObjMap, false, true)
	GenGraphqlGoFieldsFunc(genObjMap, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(`
var QueryFields graphql.Fields

func init() {
	QueryFields = make(graphql.Fields,1)
	
	QueryFields["GetCustomer"] = &graphql.Field{
		Type: customerType,
		Args: graphql.FieldConfigArgument{
			"FirstName": &graphql.ArgumentConfig{
					Type: graphql.String,
			},
			
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryCustomerObj := Customer{}
			if val, ok := p.Args["FirstName"]; ok {
				QueryCustomerObj.FirstName = val.(string)
			}
			
			var ResultCustomerObj Customer
			err := GetCustomer(DB,QueryCustomerObj,&ResultCustomerObj)
			ShoppingListObj := []string{}
			err = append(err,DB.Model(&ResultCustomerObj).Association("ShoppingList").Find(&ShoppingListObj).Error)
			ResultCustomerObj.ShoppingList = append(ResultCustomerObj.ShoppingList, ShoppingListObj...)
			if len(err) > 0 {
				return ResultCustomerObj, err[len(err)-1]
			}
			return ResultCustomerObj, nil
		},
	}

	
	QueryFields["GetAllCustomer"] = &graphql.Field{
		Type: graphql.NewList(customerType),
		Args: graphql.FieldConfigArgument{
			"FirstName": &graphql.ArgumentConfig{
					Type: graphql.String,
			},
			
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryCustomerObj := Customer{}
			if val, ok := p.Args["FirstName"]; ok {
				QueryCustomerObj.FirstName = val.(string)
			}
			
			var ResultCustomerObj []Customer
			err := GetAllCustomer(DB,QueryCustomerObj,&ResultCustomerObj)
			if len(err) > 0 {
				return ResultCustomerObj, err[len(err)-1]
			}
			return ResultCustomerObj, nil
		},
	}

}


`)

	if buff.String() != expectedBuff.String() {
		expectedBuffBytes := expectedBuff.Bytes()
		buffBytes := buff.Bytes()
		for i := range expectedBuffBytes {
			if expectedBuffBytes[i] != buffBytes[i] {
				fmt.Printf("%d Expected Char: %c\tReceived Char: %c\n", i, expectedBuffBytes[i], buffBytes[i])
			}
		}
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
		//t.Errorf("Expected\n\n'%v'\n\nReceived\n\n'%v'\n\n", expectedBuff, buff)
	}
}
func TestGenGraphqlGoTypesFunc(t *testing.T) {
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
		AltScalarType:   make(map[string]string),
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
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")
	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	gqlPlugin := Gql{}
	InitGraphqlDataTypes(&gqlPlugin)
	gqlPlugin.GenerateGraphqlGoTypesFunc(genObjMap, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(`
var customerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Customer",
		Fields: graphql.Fields{ 
			"FirstName":&graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"ShoppingList":&graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(shoppingListType)),
			},
		},
	},
)
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


	fmt.Println("Test with Get	:	curl -g 'http://localhost:8080/graphql?query={ GetCustomer{FirstName,}, }'")

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: QueryFields}
	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: MutationFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(rootMutation)}
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
		for i := range expectedBuffBytes {
			if expectedBuffBytes[i] != buffBytes[i] {
				fmt.Printf("%d Expected Char: %c\tReceived Char: %c\n", i, expectedBuffBytes[i], buffBytes[i])
			}
		}
	}
}

func TestGenGraphqlGoFieldsGetFunc(t *testing.T) {
	newGenObjType := substancegen.GenObjectType{Name: "Customer", LowerName: "customer", SourceTableName: "Customers"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "FirstName",
		ScalarNameUpper: "FirstName",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["FirstName"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["FirstName"].Tags["json"] = append(newGenObjType.Properties["FirstName"].Tags["json"], "firstName")

	newGenObjType.Properties["PhoneNumber"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "PhoneNumber",
		ScalarNameUpper: "PhoneNumber",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["PhoneNumber"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["PhoneNumber"].Tags["json"] = append(newGenObjType.Properties["PhoneNumber"].Tags["json"], "phoneNumber")

	newGenObjType.Properties["ShoppingList"] = &substancegen.GenObjectProperty{
		IsList:          true,
		IsObjectType:    true,
		KeyType:         []string{""},
		ScalarName:      "ShoppingList",
		ScalarNameUpper: "ShoppingList",
		ScalarType:      "ShoppingList",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")

	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	var buff bytes.Buffer
	gqlPlugin := Gql{}
	InitGraphqlDataTypes(&gqlPlugin)
	gqlPlugin.PopulateAltScalarType(genObjMap, false, true)
	GenGraphqlGoFieldsGetFunc(genObjMap, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(`
	QueryFields["GetCustomer"] = &graphql.Field{
		Type: customerType,
		Args: graphql.FieldConfigArgument{
			"FirstName": &graphql.ArgumentConfig{
					Type: graphql.String,
			},
			"PhoneNumber": &graphql.ArgumentConfig{
					Type: graphql.String,
			},
			
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryCustomerObj := Customer{}
			if val, ok := p.Args["FirstName"]; ok {
				QueryCustomerObj.FirstName = val.(string)
			}
			if val, ok := p.Args["PhoneNumber"]; ok {
				QueryCustomerObj.PhoneNumber = val.(string)
			}
			
			var ResultCustomerObj Customer
			err := GetCustomer(DB,QueryCustomerObj,&ResultCustomerObj)
			ShoppingListObj := []ShoppingList{}
			err = append(err,DB.Model(&ResultCustomerObj).Association("ShoppingList").Find(&ShoppingListObj).Error)
			ResultCustomerObj.ShoppingList = append(ResultCustomerObj.ShoppingList, ShoppingListObj...)
			if len(err) > 0 {
				return ResultCustomerObj, err[len(err)-1]
			}
			return ResultCustomerObj, nil
		},
	}
`)

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}

func TestGenGraphqlGoFieldsCreateFunc(t *testing.T) {
	newGenObjType := substancegen.GenObjectType{Name: "Customer", LowerName: "customer", SourceTableName: "Customers"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "FirstName",
		ScalarNameUpper: "FirstName",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["FirstName"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["FirstName"].Tags["json"] = append(newGenObjType.Properties["FirstName"].Tags["json"], "firstName")

	newGenObjType.Properties["PhoneNumber"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "PhoneNumber",
		ScalarNameUpper: "PhoneNumber",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["PhoneNumber"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["PhoneNumber"].Tags["json"] = append(newGenObjType.Properties["PhoneNumber"].Tags["json"], "phoneNumber")

	newGenObjType.Properties["ShoppingList"] = &substancegen.GenObjectProperty{
		IsList:          true,
		IsObjectType:    true,
		KeyType:         []string{""},
		ScalarName:      "ShoppingList",
		ScalarNameUpper: "ShoppingList",
		ScalarType:      "ShoppingList",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")

	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	var buff bytes.Buffer
	gqlPlugin := Gql{}
	InitGraphqlDataTypes(&gqlPlugin)
	gqlPlugin.PopulateAltScalarType(genObjMap, false, false)
	GenGraphqlGoFieldsCreateFunc(genObjMap, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(`
	MutationFields["CreateCustomer"] = &graphql.Field{
		Type: customerType,
		Args: graphql.FieldConfigArgument{
			"FirstName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
			},
			"PhoneNumber": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
			},
			
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryCustomerObj := Customer{}
			if val, ok := p.Args["FirstName"]; ok {
				QueryCustomerObj.FirstName = val.(string)
			}
			if val, ok := p.Args["PhoneNumber"]; ok {
				QueryCustomerObj.PhoneNumber = val.(string)
			}
			
			err := CreateCustomer(DB,QueryCustomerObj)
			var ResultCustomerObj Customer
			err = append(err,GetCustomer(DB,QueryCustomerObj,&ResultCustomerObj)...)
			if len(err) > 0 {
				return ResultCustomerObj, err[len(err)-1]
			}
			return ResultCustomerObj, nil
		},
	}
`)

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
		t.Errorf("Expected\n\n'%v'\n\nReceived\n\n'%v'\n\n", expectedBuff, buff)
		expectedBuffBytes := expectedBuff.Bytes()
		buffBytes := buff.Bytes()
		for i := range expectedBuffBytes {
			if expectedBuffBytes[i] != buffBytes[i] {
				fmt.Printf("%d Expected Char: %c\tReceived Char: %c\n", i, expectedBuffBytes[i], buffBytes[i])
			}
		}
	}
}

func TestGenGraphqlGoFieldsDeleteFunc(t *testing.T) {
	newGenObjType := substancegen.GenObjectType{Name: "Customer", LowerName: "customer", SourceTableName: "Customers"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "FirstName",
		ScalarNameUpper: "FirstName",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["FirstName"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["FirstName"].Tags["json"] = append(newGenObjType.Properties["FirstName"].Tags["json"], "firstName")

	newGenObjType.Properties["PhoneNumber"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "PhoneNumber",
		ScalarNameUpper: "PhoneNumber",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["PhoneNumber"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["PhoneNumber"].Tags["json"] = append(newGenObjType.Properties["PhoneNumber"].Tags["json"], "phoneNumber")

	newGenObjType.Properties["ShoppingList"] = &substancegen.GenObjectProperty{
		IsList:          true,
		IsObjectType:    true,
		KeyType:         []string{""},
		ScalarName:      "ShoppingList",
		ScalarNameUpper: "ShoppingList",
		ScalarType:      "ShoppingList",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")

	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	var buff bytes.Buffer
	gqlPlugin := Gql{}
	InitGraphqlDataTypes(&gqlPlugin)
	gqlPlugin.PopulateAltScalarType(genObjMap, false, false)
	GenGraphqlGoFieldsDeleteFunc(genObjMap, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(`
	MutationFields["DeleteCustomer"] = &graphql.Field{
		Type: customerType,
		Args: graphql.FieldConfigArgument{
			"FirstName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
			},
			"PhoneNumber": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
			},
			
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryCustomerObj := Customer{}
			if val, ok := p.Args["FirstName"]; ok {
				QueryCustomerObj.FirstName = val.(string)
			}
			if val, ok := p.Args["PhoneNumber"]; ok {
				QueryCustomerObj.PhoneNumber = val.(string)
			}
			
			err := DeleteCustomer(DB,QueryCustomerObj)
			var ResultCustomerObj Customer
			err = append(err,GetCustomer(DB,QueryCustomerObj,&ResultCustomerObj)...)
			if len(err) > 0 {
				return ResultCustomerObj, err[len(err)-1]
			}
			return ResultCustomerObj, nil
		},
	}
`)

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
		t.Errorf("Expected\n\n'%v'\n\nReceived\n\n'%v'\n\n", expectedBuff, buff)
		expectedBuffBytes := expectedBuff.Bytes()
		buffBytes := buff.Bytes()
		for i := range expectedBuffBytes {
			if expectedBuffBytes[i] != buffBytes[i] {
				fmt.Printf("%d Expected Char: %c\tReceived Char: %c\n", i, expectedBuffBytes[i], buffBytes[i])
			}
		}
	}
}

func TestGenGraphqlGoFieldsUpdateFunc(t *testing.T) {
	newGenObjType := substancegen.GenObjectType{Name: "Customer", LowerName: "customer", SourceTableName: "Customers"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "FirstName",
		ScalarNameUpper: "FirstName",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["FirstName"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["FirstName"].Tags["json"] = append(newGenObjType.Properties["FirstName"].Tags["json"], "firstName")

	newGenObjType.Properties["PhoneNumber"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "PhoneNumber",
		ScalarNameUpper: "PhoneNumber",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["PhoneNumber"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["PhoneNumber"].Tags["json"] = append(newGenObjType.Properties["PhoneNumber"].Tags["json"], "phoneNumber")

	newGenObjType.Properties["ShoppingList"] = &substancegen.GenObjectProperty{
		IsList:          true,
		IsObjectType:    true,
		KeyType:         []string{""},
		ScalarName:      "ShoppingList",
		ScalarNameUpper: "ShoppingList",
		ScalarType:      "ShoppingList",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")

	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	var buff bytes.Buffer
	gqlPlugin := Gql{}
	InitGraphqlDataTypes(&gqlPlugin)
	gqlPlugin.PopulateAltScalarType(genObjMap, false, false)
	GenGraphqlGoFieldsUpdateFunc(genObjMap, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(`
	MutationFields["UpdateCustomer"] = &graphql.Field{
		Type: customerType,
		Args: graphql.FieldConfigArgument{
			"FirstName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
			},
			"PhoneNumber": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
			},
			
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			OldCustomerObj := Customer{}
			QueryCustomerObj := Customer{}
			if val, ok := p.Args["FirstName"]; ok {
				OldCustomerObj.FirstName = val.(string)
				QueryCustomerObj.FirstName = val.(string)
			}
			if val, ok := p.Args["PhoneNumber"]; ok {
				
				QueryCustomerObj.PhoneNumber = val.(string)
			}
			
			var ResultCustomerObj Customer
			err := UpdateCustomer(DB,OldCustomerObj,QueryCustomerObj,&ResultCustomerObj)
			if len(err) > 0 {
				return ResultCustomerObj, err[len(err)-1]
			}
			return ResultCustomerObj, nil
		},
	}
`)

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
		t.Errorf("Expected\n\n'%v'\n\nReceived\n\n'%v'\n\n", expectedBuff, buff)
		expectedBuffBytes := expectedBuff.Bytes()
		buffBytes := buff.Bytes()
		for i := range expectedBuffBytes {
			if expectedBuffBytes[i] != buffBytes[i] {
				fmt.Printf("%d Expected Char: %c\tReceived Char: %c\n", i, expectedBuffBytes[i], buffBytes[i])
			}
		}
	}
}

func TestGenGraphqlGoFieldsGetAllFunc(t *testing.T) {
	newGenObjType := substancegen.GenObjectType{Name: "Customer", LowerName: "customer", SourceTableName: "Customers"}
	newGenObjType.Properties = make(substancegen.GenObjectProperties)
	newGenObjType.Properties["FirstName"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "FirstName",
		ScalarNameUpper: "FirstName",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["FirstName"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["FirstName"].Tags["json"] = append(newGenObjType.Properties["FirstName"].Tags["json"], "firstName")

	newGenObjType.Properties["PhoneNumber"] = &substancegen.GenObjectProperty{
		IsList:          false,
		IsObjectType:    false,
		KeyType:         []string{"PRIMARY KEY"},
		ScalarName:      "PhoneNumber",
		ScalarNameUpper: "PhoneNumber",
		ScalarType:      "string",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["PhoneNumber"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["PhoneNumber"].Tags["json"] = append(newGenObjType.Properties["PhoneNumber"].Tags["json"], "phoneNumber")

	newGenObjType.Properties["ShoppingList"] = &substancegen.GenObjectProperty{
		IsList:          true,
		IsObjectType:    true,
		KeyType:         []string{""},
		ScalarName:      "ShoppingList",
		ScalarNameUpper: "ShoppingList",
		ScalarType:      "ShoppingList",
		AltScalarType:   make(map[string]string),
		Nullable:        false,
	}
	newGenObjType.Properties["ShoppingList"].Tags = make(substancegen.GenObjectTag)
	newGenObjType.Properties["ShoppingList"].Tags["json"] = append(newGenObjType.Properties["ShoppingList"].Tags["json"], "shoppingList")

	genObjMap := make(map[string]substancegen.GenObjectType)
	genObjMap["Customers"] = newGenObjType
	var buff bytes.Buffer
	gqlPlugin := Gql{}
	InitGraphqlDataTypes(&gqlPlugin)
	gqlPlugin.PopulateAltScalarType(genObjMap, false, true)
	GenGraphqlGoFieldsGetAllFunc(genObjMap, &buff)

	var expectedBuff bytes.Buffer

	expectedBuff.WriteString(`
	QueryFields["GetAllCustomer"] = &graphql.Field{
		Type: graphql.NewList(customerType),
		Args: graphql.FieldConfigArgument{
			"FirstName": &graphql.ArgumentConfig{
					Type: graphql.String,
			},
			"PhoneNumber": &graphql.ArgumentConfig{
					Type: graphql.String,
			},
			
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryCustomerObj := Customer{}
			if val, ok := p.Args["FirstName"]; ok {
				QueryCustomerObj.FirstName = val.(string)
			}
			if val, ok := p.Args["PhoneNumber"]; ok {
				QueryCustomerObj.PhoneNumber = val.(string)
			}
			
			var ResultCustomerObj []Customer
			err := GetAllCustomer(DB,QueryCustomerObj,&ResultCustomerObj)
			if len(err) > 0 {
				return ResultCustomerObj, err[len(err)-1]
			}
			return ResultCustomerObj, nil
		},
	}
`)

	if buff.String() != expectedBuff.String() {
		t.Errorf("Expected\n\n'%s'\n\nReceived\n\n'%s'\n\n", expectedBuff.String(), buff.String())
	}
}
