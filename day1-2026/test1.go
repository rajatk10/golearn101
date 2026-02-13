package main

import (
	"errors"
	"fmt"
)

/*
TESTING IN GO NOTES:

TEST FILE NAMING:
- Test files must end with _test.go (e.g., test1_test.go)
- Must be in same package as code being tested
- Convention: filename_test.go tests filename.go

TEST FUNCTION RULES:
1. Function name: TestXxx (must start with Test, capital letter after)
2. Parameter: func TestXxx(t *testing.T)
3. Use t.Error() or t.Errorf() to report failures
4. Use t.Fatal() or t.Fatalf() to stop test immediately

RUNNING TESTS:
- go test                    # Run all tests in current directory
- go test -v                 # Verbose output (shows each test)
- go test -cover             # Show code coverage percentage
- go test -run TestAdd       # Run specific test
- go test file.go file_test.go  # Test specific files

TEST ASSERTIONS:
- t.Errorf("msg", args)      # Report error, continue test
- t.Fatalf("msg", args)      # Report error, stop test
- t.Error("msg")             # Simple error message
- t.Log("msg")               # Log message (only shown with -v)

COVERAGE:
- coverage: 40.0% means 40% of code lines were executed by tests
- Higher coverage = more code tested
- Aim for 70-80% coverage in production code

ERROR TESTING PATTERN:
result, err := functionThatReturnsError()
if err != nil {
    t.Errorf("unexpected error: %v", err)
}
// Or test that error SHOULD occur:
if err == nil {
    t.Error("expected error, got nil")
}
*/

func add(a, b int) int {
	return a + b
}
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ZeroDivisonError")
	}
	return a / b, nil
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(divide(1, 2))

	result, err := divide(1, 0)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println(result)
}
