$inputFile = "input.txt"
$outputFile = "output.txt"

# Read input file and replace 'XX' with 'tt' and '[.]' with '.'
$content = Get-Content $inputFile | ForEach-Object {
    $_ -replace "XX", "tt" -replace "\[\.\]", "."
}

# Write modified content to output file
Set-Content $outputFile $content
