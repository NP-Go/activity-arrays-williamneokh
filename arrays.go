package main

import "fmt"

func main() {
	//Activity #1: Declare array variable with a specific size

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
	//Activity #2: Assess a particular element in an array and update the element

	OS[0] = "Windows 10"
	OS[1] = "Linux 16.04"
	OS[2] = "Raspbian Buster"
	fmt.Println("======================")
	for i, length := range OS[:3] {
		fmt.Printf("%v %v has length of %v\n", i+1, length, len(OS[i]))
	}

	var newOS [9]string

	for i := 0; i < len(OS); i++ {
		newOS[i] = OS[i]
	}

	newOS[6] = "Ubuntu"
	newOS[7] = "MS-Dos"
	newOS[8] = "Solaris"
	fmt.Println("======================")
	for i, length := range newOS {
		fmt.Printf("%v %v has length of %v\n", i+1, length, len(newOS[i]))
	}

}
