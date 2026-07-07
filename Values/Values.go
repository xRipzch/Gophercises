package main

import "fmt"

func main() {
	fmt.Println("And" + "ers")

	fmt.Println("1+1 =", 1+1) // strings kan også bruges, kan ikke se en usecase tho
	fmt.Println(1 + 1)        // ints kan bare bruges til regne stykket
	fmt.Println("7/3 = ", 7/3)

	fmt.Println(true && false) // Fordi den ene del af && er false, bliver hele udtrykket false.
	fmt.Println(true || false) // returnere true, hvis den ene ||(enten eller) er true
	fmt.Println(!true)         // Returnere false, da !(modsætning) af true
}
