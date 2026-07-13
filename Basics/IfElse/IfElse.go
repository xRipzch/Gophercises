package main

import "fmt"

func main() {
	// Basalt eksempel, odd eller even?
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd\n")
	}
	// Man kan have if statements uden else
	if 8%4 == 0 {
		fmt.Println("8 er delbart med 4\n")
	}
	// Logiske operatore som &&(og), ||(eller) kan også bruges
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Enten er 8, eller 7 et ligetal\n")
	}
	// I go kan man lave en statement inden i selve blokken, den lever kun her
	// og bruges kun en gang, før if statement blokken reelt starter
	// num := 9 er den statement, der kører 1 gang før if-tjekket.
	// num er tilgængelig i alle grene (if/else if/else), men findes kun inde i denne blok.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// I go er der ingen ternary if
// Python eksempel:
// result = x if a > b else y
