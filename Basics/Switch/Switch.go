package main

// Flere importes sker i parantecer
import (
	"fmt"
	"time"
)

func main() {
	// basal switch case
	i := 2
	fmt.Println("Write", i, "as: ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two\n")
	case 3:
		fmt.Println("Three")
	}

	// Man kan også bruge commaer til at sepere flere expressions i samme case statement
	// Der er også en "default" som fungere som fallback, hvis ingen statement passer
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // her er der to statements, er det lørdag eller søndag
		fmt.Println("Det er endelig weekend")
	default:
		fmt.Println("Det er hverdag :-( \n")
	}

	// en switch uden en expression er en anden måde at lave if/else logik
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Klokken er før 12\n")
	default:
		fmt.Println("Klokken er efter 12\n")
	}

	// go har noget der hedder en type switch, der sammenligner types, ikke values.
	// et interface{} eller any kan indeholde hvad som helst i go, og derfor kan go ikke
	// finde den rigtige type ved compile time.
	// så en type switch lader dig spørge om den konkrete type ved runtime,
	// og reagere forskelligt alt efter hvad svaret er.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Jeg er en boolean")
		case int:
			fmt.Println("Jeg er en integer")
		default:
			fmt.Printf("Default: kan eventuelt bruge '%%T', til at tjekke: %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(78)
	whatAmI("String check")
}
