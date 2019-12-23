package day2partA

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day2 PartA")

	values := ReadFile("./Day2/input.txt")

	values[1] = 12
	values[2] = 2
	finalValues := intCode(values)

	fmt.Println(finalValues)
}

func ReadFile(file string) []int {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return toInt(bytes)
}

func toInt(input []byte) []int {
	fields := bytes.Fields(input)

	var array []int
	for _, v := range fields {
		values := strings.Split(string(v), ",")

		for _, v := range values {
			intField, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			array = append(array, intField)
		}
	}

	return array
}

func intCode(input []int) []int {
	var (
		i                    = 0
		posInputA, posInputB int
	)

	for {
		opcode := input[i]

		switch opcode {
		case 1:
			posInputA, posInputB = input[i+1], input[i+2]
			input[input[i+3]] = input[posInputA] + input[posInputB]
		case 2:
			posInputA, posInputB = input[i+1], input[i+2]
			input[input[i+3]] = input[posInputA] * input[posInputB]
		case 99:
			return input
		default:
			log.Fatalf("Cant handle opcode: %v", opcode)
		}

		i += 4
	}

}
