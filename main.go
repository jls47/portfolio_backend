package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", "0:"
	for i, arg := range os.Args[1:] {
		sep = strconv.FormatInt(int64(i), 10) + ": "
		s += sep + arg + " "
	}
	fmt.Println(s)
}
