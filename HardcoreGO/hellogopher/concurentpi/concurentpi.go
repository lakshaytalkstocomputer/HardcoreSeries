package concurentpi

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// We will be using Neelkantha's Formula to generate value of PI

/* We want to create a virtual scoreboard where we can simultaneously
 *  show the current value of PI and how many Nilkantha terms we have
 *  calculated.
 *
 * We can create virtual score board by using ANSI escape codes
 *  Here's wiki article giving us the complete list of ANSI escape codes
 *   http://en.wikipedia.org/wiki/ANSI_escape_code
 */

const ANSIClearScreenSequence = "\033[H\033[2J"
const ANSIFirstSlotScreenSequence = "\033[2;0H"
const ANSISecondSLotScreenSequence = "\033[3;0H"

// The Channel used to update the current value of pi on scoreboard
var pichan chan float64 = make(chan float64)

// THe channel that we use to indicate that the program can exit
var computationDone chan bool = make(chan bool, 1)

// Number of Nilakantha terms for the scoreboard
var termsCount int

// This function serves as out virtual scoreboard to show the current
//  computed value of PI using Nilankantha's formula
func printCalculationSummary(){
	fmt.Print(ANSIClearScreenSequence)
	fmt.Println(ANSIFirstSlotScreenSequence, "Computed Value of PI:\t\t", <-pichan)
	fmt.Println(ANSISecondSLotScreenSequence, "# of Nilakantha Terms:\t\t", termsCount)
}

func pi(n int) float64{
	ch := make(chan float64)

	// The k variable is considered to be the current step
	for k:=1; k<=n; k++ {
		go nilakanthaTerm(ch, float64(k))
	}

	f := 3.0

	for k:= 1; k<=n ; k++{
		termsCount++
		f += <-ch
		pichan <- f
	}

	computationDone <- true
	return f
}

func nilakanthaTerm(ch chan float64, k float64){
	j := 2*k

	if int64(k)%2 == 1{
		ch <- 4.0 / (j * (j + 1) * (j + 2))
	}else {
		ch <- -4.0 /(j * (j + 1) * (j +2 ))
	}
}

func main(){
	ticker := time.NewTicker(time.Millisecond * 108)

	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt)

	go pi(5000)

	go func(){
		for range ticker.C{
			printCalculationSummary()
		}
	}()

	for {
		select {

		case <- computationDone:
			ticker.Stop()
			fmt.Println("Program done calculating PI.")
			os.Exit(0)
		case <-interrupt:
			ticker.Stop()
			fmt.Println("Program interrupted by the user")
			return
		}
	}

}


