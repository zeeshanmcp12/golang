package main

import "fmt"

/*func main() {
	intelligenceQuotient := 180.00
	fmt.Printf("the type of intelligenceQuotient is %T\n", intelligenceQuotient)
}*/

/*
func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}*/

const OvenTime = 40

func main() {
	fmt.Printf("Lasagna preparation time is %v mins\n", OvenTime)
	fmt.Printf("Remaining minutes in oven: %v \n", RemainingOvenTime(30))
	fmt.Printf("Time spent on preparation: %v mins\n", PreparationTime(1))
	fmt.Println("Total time in minutes:", ElapsedTime(1, 30))

	// fmt.Printf("Total Elapsed time in %v mins\n", ElapsedTime(3, 15))

}

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven

}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	// numoflayer := PreparationTime(numOfLayers)
	// timeSpent := RemainingOvenTime(actualMinsInOven)

	// return numoflayer + timeSpent
	// fmt.Printf("Total elapsed time %v\n", PreparationTime(numberOfLayers)+RemainingOvenTime(actualMinutesInOven))
	// result := PreparationTime(numberOfLayers) + RemainingOvenTime(actualMinutesInOven)
	// fmt.Printf("Total elapsed time %v\n", result)

	return (numberOfLayers * 2) + actualMinutesInOven
	// return RemainingOvenTime(actualMinutesInOven) + PreparationTime(numberOfLayers)

}
