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
		}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls){
				t.Errorf("got %v, eamr %v,",got, test.ExpectedCalls)
			}

		})
}

	}