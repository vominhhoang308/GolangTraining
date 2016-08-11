package main

import (
	"fmt"
	"log"
)

// Just like how we create our custom err message inside function Sqrt() in file customFucntionWithErr.go
// now we put the err message into this function fmt.Errorf.
// by this way you can actually print value too

func Sqroot(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("inccorect value... : %v", f) // Please ctrl+cliock the Errorf to see it in more detail when YOU REREAD THIS
	}
	return 42, nil
}

func main() {
	_, err := Sqroot(-10)
	if err != nil {
		log.Fatal(err)
	}
}
