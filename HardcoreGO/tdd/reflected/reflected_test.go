package reflected

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T){

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
				}{"Lakshay"},
			[]string{"Lakshay"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},

	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls){
				t.Errorf("got %v, expected %v,",got, test.ExpectedCalls)
			}

		})
}

	}