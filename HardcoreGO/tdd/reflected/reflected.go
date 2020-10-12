package reflected


func Walk(x interface{}, fn func(input string) ){
	fn("Some Text")
}
