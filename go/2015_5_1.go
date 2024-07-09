package main

import (
	"fmt"
	"strings"
)

func main(){

	nice := 0
	var linha string
	for {
		_ , errorvar := fmt.Scanln(&linha)
		if errorvar != nil{
			break
		}
		
		if is_nice(linha){
			nice++
		}

		fmt.Print(linha,"\n")

	}

	fmt.Print(nice,"\n")

}

func contains_3_vowels(linha string) bool {
	n_vowels := 0
	for i := 0; i < len(linha); i++ {
		if linha[i] == 'a' || linha[i] == 'e' || linha[i] == 'o' || linha[i] == 'i' || linha[i] == 'u' {
			n_vowels++
		}
	}
	if n_vowels >= 3 {
		return true
	}
	return false
}

func contains_2_letters_in_row(linha string) bool{
	for i:= 1; i < len(linha) ; i++ {
		if linha[i] == linha[i-1] {
			return true
		}
	}
	return false
}

func not_abcdpqxy(linha string) bool{
	return !(strings.Contains(linha,"ab") || strings.Contains(linha,"cd") || strings.Contains(linha,"pq") || strings.Contains(linha,"xy"))
}

func is_nice(linha string) bool {

	return contains_3_vowels(linha) && contains_2_letters_in_row(linha) && not_abcdpqxy(linha)
}