#!/usr/bin/env python
with open('input') as f:
    input_data = [int(i) for i in f.readlines()]
fuel = map(lambda x: int(x/3) - 2, input_data)
print(sum(fuel))
# reduce(lambda x, y: x+y, map(lambda x: int(x/3) - 2, input_data))

def total_fuel_needed(x):
    ans = int(x/3) - 2
    if ans <= 0:
        return 0
    return ans + total_fuel_needed(ans)

print(sum(map(total_fuel_needed, input_data)))

