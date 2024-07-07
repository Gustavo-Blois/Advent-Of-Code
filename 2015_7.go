package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isNumber(str string) (int, bool) {
	num, err := strconv.Atoi(str)
	return num, err == nil
}

func evaluateWire(wire string, wiremap map[string]int, text string) int {
	if val, ok := wiremap[wire]; ok {
		return val
	}

	regex := regexp.MustCompile(`(?:(NOT) (\w+)|(\w+) (AND|OR|LSHIFT|RSHIFT) (\w+)|(\w+)) -> (\w+)`)
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)

		if match[7] == wire {
			if match[1] == "NOT" {
				// Caso NOT
				val := resolveOperand(match[2], wiremap, text)
				wiremap[wire] = ^val
			} else if match[3] != "" {

				val1 := resolveOperand(match[3], wiremap, text)
				val2 := resolveOperand(match[5], wiremap, text)

				switch match[4] {
				case "AND":
					wiremap[wire] = val1 & val2
				case "OR":
					wiremap[wire] = val1 | val2
				case "LSHIFT":
					wiremap[wire] = val1 << val2
				case "RSHIFT":
					wiremap[wire] = val1 >> val2
				}
			} else if match[6] != "" {
				// Caso de atribuição direta
				val := resolveOperand(match[6], wiremap, text)
				wiremap[wire] = val
			}
			break
		}
	}

	return wiremap[wire]
}

func resolveOperand(operand string, wiremap map[string]int, text string) int {
	num, isNum := isNumber(operand)
	if isNum {
		return num
	}
	return evaluateWire(operand, wiremap, text)
}

func main() {
	file, err := os.ReadFile("in_2015_7.txt")
	if err != nil {
		panic(err)
	}
	text := string(file)

	wiremap := map[string]int{}
	wiremap["b"] = 3176
	fmt.Println(evaluateWire("a", wiremap, text))
}
