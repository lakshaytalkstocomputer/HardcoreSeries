package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T){
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

// * The following function will not be executed if I remove the `comment`:Output :6
// * The function will get compiled, but will not be executed.
// * This type of function helps to build documentation.
// * This will appear in documentation inside godoc
func ExampleAdd(){
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}