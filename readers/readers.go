package readers

import (
    "bufio"
    "fmt"
    "os"
)

func Readfile() string {
    // Prompt the user to enter the file path
    fmt.Print("Enter the file path: ")
    
    // Create a scanner to read user input from standard input (keyboard)
    scanner := bufio.NewScanner(os.Stdin)
    
    // Read the input from the user
    if scanner.Scan() {
        // Get the file path entered by the user
        inputFile := scanner.Text()
        
        // Read the contents of the file
        data, err := os.ReadFile(inputFile)
        if err != nil {
            
            return err.Error()
        }
        fmt.Println("File contents:", string(data))
        return string(data)

    } else {
        fmt.Println("Error reading input:", scanner.Err())
        
    } 
    
}
