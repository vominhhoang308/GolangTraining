package main

//In this example, we can see how GO LET US FULLY OPTIMIZE OUR ERR
// so , basically when we dealt with err, we need to have the error interface
// and the error interface have the method Error() string
// So , we can create our own type of struct (class) that can store the error so we can print or do stuff with that error and other field in the struct
// So , we have to implement method Error() string in our newly created type, so any variable of that type will be of type error interface too

import (
	"fmt"
	"log"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

// THIS LINE IS JUST EQUAL TO public class NorgateMatherror implements error {} in java
func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := SqrtCustom(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func SqrtCustom(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		var1 := NorgateMathError{"50.2289 N", "99.4656 W", nme}
		return 0, &var1
	}
	// implementation
	return 42, nil
}

// see use of structs with error type in standard library:
// http://www.goinggo.net/2014/11/error-handling-in-go-part-ii.html
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
