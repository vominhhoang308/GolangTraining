package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

//This function is to create a BUCKET which the WORD GIVEN will lie into
func HashBucket(word string, numberOfBuckets int) int {
	var sumWord int          // This is the sum of INT32 of each letter in a word (int32 is a rune)
	for _, v := range word { //This loop through each letter in the word, the "_" is suppose to be an index, but we leave it out
		sumWord += int(v) // convert each LETTER into a INT32 type and add that to the SUM
	}

	return sumWord % numberOfBuckets // return the REMAINDER, which is also a BUCKET NUMBER that the word would lie into
}

func main() {
	// Get the text
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt") // this STATEMENT using HTTP GET METHOD to get the content from a URL and parse that into
	// a SLICE of BYTES ("res") and also return the errors if theres any
	if err != nil {
		log.Fatal(err) // handle Error
	}

	// we scan the TEXT and PUT THAT INTO a BUFFER, it would return a SCANNER type
	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close() // CLOSE the scanning thing BEFORE the main exit

	// set the SPLIT FUNCTION to split all the texts into words
	scanner.Split(bufio.ScanWords)

	// Create a slice of (slice of string) to hold slices of words with length 12
	mainBucket := make([][]string, 12)

	//Create slices to hold words
	for i := 0; i < 12; i++ {
		mainBucket = append(mainBucket, []string{}) // we append 12 empty slices of string to a mainBucket
	}

	// this function is too LOOP OVER THE LIST OF WORDS from the text from URL
	for scanner.Scan() {
		word := scanner.Text()                      // this will return the word that is most recently scanned from the Scan function
		n := HashBucket(word, 12)                   // Find the number of BUCKET which the word will lie into, that would be at bucket N
		mainBucket[n] = append(mainBucket[n], word) // bucket N is also an index in the slices of buckets (stored into mainBuckets)
		// then we append the word into that BucketN (at the position N in mainBucket)
		// now BucketN will got a list of words from a text
	}

	//Print the length of each bucket in mainBucket
	for i := 0; i < 12; i++ {
		fmt.Println("Bucket number ", i, " stores ", len(mainBucket[i]), " words.")
	}
}
