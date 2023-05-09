package main

import (
    "bufio"
    "os"
    "strings"
)

func main() {
    inputFile := "input.txt"
    outputFile := "output.txt"

    // Open input file for reading
    input, err := os.Open(inputFile)
    if err != nil {
        panic(err)
    }
    defer input.Close()

    // Open output file for writing
    output, err := os.Create(outputFile)
    if err != nil {
        panic(err)
    }
    defer output.Close()

    // Read input file line by line and replace 'XX' with 'tt' and '[.]' with '.'
    scanner := bufio.NewScanner(input)
    for scanner.Scan() {
        line := scanner.Text()
        line = strings.ReplaceAll(line, "XX", "tt")
        line = strings.ReplaceAll(line, "[.]", ".")
        
        // Write modified line to output file
        _, err := output.WriteString(line + "\n")
        if err != nil {
            panic(err)
        }
    }
}
