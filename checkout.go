package main

import (
	"fmt"
	"math"
	"strconv"
)

var interarrivalTimes = [6]int{1, 1, 6, 3, 7, 5} // Rime between each arrival
var serviceTimes = [6]int{4, 2, 5, 4, 1, 5}      // Time between each departure

//var interarrivalTimes = [6]int{4, 5, 2, 8, 3, 7} // Rime between each arrival
//var serviceTimes = [6]int{5, 3, 4, 6, 2, 7}      // Time between each departure
var lq_t = 0 // Number of Customers in Waiting Line
var ls_t = 0 // Number of Customers being served
var s = 0    // Sum of Customer response times by current time
var nD = 0   // Number of Customers who were in the checkout line for >= 5 min
var f = 0    // Number of Departures by current time

var clock = 0     // Current Clock Time
var stopTime = 25 // System Stop Time

var arrivalIndex = 0 // Index for the interarrival times
var serviceIndex = 0 // Index for the service times

type customer struct {
	Index       int
	ArrivalTime int
}

func main() {
	checkoutLine := make([]customer, 0)

	fmt.Println("CLOCK\tLQ(t)\tLS(t)\tCHECKOUT LINE\t\t\tFuture Event List\t\t\tS\tN_D\tF")

	for {
		if clock >= stopTime {
			break
		}

		//If clock is 0, move to first arrival
		if clock == 0 {
			// -1 to adjust for the minute spent initalizing the simulation
			printEvent(checkoutLine)
			c := customer{arrivalIndex, clock}
			clock += interarrivalTimes[arrivalIndex]

			checkoutLine = append(checkoutLine, c)
			arrivalIndex++
			printEvent(checkoutLine)
			continue
		}

		//increment clock to nearest arrival or departure, which is the smallest of the next 2 indexes
		min := int(math.Min(float64(interarrivalTimes[arrivalIndex]), float64(serviceTimes[serviceIndex])))

		//time to arrive
		if min == interarrivalTimes[arrivalIndex] {
			c := customer{arrivalIndex, clock}

			checkoutLine = append(checkoutLine, c)
			arrivalIndex++
		}

		//time to depart
		if min == serviceTimes[serviceIndex] {
			//var c = checkoutLine[0]

			checkoutLine = append(checkoutLine[1:])
			serviceIndex++
		}

		//set the number of customers waiting in line and being served
		if len(checkoutLine) > 0 {
			ls_t = 1
			lq_t = len(checkoutLine) - 1
		} else {
			ls_t = 0
			lq_t = 0
		}

		clock += min

		printEvent(checkoutLine)
		//break

	}
}

func printEvent(checkoutLine []customer) {

	var checkout = ""
	for _, c := range checkoutLine {
		checkout = checkout + "(C" + strconv.Itoa(c.Index+1) + "," + strconv.Itoa(c.ArrivalTime) + ")"
	}

	var nextArrival = "(A," + strconv.Itoa(clock+interarrivalTimes[arrivalIndex]-1) + ",C" + strconv.Itoa(arrivalIndex) + ")"

	fmt.Printf("%d\t%d\t%d\t%s\t\t\t%s\n", clock, lq_t, ls_t, checkout, nextArrival)
}
