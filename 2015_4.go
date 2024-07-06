package main

import (
	"fmt"
	"strconv"
	"crypto/md5"
)

func main() {
	chave_inicial := "yzbqklnj"
	i := 0
	for {
		chave := chave_inicial + strconv.Itoa(i)
		hash := md5.Sum([]byte(chave))
		hashstring := fmt.Sprintf("%x",hash)
		if hashstring[:6] == "000000" {
			fmt.Print(strconv.Itoa(i),"\n")
			break
		}

		i++
	}
}