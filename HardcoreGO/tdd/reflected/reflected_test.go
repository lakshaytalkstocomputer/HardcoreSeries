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
		{
			"Strruct with non string field",
			struct {
				Name string
				Age int
			}{"lakshay", 24},
			[]string{"lakshay"},
		},
		{
			"Struct with another struct within - Nested Structs",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{
					"Lakshay",
					struct {
						Age	int
						City string
					}{
						33,
						"London",
					},
			},
			[]string{"Lakshay", "London"},
		},
		{
			"Pointer to Struct with another struct within - Nested Structs",
			&struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{
				"Lakshay",
				struct {
					Age	int
					City string
				}{
					33,
					"London",
				},
			},
			[]string{"Lakshay", "London"},
		},
		{
			"Slices",
			[]struct {
					Age  int
					City string
				}{
					{
						33,
						"London",
					},
					{
						44,
						"Chandigarh",
					},
			},
			[]string{"London", "Chandigarh"},
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