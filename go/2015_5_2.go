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

func pairs_dont_overlap(linha string) bool{
	for i := 0; i < len(linha) -1; i++{
		par_atual := linha[i:i+2]
		if strings.Contains(linha[i+2:len(linha)],par_atual) {
			return true
		}	
	}
	return false
}


func repeats_with_one_in_between(linha string) bool{
	for i:= 2; i < len(linha) ; i++ {
		if linha[i] == linha[i-2] {
			return true
		}
	}
	return false
}



func is_nice(linha string) bool {

	return repeats_with_one_in_between(linha) && pairs_dont_overlap(linha)  

}