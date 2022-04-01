package main

import "fmt"

func main() {
	//Insert Code here
	var OS [6]string

	OS[0] = "Windows XP"
	OS[1] = "Linux 1.0"
	OS[2] = "Raspbian Teensy"
	OS[3] = "MacOS"
	OS[4] = "IOS"
	OS[5] = "Google Android"

	for i, length := range OS {
		fmt.Printf("%v %v has length of %v\n", i+1, length, len(OS[i]))
	}
}
