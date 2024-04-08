package main

import "fmt"

func SlurredTalk(pwords *string) {
	// TODO: answer here
	result := ""
	for _, c := range *pwords {
		if c == 'S' || c == 'R' || c == 'Z' {
			result += "L"
		} else if c == 's' || c == 'r' || c == 'z' {
			result += "l"
		} else {
			result += string(c)
		}
	}
	fmt.Println(result)
	*pwords = result
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	var pwords *string = &words
	// fmt.Println("Before", words)
	SlurredTalk(pwords)
	// fmt.Println("After", words)
}
