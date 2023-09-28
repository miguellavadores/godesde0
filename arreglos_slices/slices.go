package arreglos_slices

import "fmt"

var tablaS []int = []int{2,5,4}
var arreglo [10]int = [10]int{6,98,21,674,18,36,78,9}
func MuestroSlice(){
	fmt.Println(tablaS)
	
	porcion := arreglo[3:]
	porcion2 := arreglo[:5]
	porcion3 := arreglo[6:7]
	
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("La capacidad es %d y el tamaño es %d\n", len(elementos), cap(elementos))
	
	nums := make([]int, 0, 0)
	for i:=0; i<100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLa capacidad es %d y el tamaño es %d\n", len(nums), cap(nums))
}