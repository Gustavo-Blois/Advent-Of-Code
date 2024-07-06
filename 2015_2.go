package main

import (
	"math"
	"fmt"
	"sort"
)

func main() {
	var l int
	var w int
	var h int
	i := 0
	surface := 0
	ribbontotal := 0
	for {
		_, errorvar := fmt.Scanf("%dx%dx%d\n", &l, &w, &h)
		if errorvar != nil {
			break
		}

		var arrayOfRibbon = []int{l,w,h}
		sort.Ints(arrayOfRibbon) 
		ribbon := 2*arrayOfRibbon[0] + 2* arrayOfRibbon[1] 
		bow := l*w*h
		ribbontotal += bow + ribbon

		a1 := l*w
		a2 := w*h
		a3 := l*h
		extra := int(math.Min(float64(a1), math.Min(float64(a2),float64(a3)))) 
		surface += 2*a1 + 2*a2 + 2*a3 + extra
		i++
	}
	fmt.Print(surface,"\n")
	fmt.Print(ribbontotal,"\n")
}
