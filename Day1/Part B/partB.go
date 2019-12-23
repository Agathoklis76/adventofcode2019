package day1partB

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func Run() {
	input, err := ioutil.ReadFile("./Day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fields := bytes.Fields(input)
	var totalRequiredFuel int
	for _, v := range fields {

		mass, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}

		fuelReq := (mass / 3) - 2

		fuelReq += ExtraFuel(fuelReq)

		totalRequiredFuel += fuelReq
	}

	fmt.Println("Total fuel required: ", totalRequiredFuel)
}

func ExtraFuel(fuel int) int {
	var totalRequiredFuel int
	for fuelToAdd := fuel/3 - 2; fuelToAdd > 0; fuelToAdd = fuelToAdd/3 - 2 {
		totalRequiredFuel += fuelToAdd
	}

	return totalRequiredFuel
}
