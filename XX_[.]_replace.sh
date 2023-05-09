#!/bin/bash

inputFile="input.txt"
outputFile="output.txt"

# Read input file and replace 'XX' with 'tt' and '[.]' with '.'
while read line; do
    line=${line//XX/tt}
    line=${line//"[.]"/"."}
    
    # Write modified line to output file
    echo "$line" >> "$outputFile"
done < "$inputFile"
