package main

import (
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)

// This is without the blank identifier
//func main() {
//	res, err := http.Get("http://www.mcleods.com/")
//	if err != nil {
//		log.Fatal(err)
//	}
//	page, err := ioutil.ReadAll(res.Body)
//	res.Body.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%s", page)
//}

//This is with the blank identifier
func main() {
	res, _ := http.Get("http://www.mcleods.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}