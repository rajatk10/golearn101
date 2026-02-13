package main

import "fmt"

/*
CONTROL FLOW NOTES:
IF-ELSE:
1. Basic: if condition { } else { }
2. No parentheses around condition (unlike C/Java)
3. Short statement: if x := getValue(); x > 0 { }

FOR LOOPS:
1. Standard: for i := 0; i < 10; i++ { }
2. While-style: for condition { } (no while keyword in Go)
3. Infinite: for { } (use break to exit)
4. Range: for i, val := range slice { }
5. Range over int (Go 1.22+): for i := range 10 { } (0 to 9)
6. Custom increment: for i := 0; i < 10; i += 2 { }

SWITCH:
1. Basic: switch variable { case value: ... }
2. No break needed (automatic)
3. Multiple values: case 1, 2, 3:
4. No condition: switch { case x > 0: ... } (like if-else chain)
5. Fallthrough: use fallthrough keyword to continue to next case
*/

func main() {
	//control flow learning
	var bo string
	var flag bool = true
	if flag {
		fmt.Println("Flag is true, set value of bo")
		bo = "bo is true"
		fmt.Println(bo)
	} else {
		fmt.Println("Flag is unset")
		bo = "bo is false"
		fmt.Println(bo)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Usual while loop", i)
	}

	// There is no while loop but it can be mimiced using for

	j := 0
	for j < 5 {
		fmt.Println("Mimic while loop using for ", j)
		j++
	}

	// switch case
	var day int = 3
	switch day {
	case 1:
		fmt.Println("Day is Monday")
	case 2:
		fmt.Println("Day is Tuesday")
	case 3:
		fmt.Println("Day is Wednesday")
	default:
		fmt.Println("Day is unknown, unspecified")
	}

}
