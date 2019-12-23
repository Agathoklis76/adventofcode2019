package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadFile(file string) []int {
	bytess, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return toInt(bytess)
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

func IntCode(input []int) []int {
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
			panic(fmt.Sprintf("Cant handle opcode: %v", opcode))
			//log.Fatalf("Cant handle opcode: %v", opcode)
		}

		i += 4
	}

}
