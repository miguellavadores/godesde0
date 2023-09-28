package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)
	
	paises["Mexico"] = "CDMX"
	paises["Argentina"] = "Buenos Aires"
	paises["Peru"] = "Lima"
	//fmt.Println(paises)
	//fmt.Println(paises["Peru"])
	
	campeonato := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"Boca Juniors": 30,
	}
	
	fmt.Println(campeonato)
	
	/*for equipo, puntaje := range campeonato {
		fmt.Printf("equipo %s, tiene un puntaje de %d\n", equipo, puntaje)
	}*/
	
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)
	
	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t\n", puntaje, existe)
}