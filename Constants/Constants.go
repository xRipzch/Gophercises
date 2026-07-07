// go supporter consts af chars, string, bools og numeric værdier
// const deklare en constant værdi, som JS
// const opretter en fast værdi (kan ikke ændres, kendes ved kompilering)

package main

import (
	"fmt"
	"math"
)

const s string = "constant" // låst, også fordi det er ude fra function

func main() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d, ":: <-  const n(500000), dividerde med 3e20")

	fmt.Println(int64(d))
	fmt.Println("overstående: int64(d)")

	fmt.Println(math.Sin(n), "math.Sin(n)")
}
