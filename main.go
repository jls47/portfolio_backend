package main

import (
	"fmt"
	"os"
	"strconv"
)
//Alright, not sure why the i164 branch in itch isn't being tracked for contributions AT ALL
//One more
//test
func main() {
	s, sep := "", "0:"
	for i, arg := range os.Args[1:] {
		sep = strconv.FormatInt(int64(i), 10) + ": "
		s += sep + arg + " "
	}
	fmt.Println(s)
}
