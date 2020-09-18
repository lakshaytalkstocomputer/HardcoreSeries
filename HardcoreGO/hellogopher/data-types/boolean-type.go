package main


import "fmt"

var LightSwitchIsOn bool // Zero-Value of bool is false
func main(){

	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()

}

func printLightSwitchState(){
	fmt.Println("Light Switch is off: ", LightSwitchIsOn)
}

func toggleLightSwitch(){
	LightSwitchIsOn = !LightSwitchIsOn
}