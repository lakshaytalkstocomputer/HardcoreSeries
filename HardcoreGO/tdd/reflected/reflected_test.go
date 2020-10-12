package reflected

import "testing"

func TestWalk(t *testing.T){
	t.Run("string", func(t *testing.T) {

		expected := "Chris"
		var got []string

		x := struct{
			Name string
		}{expected}

		Walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1{
			t.Errorf("Wrong number of function calls, got %d want %d", len(got), 1)
		}

		if got[0] != expected{
			t.Errorf("Got %v, expected %q", got[0], expected)
		}
	})
}




