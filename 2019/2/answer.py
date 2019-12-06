#!/usr/bin/env python3
with open('input') as f:
    input_data = f.read()
input_data = input_data.rstrip()
program_code = list(map(int, input_data.split(',')))

def opcode(program, position):
    if program[position] == 1:
        program[program[position + 3]] = program[program[position + 1]] + program[program[position + 2]]
    elif program[position] == 2:
        program[program[position + 3]] = program[program[position + 1]] * program[program[position + 2]]

program_code[1], program_code[2] = 12, 2

def execute_program(program_code):
    current_position = 0
    while program_code[current_position] != 99:
        opcode(program_code, current_position)
        current_position += 4

execute_program(program_code)
print(program_code[0])

# TODO: something better than brute force
for i in range(1,99):
    for j in range(1,99):
        program_code = list(map(int, input_data.split(',')))
        program_code[1], program_code[2] = i, j
        execute_program(program_code)
        if program_code[0] == 19690720:
            print(i * 100 + j)
            break
