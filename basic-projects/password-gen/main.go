package main

import (
	"flag"
	"fmt"
)

func main() {
	passLength := flag.Int("length", 16, "Password Length")
	includeSpecial := flag.Bool("special", true, "Include Special Characters")
	includeSmallCase := flag.Bool("small", true, "Include Small Case")
	includeLargeCase := flag.Bool("large", true, "Include Large Case")
	includeNumbers := flag.Bool("numbers", true, "Include Numbers")
	flag.Parse()
	password, err := generatePassword(*passLength, *includeSpecial, *includeSmallCase, *includeLargeCase, *includeNumbers)
	if err != nil {
		fmt.Println("Error While creating password: ", err)
		return
	}
	fmt.Println("Here you go :", password)

}
