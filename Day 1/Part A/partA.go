package day1partA

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func Run()  {
	input,err := ioutil.ReadFile("./Day 1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fields := bytes.Fields(input)
	var totalRequiredFuel int
	for _,v  := range fields {

		mass,err  := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}

		fuelReq := (mass/3) - 2

		totalRequiredFuel += fuelReq
	}

	fmt.Println("Total fuel required: ", totalRequiredFuel)
}
