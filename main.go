package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var Total int
	enable := true
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("no data file found")
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {

		for i := 0; i < len(line)-3; i++ {
			if enable {
				if i+7 < len(line) {
					if line[i:i+7] == "don't()" {
						enable = false
					}
				}
			} else {
				if i+4 < len(line) {
					if line[i:i+4] == "do()" {
						enable = true
					}
				}
			}

			if line[i:i+3] == "mul" {
				if string(line[i+3]) == "(" {
					inputs := strings.Split(line[i+4:], ")")
					if strings.Contains(inputs[0], ",") {
						//fmt.Println(inputs[0])
						if !strings.Contains(inputs[0], " ") {
							operands := strings.Split(inputs[0], ",")
							if len(operands) == 2 {
								operand1, err1 := strconv.Atoi(operands[0])
								if err1 != nil {
									continue
									//log.Fatal("operand1 error: ", err1)
								}
								operand2, err2 := strconv.Atoi(operands[1])
								if err2 != nil {
									continue
									//log.Fatal("operand2 error: ", err2)
								}
								if enable {
									Total += operand1 * operand2
								}
							}
						}
					}
					continue
					//fmt.Println(operands)
				}
			} //*/
		}
	}
	fmt.Println("total muls in data:", Total)
}
