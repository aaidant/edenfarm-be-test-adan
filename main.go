package main

import (
	"fmt"

	"edenfarm.com/test-edennfarm-adan/service"
)

var (
	s = service.ConsecutiveService{}
)

func main() {
	var str = ""
	fmt.Print("Input arrays line up : ")
	fmt.Scanln(&str)

	largestValue, err := s.FindLargestValue(str)

	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	fmt.Println("Largest value of an array of consecutive values : ", largestValue)
}
