#!/usr/bin/env python3

ith open('input') as f:
    lines = [i.rstrip('\n') for i in f.readlines()]
lines.sort()

guard_lines = OrderedDict()
for l in lines:
    if 'begins shift' in l:
        current_guard = l.split()[3]
        if current_guard not in guard_lines:
            guard_lines[current_guard] = []
    guard_lines[current_guard].append(l)

def get_sleep_wakeup_ranges(lines):
    times = [int(i.split(' ')[1][:-1].split(':')[1]) for i in lines if 'begins shift' not in i]
    return zip(times[::2], times[1::2])

for guard in guard_lines.keys():
    guard_naps = get_sleep_wakeup_ranges(guard_lines[guard])
    guard_minutes = [0 for i in range(60)]
    for i in guard_naps:
        for j in range(i[0], i[1]):
            guard_minutes[j] += 1
    guard_minutes_asleep[guard] = guard_minutes


for k,v in guard_minutes_asleep.items():
    print(k, sum(v))


for k,v in guard_minutes_asleep.items():
    print(k, sum(v))
for i in enumerate(guard_minutes_asleep['#2441']):
    if i[1] == max(guard_minutes_asleep['#2441']):
        print(i[0])



