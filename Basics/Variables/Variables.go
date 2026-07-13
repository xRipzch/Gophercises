// I go er variabler explicit declared, og brugt af compileren.
// var declars 1 eller flere variabler, go vil selv konkludere typen
// := syntaxen er en kort og smart måde at initalisere en variable.
// 			Dog er denne syntax kun tilgængelig i funktioner, ikke til en global variable.

package main

import "fmt"

func main() {
	a := "fodbold"
	fmt.Println(a)

	var b, c int = 1, 2 // Deklare flere variable, hvis samme type
	fmt.Println("Flere variabler i en var ", b, c)

	d := true
	fmt.Println(d)

	var e int      // Variabler uden startværdi får automatisk deres "zero-value" (f.eks. 0 for int).
	fmt.Println(e) // vil printe 0

	f := "apple"
	fmt.Println(f)
}
