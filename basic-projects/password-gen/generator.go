package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func generatePassword(passLength int, includeSpecial bool, includeSmallCase bool, includeLargeCase bool, includeNumbers bool) (string, error) {
	if passLength < 16 || passLength > 28 {
		//fmt.Println("Password length should be between 16 and 28")
		return "", errors.New("Password length should be between 16 and 28")
	}
	charset := ""
	if includeSpecial {
		charset += "!@#$%^&*()_+"
	}
	if includeNumbers {
		charset += "0123456789"
	}
	if includeLargeCase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if includeSmallCase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if len(charset) == 0 {
		return "", errors.New("Password Length is empty")
	}
	var password []byte

	for _ = range passLength {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		random := randomIndex.Int64()
		password = append(password, charset[random])
		//fmt.Println(password) //debug issues
	}
	return string(password), nil
}
