package main

import "fmt"

func main() {
	data := []float64{40, 50, 60, 1000}
	n := average(4, 3)
	fmt.Println(n)
}

func average(sf []float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

// THIS IS THE SAME AS TWO PREVIOUS FILE; JUST ALTENATION
