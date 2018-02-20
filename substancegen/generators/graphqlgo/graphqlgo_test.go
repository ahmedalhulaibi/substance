package graphqlgo

import (
	"testing"
)

func TestGql_GetObjectTypesFunc(t *testing.T) {

}

func TestGql_ResolveRelationshipsFunc(t *testing.T) {

}

func Test_stringInSlice(t *testing.T) {
	type args struct {
		searchVal string
		list      []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Primary UPPER true", args{searchVal: "PRIMARY", list: []string{"PRIMARY", "FOREIGN", "UNIQUE"}}, true},
		{"Foreign UPPER true", args{searchVal: "FOREIGN", list: []string{"PRIMARY", "FOREIGN", "UNIQUE"}}, true},
		{"Unique UPPER true", args{searchVal: "UNIQUE", list: []string{"PRIMARY", "FOREIGN", "UNIQUE"}}, true},
		{"Primary lower true", args{searchVal: "p", list: []string{"p", "f", "u"}}, true},
		{"Foreign lower true", args{searchVal: "f", list: []string{"p", "f", "u"}}, true},
		{"Unique lower true", args{searchVal: "u", list: []string{"p", "f", "u"}}, true},

		{"Primary UPPER false", args{searchVal: "PRIMARY", list: []string{"FOREIGN", "UNIQUE"}}, false},
		{"Foreign UPPER false", args{searchVal: "FOREIGN", list: []string{"PRIMARY", "UNIQUE"}}, false},
		{"Unique UPPER false", args{searchVal: "UNIQUE", list: []string{"PRIMARY", "FOREIGN"}}, false},
		{"Primary lower false", args{searchVal: "p", list: []string{"f", "u"}}, false},
		{"Foreign lower false", args{searchVal: "f", list: []string{"p", "u"}}, false},
		{"Unique lower false", args{searchVal: "u", list: []string{"p", "f"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gql{}
			if got := g.StringInSlice(tt.args.searchVal, tt.args.list); got != tt.want {
				t.Errorf("stringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
