#!/usr/bin/env python3
with open('input') as f:
    lines = [i.rstrip('\n') for i in f.readlines()]

def get_inches(line):
    start = [int(i) for i in line.split()[2][:-1].split(',')]
    max_i, max_j = [int(i) for i in line.split()[3].split('x')]
    result = []
    for i in range(max_i):
        for j in range(max_j):
            result.append((start[0] + i, start[1] + j))
    return result

already_claimed = set()
duplicate_claims = set()

for l in lines:
    current_inches = get_inches(l)
    for i in current_inches:
        if i in already_claimed:
            duplicate_claims.add(i)
        already_claimed.add(i)

print(len(duplicate_claims))

def check_overlap(line1, line2):
    ends1, ends2 = get_ends(line1), get_ends(line2)
    if (ends1[0] > ends2[0] and ends1[0] < ends2[2]) or (ends1[2] > ends2[0] and ends1[2] < ends2[2]):
        if (ends1[1] > ends2[1] and ends1[1] < ends2[3]) or (ends1[3] > ends2[1] and ends1[3] < ends2[3]):
            return True
    return False


has_overlap = [False for i in lines]

for i in enumerate(lines):
    for j in enumerate(lines):
        if i[0] >= j[0]:
            continue
        if check_overlap(i[1], j[1]):
            has_overlap[i[0]], has_overlap[j[0]] = True, True

for i in zip(lines, has_overlap):
    if not i[1]:
        print(i[0])


def check_overlap(line1, line2):
    ends1, ends2 = get_ends(line1), get_ends(line2)
    if (ends1[0] >= ends2[0] and ends1[0] <= ends2[2]) or (ends1[2] >= ends2[0] and ends1[2] <= ends2[2]):
        if (ends1[1] >= ends2[1] and ends1[1] <= ends2[3]) or (ends1[3] >= ends2[1] and ends1[3] <= ends2[3]):
            return True
    if (ends2[0] >= ends1[0] and ends2[0] <= ends1[2]) or (ends2[2] >= ends1[0] and ends2[2] <= ends1[2]):
        if (ends2[1] >= ends1[1] and ends2[1] <= ends1[3]) or (ends2[3] >= ends1[1] and ends2[3] <= ends1[3]):
            return True
    return False
has_overlap = [False for i in lines]
for i in enumerate(lines):
    for j in enumerate(lines):
        if i[0] >= j[0]:
            continue
        if check_overlap(i[1], j[1]):
            has_overlap[i[0]], has_overlap[j[0]] = True, True
for i in zip(lines, has_overlap):
    if not i[1]:
        print(i[0])


