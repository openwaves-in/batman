package readerss

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Readd() string {
	// Prompt the user to input the file path
	fmt.Print("Enter the file path: ")

	// Read the input from the user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := strings.TrimSpace(scanner.Text())
	return a
}
