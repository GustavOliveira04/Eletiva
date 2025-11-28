package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("é fim de semana")
	default:
		fmt.Println("é dia de semana")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("é antes de meio-dia")
	default:
		fmt.Println("é depois de meio-dia")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("sou um booleano")
		case int:
			fmt.Println("sou um inteiro")
		default:
			fmt.Printf("não sei o tipo de %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("oi")
}
