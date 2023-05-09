input_file = "input.txt"
output_file = "output.txt"

with open(input_file, "r") as f1, open(output_file, "w") as f2:
    for line in f1:
        line = line.replace('XX', 'tt').replace('[.]', '.')
        f2.write(line)
