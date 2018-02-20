package graphqlgo

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/ahmedalhulaibi/substance/substancegen"
)

func TestGql_AddJSONTagsToProperties(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectTypes map[string]substancegen.GenObjectType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.AddJSONTagsToProperties(tt.args.gqlObjectTypes)
		})
	}
}

func TestGql_GenPackageImports(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		dbType string
		buff   *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenPackageImports(tt.args.dbType, tt.args.buff)
		})
	}
}

func TestGql_GenObjectTypeToStringFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectType substancegen.GenObjectType
		buff          *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenObjectTypeToStringFunc(tt.args.gqlObjectType, tt.args.buff)
		})
	}
}

func TestGql_GenObjectPropertyToStringFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectProperty substancegen.GenObjectProperty
		buff              *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenObjectPropertyToStringFunc(tt.args.gqlObjectProperty, tt.args.buff)
		})
	}
}

func TestGql_GenObjectTagToStringFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		genObjectTags substancegen.GenObjectTag
		buff          *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenObjectTagToStringFunc(tt.args.genObjectTags, tt.args.buff)
		})
	}
}

func TestGql_GenGormObjectTableNameOverrideFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectType substancegen.GenObjectType
		buff          *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenGormObjectTableNameOverrideFunc(tt.args.gqlObjectType, tt.args.buff)
		})
	}
}

func TestGql_GenGraphqlGoTypeFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectType substancegen.GenObjectType
		buff          *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenGraphqlGoTypeFunc(tt.args.gqlObjectType, tt.args.buff)
		})
	}
}

func TestGql_GenGraphqlGoTypePropertyFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectProperty substancegen.GenObjectProperty
		buff              *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenGraphqlGoTypePropertyFunc(tt.args.gqlObjectProperty, tt.args.buff)
		})
	}
}

func TestGql_ResolveGraphqlGoFieldType(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectProperty substancegen.GenObjectProperty
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			if got := g.ResolveGraphqlGoFieldType(tt.args.gqlObjectProperty); got != tt.want {
				t.Errorf("Gql.ResolveGraphqlGoFieldType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGql_GenGraphqlGoMainFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		dbType           string
		connectionString string
		gqlObjectTypes   map[string]substancegen.GenObjectType
		buff             *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenGraphqlGoMainFunc(tt.args.dbType, tt.args.connectionString, tt.args.gqlObjectTypes, tt.args.buff)
		})
	}
}

func TestGql_GenGraphqlGoQueryFieldsFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectType substancegen.GenObjectType
		buff          *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenGraphqlGoQueryFieldsFunc(tt.args.gqlObjectType, tt.args.buff)
		})
	}
}

func TestGql_GenGraphqlGoSampleQuery(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectTypes map[string]substancegen.GenObjectType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bytes.Buffer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			if got := g.GenGraphqlGoSampleQuery(tt.args.gqlObjectTypes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gql.GenGraphqlGoSampleQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGql_GenGraphlGoSampleObjectQuery(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectTypes map[string]substancegen.GenObjectType
		gqlObjectType  substancegen.GenObjectType
		buff           *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenGraphlGoSampleObjectQuery(tt.args.gqlObjectTypes, tt.args.gqlObjectType, tt.args.buff)
		})
	}
}

func TestGql_OutputGraphqlSchema(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectTypes map[string]substancegen.GenObjectType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bytes.Buffer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			if got := g.OutputGraphqlSchema(tt.args.gqlObjectTypes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gql.OutputGraphqlSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGql_OutputCodeFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		dbType           string
		connectionString string
		gqlObjectTypes   map[string]substancegen.GenObjectType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bytes.Buffer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			if got := g.OutputCodeFunc(tt.args.dbType, tt.args.connectionString, tt.args.gqlObjectTypes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gql.OutputCodeFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGql_GenGraphqlGoRootQueryFunc(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectTypes map[string]substancegen.GenObjectType
		buff           *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenGraphqlGoRootQueryFunc(tt.args.gqlObjectTypes, tt.args.buff)
		})
	}
}

func TestGql_GenObjectGormCrud(t *testing.T) {
	type fields struct {
		Name                  string
		GraphqlDataTypes      map[string]string
		GraphqlDbTypeGormFlag map[string]bool
		GraphqlDbTypeImports  map[string]string
	}
	type args struct {
		gqlObjectType substancegen.GenObjectType
		buff          *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{
				Name:                  tt.fields.Name,
				GraphqlDataTypes:      tt.fields.GraphqlDataTypes,
				GraphqlDbTypeGormFlag: tt.fields.GraphqlDbTypeGormFlag,
				GraphqlDbTypeImports:  tt.fields.GraphqlDbTypeImports,
			}
			g.GenObjectGormCrud(tt.args.gqlObjectType, tt.args.buff)
		})
	}
}
