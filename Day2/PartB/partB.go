package day2partB

import (
	"fmt"
	"github.com/Agathoklis76/adventofcode2019/util"
)

func Run() {
	fmt.Println("Day2 Part B")

	instructionSet := util.ReadFile("./Day2/input.txt")

	bruteForce(instructionSet)
}

func bruteForce(input []int) {

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {

			var inputCopy = make([]int, len(input))
			copy(inputCopy, input)

			inputCopy[1] = x
			inputCopy[2] = y

			var result []int
			func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered ", r)
					}
				}()

				result = util.IntCode(inputCopy)

				if result[0] == 19690720 {
					fmt.Println("X:: ", x, "Y:: ", y)
					fmt.Println("RESULT:: ", (x*100)+y)
					return
				}
			}()
		}
	}

}
