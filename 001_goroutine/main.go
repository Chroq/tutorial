package main

import (
	"fmt"
	"strings"
	"time"
)

// shout prints an uppercased name n times, the longer the name the longer the delay
func shout(number int, name string) {
	for j := 1; j <= number; j++ {
		fmt.Println(strings.ToUpper(name))
		time.Sleep(time.Duration(len(name)*100) * time.Millisecond)
	}
}

func main() {
	go shout(3, "Galadrielle")
	go shout(3, "Arwen")
	var a string
	fmt.Scanln(&a)
}
