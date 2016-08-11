package main

import (
	"fmt"
	"log"
	"os"
)

// Error handling in GO is treated diffrently compare to any other programming languages
// It treated without the trycatch blocks, which will makes the program looks more redundancy.

// function init is to initialize the program before main, it run once, we can have multiple init() in nay package or files.
func init() {
	newCreatedFiles, err := os.Create("logErr.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(newCreatedFiles) // THIS LINE WILL WRITE THE ERR INTO the file that is newly created instead of stdout.
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		//log.Println("err happened", err) // this will do the same with fmt.Println(), however, it will
		// also print out the log time.

		log.Fatalln(err) // this will give us the log message of the error, just like the Println()
		// however, it will also run os.Exit() after that, os.Exit() is to indicate which
		// go routines status has been terminate by the err, 0 for non-err, and other number for theres-are-some-errs

		//panic(err) // this will give you the log message AND THE STACK OF HOW THE MESSAGE BEEN CREATED, including the os.Exit()
		// we can fix it by using recover() after that, like try-catch, but only for important matters.
		// not recommended this way.
	}
}
