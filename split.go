package main

import (
	"fmt"	
	"strings"	
)

func main() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a cancal anama", "a "))
	fmt.Printf("%q\n", strings.Split("", "Bernard 0'Hiisnfs"))
}
