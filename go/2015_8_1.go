package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	actual_chars := 0
	mem_chars := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		actual_chars += len(line)

		for i := 0; i < len(line); i++ {
			if string(line[i]) == `\` {
				if string(line[i+1]) == `x` {
					i += 3
					mem_chars++
				} else {
					i++
					mem_chars++
				}
			} else if string(line[i]) != `"` {
				mem_chars++
			}
		}
	}
	x := actual_chars - mem_chars
	fmt.Print(x, "\n")

}
