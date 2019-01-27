#!/usr/bin/env python3
with open('input') as f:
    lines = [i.rstrip('\n') for i in f.readlines()]
lines = lines[:-1]
ints = list(map(int, lines))
print(sum(ints))

done = False
current_freq = 0
already_there = set([0])
while not done:
    for i in ints:
        current_freq += i
        if current_freq in already_there:
            print(current_freq)
            done = True
            break
        already_there.add(current_freq)

