#!/usr/bin/env python3
# I'm not using collections, it swould be too easy
with open('input') as f:
    lines = [i.rstrip('\n') for i in f.readlines()]

def repeats_2_or_3_times(input_str):
    counts = {}
    for i in input_str:
        if i in counts:
            counts[i] += 1
        else:
            counts[i] = 1
    times2, times3 = False, False
    for k,v in counts.items():
        if v==2:
            times2 = True
        elif v==3:
            times3 = True
    return (times2, times3)

counted_repeats = list([repeats_2_or_3_times(i) for i in lines])
print(sum([i[0] for i in counted_repeats]) * sum([i[1] for i in counted_repeats]))


def common_letters(s1, s2):
    result = []
    for i in zip(s1, s2):
        if i[0] == i[1]:
            result.append(i[0])
    return ''.join(result)

current_common_letters = ''
target_len = len(lines[0]) - 1
for l1 in lines:
    for l2 in lines:
        current_common_letters = common_letters(l1, l2)
        if len(current_common_letters) == target_len:
            print(current_common_letters)
            break
    if len(current_common_letters) == target_len:
        break

