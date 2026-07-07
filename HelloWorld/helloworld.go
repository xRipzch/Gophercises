package main

import "fmt" // FMT er go's standard libary til input og output

func main() {
	fmt.Println("hello world")
	greetings()
}

func greetings() {
	name := "Anders"
	age := 31
	// := opretter en variabel (Go gætter typen, udfra værdien) smort
	fmt.Printf("Navn: %s, Alder: %d \n", name, age)
}

// Formatings verbs
/*
%s  er string                                    // "Anna"
%d  er decimal (int)                             // 42
%f  er float (flydende decimaltal)               // 3.140000
%v  er standard-repræsentation af enhver værdi   // {Anna 25}
%+v er som %v, men viser feltnavne for structs   // {Navn:Anna Alder:25}
%#v er Go-syntaks-repræsentation af værdien      // main.Person{Navn:"Anna", Alder:25}
%T  er typen af værdien                          // main.Person
%t  er boolean (true/false)                      // true
%c  er et character (rune) ud fra Unicode-værdi  // 'A' (fra 65)
%q  er en "quoted" string                        // "Anna"
%x  er hexadecimal (småt)                        // 2a (fra 42)
%X  er hexadecimal (stort)                       // 2A (fra 42)
%o  er oktal                                     // 52 (fra 42)
%b  er binær                                     // 101010 (fra 42)
%e  er float i videnskabelig notation            // 3.140000e+00
%p  er pointer-adresse                           // 0xc0000140a0
%%  er et literal procenttegn                    // %
*/
