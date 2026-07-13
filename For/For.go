// For er et klassisk Initial/condition/after for loop
// for er Go's eneste looping konstrukt.

package main

import "fmt"

func main() {
	// Basalt for loop med en single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("Basic for loop, single condition \n")

	// Klassisk for-løkke: init (j:=0) kører én gang, condition (j<3) tjekkes før hver
	// iteration, og post (j++) kører efter hver iteration.
	for j := 0; j <= 3; j++ {
		fmt.Println(j)
	}
	fmt.Println("Klassisk for-loop: init kører kun en gang, condition tjekkes før hvert loop,")
	fmt.Println("og post kører efter hvert loop \n")

	// En anden metode for at få "do this N times" er range.
	// range 3 genererer tallene 0,1,2.
	// Bruges typisk til simpel "gør noget n gange"
	for i := range 3 {
		fmt.Println("range:", i)
	}
	fmt.Println("range kan bruges til simpel 'do this N times' ")
	fmt.Println("men bruges typsik til iteration over collections (slices, maps, arrays etc..)\n")

	// en for loop uden condition vil loope uendeligt, indtil det rammer et break, eller return
	// Break: breaker ud af loopet, men ikke resten af funktionen
	// Return: stopper hele funktionen og går tilbage til der hvor funktionen blev kaldt.
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("for uden condition kører indtil et break eller en return \n")

	// Continue: springer resten af den nuværende iteration over og fortsætter til næste,
	// så loopet stopper ikke, men den ene iteration der passer med condition afsluttes bare tidligt.
	for j := 0; j < 5; j++ {
		if j == 2 {
			continue // springer resten af krop over for j==2, går videre til j++
		}
		fmt.Println(j)
	}
	fmt.Println("output: 0, 1, 3, 4 (2 bliver sprunget over)")
}
