package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}
