package main

import (
	"fmt"
	"strings"
)

/*
STRINGS NOTES:
1. Strings are immutable - each operation creates a new string
2. Strings are UTF-8 encoded byte sequences
3. len(str) returns number of BYTES, not characters
4. Access by index str[i] gives BYTE (uint8), not character

STRING OPERATIONS:
- Concatenation: str1 + str2
- Length: len(str) (bytes, not characters)
- Indexing: str[0] (returns byte)
- Iteration: for i := range str (i is index)
- Two-variable range: for i, r := range str (r is rune)

RUNE NOTES:
1. Rune = int32 type representing Unicode code point
2. Rune holds the numeric value of a character
3. Byte (uint8) = 0-255, only handles ASCII
4. Rune (int32) = can hold all Unicode characters
5. Use %c format verb to print rune/byte as character

BYTE vs RUNE:
- str[i] → byte (uint8) → ASCII only, prints as number
- range str with 2 vars → rune (int32) → Unicode, prints as number
- Use %c to display as character: fmt.Printf("%c", r)

UNICODE vs UTF-8:
- Unicode: Standard that assigns numbers to characters (U+0041 = 'A')
- UTF-8: Encoding format that stores Unicode as bytes (1-4 bytes per char)
- Go strings: UTF-8 encoded
- Go runes: Unicode code points

EXAMPLE:
str := "世" (Chinese character)
- Unicode code point: U+4E16
- UTF-8 encoding: 3 bytes (0xE4, 0xB8, 0x96)
- len(str) = 3 (bytes)
- Rune count = 1 (character)
*/

func main() {
	str := "Hello World"
	fmt.Println(str)
	//iterate loop over characters
	for i := range str {
		fmt.Println("Index : ", i, "Character is: ", str[i]) //prints ASCII value instead of character
		fmt.Printf("Index : %d Character is: %c\n", i, str[i])
		//character is string and which is where rune comes into picture
		//fmt.Println("Index : ", i, "Character is: ", r)

	}
	fmt.Println("Length of string is ", len(str))
	fmt.Printf("%s does it contain \"%s\"? %t \n", str, "World", strings.Contains(str, "World"))
	fmt.Printf("String replace World %s \n", strings.ReplaceAll(str, "World", "Universe"))
	fmt.Printf("String to all caps %s \nAnd LOWER is %s \n", strings.ToUpper(str), strings.ToLower(str))
	str2 := "   Hello Space String  "
	fmt.Printf("String trim spaces %s \n", strings.TrimSpace(str2))
	fmt.Printf("String as prefix 'Hello' in str : %t \n", strings.HasPrefix(str, "He"))
}
