package main

import "fmt"

func main() {
	var linha string
	andar := 0
	fmt.Scanln(&linha)

	for i := 0; i < len(linha); i++ {
		if linha[i] == '(' {
			andar++
		} else {
			andar--
		}
		if andar == -1 {
			print(i + 1)
			print("\n")
			break
		}
	}
	print(andar)
	print("\n")

}
