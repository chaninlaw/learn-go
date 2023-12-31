package control

import (
	"fmt"
)

func Loop() {
	i := 1
	for {
		if i > 5 {
			break
		}
		fmt.Printf("number: %d\n", i)
		i++
	}
}
