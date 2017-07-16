package main

import (
	"fmt"
	"strings"	
)

func main() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("team", ""))
	fmt.Println(strings.ContainsAny("", ""))
}
