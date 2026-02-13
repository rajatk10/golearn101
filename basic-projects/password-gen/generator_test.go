package main

import (
	"strings"
	"testing"
)

func TestBasicPassword(t *testing.T) {
	password, err := generatePassword(16, true, true, true, true)
	//Assert1 check length
	if len(password) != 16 {
		t.Errorf("Password length is not 16")
	}
	//Assert2 check error is nil
	if err != nil {
		t.Errorf("Error should be nil")
	}
	//Assert3 check it is not empty
	if password == "" {
		t.Errorf("Password is empty")
	}
	//Assert4 check it contains special characters
	if !strings.ContainsAny(password, "!@#$%^&*()_+") {
		t.Errorf("Password does not contain special characters")
	}
	//Assert5 check it contains smallcase
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		t.Errorf("Password does not contain smallcase characters")
	}
	//Assert6 Check it contains largecase characters
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		t.Errorf("Password does not contain largecase characters")
	}
	//Assert7 Check it contains numbers
	if !strings.ContainsAny(password, "0123456789") {
		t.Errorf("Password does not contain numbers")
	}
}

func TestInvalidSmallerLength(t *testing.T) {
	password, err := generatePassword(15, true, true, true, true)
	//Assert1
	if err == nil {
		t.Errorf("Error should not be nil")
	}
	if password != "" {
		t.Errorf("Expected empty password on error")
	}
}

func TestInvalidLargerLength(t *testing.T) {
	password, err := generatePassword(29, true, true, true, true)
	if err == nil {
		t.Errorf("Error should not be nil")
	}
	if password != "" {
		t.Errorf("Expected empty password on error")
	}
}

func TestUnoptedSpecialChars(t *testing.T) {
	password, err := generatePassword(18, false, true, true, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	if strings.Contains(password, "!@#$%^&*()_+") {
		t.Errorf("Password should not contain special characters")
	}
}

func TestUnoptedSmallCaseChars(t *testing.T) {
	password, err := generatePassword(16, true, false, true, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	if strings.Contains(password, "abcdefghijklmnopqrstuvwxyz") {
		t.Errorf("Password should not contain smallcase characters")
	}
}

func TestAllFlagsFalse(t *testing.T) {
	password, err := generatePassword(16, false, false, false, false)
	if err == nil {
		t.Errorf("Error should not be nil")
	}
	if password != "" {
		t.Errorf("Expected empty password on error")
	}
}
