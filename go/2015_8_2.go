package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	chars := 0
	post_chars := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		post_chars += len(line)

		chars += 2
		for i := 0; i < len(line); i++ {
			if string(line[i]) == `"` || string(line[i]) == `\` {
				chars += 2
			} else {
				chars++
			}
		}

	}
	x := chars - post_chars
	fmt.Print(x, "\n")

}
