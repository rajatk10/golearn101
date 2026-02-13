package main

//Reading and Writing files in go.
import (
	"bufio"
	"fmt"
	"os"
)

/*
FILE I/O NOTES:

READING FILES:
1. os.Open(filename) - Opens file for reading, returns *os.File and error
2. os.ReadFile(filename) - Reads entire file into []byte
3. bufio.NewScanner(file) - Line-by-line reading
4. scanner.Scan() - Reads next line (returns false at EOF)
5. scanner.Text() - Gets current line as string

WRITING FILES:
1. os.Create(filename) - Creates new file or truncates existing
2. os.WriteFile(filename, data, perm) - Writes entire []byte to file
3. file.WriteString(content) - Writes string to file
4. os.OpenFile() - Open with flags (append, create, write)

IMPORTANT PATTERNS:
- defer file.Close() - Always close files (runs when function exits)
- Check errors immediately after each operation
- return in error blocks prevents using nil values

ERROR HANDLING IN GO:
- Errors are return values, not exceptions (unlike Python try/except)
- Must check err != nil after every operation
- return statement exits function to prevent crashes with nil values
- if err := func(); err != nil { } - compact error check pattern

COMMON OPERATIONS:
- os.Stat(filename) - Get file info (size, modified time)
- os.Remove(filename) - Delete file
- os.Rename(old, new) - Rename/move file
- os.IsNotExist(err) - Check if file doesn't exist

EXAMPLE PATTERN:
file, err := os.Open("file.txt")
if err != nil {
    return err  // Exit if error
}
defer file.Close()  // Cleanup when done
// Use file safely here
*/

func main() {
	file, err := os.Open("hello.go")
	if err != nil {
		fmt.Println("Unable to open file, please review the directory and file type")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	//Writing to file.
	filew, err := os.Create("Output1.txt")
	if err != nil {
		fmt.Println("Unable to write file, there is some issue with host/server")
		return
	}
	defer filew.Close()
	_, err = filew.WriteString("Hello from Go Read/Write Files i.e. file.go")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
	fmt.Println("File Written successfully")

	//Above there is basic and repetitive code so we are creating a function for it.
	if err := readAndPrintFile("maps1.go"); err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	if err := writeFile("Output2.txt", "Hello from Go Read/Write Files i.e. file.go"); err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
	fmt.Println("File Written successfully")

}

func readAndPrintFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Unable to open file %s : %w \n", filename, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

func writeFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Unable to create file %s : %w \n", filename, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("Unable to write to file %s : %w \n", filename, err)
	}
	return nil
}
