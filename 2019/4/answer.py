#!/usr/bin/env python3
with open('input') as f:
    input_data = f.read()
input_data
input_data.rstrip()
input_data.rstrip().split('-')
pass_min, pass_max = input_data.rstrip().split('-')
pass_min, pass_max = int(pass_min), int(pass_max)

def verify(password):
    digits = str(password)
    duplicate = False
    previous = -1
    for d in digits:
        if previous == d:
            duplicate = True
        if int(d) < int(previous):
            return False
        previous = d
    return duplicate

valid_passwords = 0
print(sum(map(verify, range(pass_min, pass_max))))

def verify2(password):
    digits = str(password)
    counts = [0] * 10
    previous = -1
    for d in digits:
        counts[int(d)] += 1
        if int(d) < int(previous):
            return False
        previous = d
    return 2 in counts

valid_passwords = 0
print(sum(map(verify2, range(pass_min, pass_max))))

