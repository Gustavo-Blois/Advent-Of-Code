package main
import (
	"fmt"
)

type Coordinates struct {
	x int 
	y int 
}

func main() {
	visitados := []Coordinates{}
	var linha string
	fmt.Scanln(&linha)
	
	papai_noel := Coordinates{0,0}
	robo_noel := Coordinates{0,0}
	visitados = append(visitados, papai_noel)
	presentes := 1
	for i := 0; i < len(linha); i++ {
		if i % 2 == 0{
			switch linha[i]{
			case '^': papai_noel.y++
			case 'v': papai_noel.y--
			case '<': papai_noel.x--
			case '>': papai_noel.x++
			}
			if nao_visitado(visitados,papai_noel) {
				visitados = append(visitados,papai_noel)
				presentes++
			} 
		} else {
			switch linha[i]{
			case '^': robo_noel.y++
			case 'v': robo_noel.y--
			case '<': robo_noel.x--
			case '>': robo_noel.x++
			}
				if nao_visitado(visitados,robo_noel) {
					visitados = append(visitados,robo_noel)
					presentes++
				}
		}
	}

	fmt.Print(presentes,"\n")

}

func nao_visitado(visitados []Coordinates, atual Coordinates) bool{
	for i := 0; i < len(visitados) ; i++ {
		if (visitados[i].x == atual.x && visitados[i].y == atual.y) {
			return false
		} 
	}
	return true
}