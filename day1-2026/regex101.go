package main

import (
	"fmt"
	"regexp"
)

/*
Find pattern
| Pattern | Meaning                 |
| ------- | ----------------------- |
| `.`     | any character           |
| `\d`    | digit (0–9)             |
| `\w`    | word (a-z, A-Z, 0-9, _) |
| `\s`    | whitespace              |
| `^`     | start of string         |
| `$`     | end of string           |

# How to quantify

| Pattern | Meaning         |
| ------- | --------------- |
| `*`     | 0 or more       |
| `+`     | 1 or more       |
| `?`     | 0 or 1          |
| `{n}`   | exactly n       |
| `{n,m}` | between n and m |

Some examples
*/
func userDetailsFromEmail(email string) {
	patternGroup := regexp.MustCompile(`^(\w+)@(\w+)\.(\w+)$`)
	matches := patternGroup.FindStringSubmatch(email)
	if len(matches) == 0 {
		fmt.Println("Invalid Email ", email)
		return
	}
	fmt.Println(matches)    //Print Whole Group
	fmt.Println(matches[1]) //Subsequent Group
	fmt.Println(matches[2])
	fmt.Println(matches[3])

	domain := matches[2] + matches[3]
	fmt.Println("Dmoain is ", domain)

	if numericsFromText(domain) {
		fmt.Println("Domain contains numerics")
		fmt.Println("Invalid Email ", email)
	} else {
		fmt.Println("Valid Email ", email)
	}
}
func numericsFromText(text string) bool {
	patternGroup := regexp.MustCompile(`(\d+)`)
	matches := patternGroup.MatchString(text)
	fmt.Println(matches)
	if matches {
		fmt.Println("Domain contains numerics ", text)
		return true
	}
	return false
}
func getPhoneNumber(mobileNo string) {
	pattern := regexp.MustCompile(`^\d{10}$`)
	matches := pattern.FindStringSubmatch(mobileNo)
	if len(matches) == 0 {
		fmt.Println("Invalid Mobile Number ", mobileNo)
		return
	}
	fmt.Println(matches)
	fmt.Println("Mobile no. is valid", mobileNo)
}

func getPhoneFromForms(text string) {
	pattern := regexp.MustCompile(`(\d{10})`)
	matches := pattern.FindStringSubmatch(text)
	if len(matches) == 0 {
		fmt.Println("Invalid Mobile Number ", text)
		return
	}
	fmt.Println(matches)
	fmt.Println("Mobile no. is valid", matches[1])
}

func editSentence(text string) {
	pattern := regexp.MustCompile(`(\d{10})`)
	clean := pattern.ReplaceAllString(text, "Valid Mobile Number")
	fmt.Println("Edit out the Mobile No. from User details")
	fmt.Println(clean)
}

func main() {
	//userDetailsFromEmail("apadma88@gmail.com")  //valid
	//userDetailsFromEmail("apadma88@gmail9.com") //Invalid email
	//userDetailsFromEmail("apadma88gmail.com")   //Invalid email
	//getPhoneNumber("1234567890")   //Valid
	//getPhoneNumber("gahah3456789") //Invalid
	//getPhoneFromForms("May name is Capgemini and my phone number is 1234567890")
	editSentence("May name is Capgemini and my phone number is 1234567890")
}
