package reflected

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T){

	t.Run("normal data-types", func(t *testing.T) {
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
			{
				"Arrays",
				[2]struct{
					Age int
					City string
				}{
					{
						33, "Chand",
					},
					{
						44, "Banglore",
					},
				},
				[]string{"Chand","Banglore"},
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

	})

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got[] string
		Walk(aMap, func(input string){
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string){
	t.Helper()

	contains := false
	for _, x :=  range haystack{
		if x == needle {
			contains = true
		}
	}

	if !contains{
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}

}