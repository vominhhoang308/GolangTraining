package main

import "fmt"

func main() {
	data := []float64{40, 60, 100, 200, 300}
	n := average(data...) // the variadics (the ... notation) in here comes AFTER the variables.
	// that means, it takes EACH value IN THE LIST ONE AT THE TIME,
	//and PASS IT TO THE PARAMS. Different from the variadics below
	fmt.Println(n)
}

func average(sf ...float64) float64 { //This variadic comes BEFORE the type, that mean we have to pass
	//ZERO or MULTIPLE float64 variables
	//In the main function, we called a "data", which is a list, so it is not valid
	//however, we use variadic after "data", so it pass each value.
	// average(data...) will function just like average (40, 60, 100, 200, 300)
	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
