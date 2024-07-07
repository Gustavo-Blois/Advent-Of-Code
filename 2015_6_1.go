package main

import (
	"fmt"
	"strings"
	"regexp"
	"strconv"
	"os"
	"bufio"
)

func main(){
	var grid_de [1000][1000]int
	grid := &grid_de

	turnon := regexp.MustCompile(`turn on (\d+),(\d+) through (\d+),(\d+)`)
	turnoff := regexp.MustCompile(`turn off (\d+),(\d+) through (\d+),(\d+)`)
	togglevar := regexp.MustCompile(`toggle (\d+),(\d+) through (\d+),(\d+)`)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linha := scanner.Text()
		
		
		if strings.Contains(linha,"turn on") {
			valores := turnon.FindStringSubmatch(linha)
			x1,_ := strconv.Atoi(valores[1])
			y1,_ := strconv.Atoi(valores[2])
			x2,_ := strconv.Atoi(valores[3])
			y2,_ := strconv.Atoi(valores[4])
			turn_on(grid,x1,y1,x2,y2)
		}
		if strings.Contains(linha,"turn off") {
			valores := turnoff.FindStringSubmatch(linha)
			x1,_ := strconv.Atoi(valores[1])
			y1,_ := strconv.Atoi(valores[2])
			x2,_ := strconv.Atoi(valores[3])
			y2,_ := strconv.Atoi(valores[4])
			turn_off(grid,x1,y1,x2,y2)
		}
		if strings.Contains(linha,"toggle") {
			valores := togglevar.FindStringSubmatch(linha)
			x1,_ := strconv.Atoi(valores[1])
			y1,_ := strconv.Atoi(valores[2])
			x2,_ := strconv.Atoi(valores[3])
			y2,_ := strconv.Atoi(valores[4])
			toggle(grid,x1,y1,x2,y2)
		}

	}

	fmt.Print(count_lights(grid),"\n")
	print_leds(grid)

}

func toggle(grid *[1000][1000]int, x1 int, y1 int, x2 int, y2 int){
	for i := x1; i <= x2; i++{
		for k := y1; k <= y2; k++{
			if grid[i][k] == 0{
				grid[i][k] = 1
			} else {
				grid[i][k] = 0
			}
		}
	}
}

func turn_on(grid *[1000][1000]int, x1 int, y1 int, x2 int, y2 int){
	for i := x1; i <= x2; i++{
		for k := y1; k <= y2; k++{
			grid[i][k] = 1
		}
	}
}

func turn_off(grid *[1000][1000]int, x1 int, y1 int, x2 int, y2 int){
	for i := x1; i <= x2; i++{
		for k := y1; k <= y2; k++{
			grid[i][k] = 0
		}
	}
}


func print_leds(grid *[1000][1000]int){
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++{
			if grid[i][j] == 1{
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}


func count_lights(grid *[1000][1000]int) int{
	lights := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++{
			if grid[i][j] == 1{
				lights++
			}
		}
	}
	return lights
}
