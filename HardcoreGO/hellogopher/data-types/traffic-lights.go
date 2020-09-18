package main


import "fmt"

const (
	// "_" is called blank identifier, it is used to ignore the value
	_ = iota
	TrafficLightStateRedLight
	TrafficLightStateYellowLight
	TrafficLightStateGreenLight
)


func main(){
	fmt.Println("Red Light State Code: ",TrafficLightStateRedLight)
	fmt.Println("Yellow Light State Code: ",TrafficLightStateYellowLight)
	fmt.Println("Green Light State Code: ",TrafficLightStateGreenLight)
}