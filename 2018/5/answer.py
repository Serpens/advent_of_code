molecule = 'dabAcCaCBAcCcaDA'
def transmute_once(molecule):
    result = ''
    for i in zip(molecule[:-1], molecule[1:]):
        if i[0] != upper(i[1]) and i[0] != lower(i[1]):
            result += i[0]
    return result
transmute_once(molecule)
def transmute_once(molecule):
    result = ''
    for i in zip(molecule[:-1], molecule[1:]):
        if i[0] != i[1].upper() and i[0] != i[1].lower():
            result += i[0]
    return result
transmute_once(molecule)
molecule
def transmute_once(molecule):
    result = ''
    for i in zip(molecule[:-1], molecule[1:]):
        if i[0] != i[1].upper() and i[0] != i[1].lower():
            result += i[0]
    return result
molecule
transmute_once(molecule)
def transmute_once(molecule):
    result = ''
    for i in range(len(molecule) - 1):
        if molecule[i].upper() == molecule[i+1] or molecule[i].lower() == molecule[i+1]:
            return molecule[:i] + molecule[(i+1):]
    return molecule
transmute_once(molecule)
molecule
def transmute_once(molecule):
    result = ''
    for i in range(len(molecule) - 1):
        if molecule[i].upper() == molecule[i+1] or molecule[i].lower() == molecule[i+1]:
            return molecule[:i-1] + molecule[(i+1):]
    return molecule
transmute_once(molecule)
molecule
def transmute_once(molecule):
    result = ''
    for i in range(len(molecule) - 1):
        if molecule[i].upper() == molecule[i+1] or molecule[i].lower() == molecule[i+1]:
            return molecule[:i] + molecule[(i+2):]
    return molecule
molecule
transmute_once(molecule)
transmute_once(transmute_once(molecule))
transmute_once(transmute_once(transmute_once(molecule)))
df transmute(molecule):
def transmute(molecule):
    old_molecule = molecule
    new_molecule = transmute_once(old_molecule)
    while old_molecule != new_molecule:
        old_molecule = new_molecule
        new_molecule = transmute_once(old_molecule)
    return new_molecule
transmute(molecule)
with open('input') as f:
    input = f.read()
input[-10:]
with open('input') as f:
    input_molecule = f.read()
input_molecule = input_molecule.rstrip('\n')
len(input_molecule)
transmute(input_molecule)
transmuted_molecule = transmute(input_molecule)
len(transmuted_molecule)
len(input_molecule)

for i in string.ascii_lowercase:
    potential_molecule = input_molecule.replace(i, '').replace(i.upper(), '')
    potential_molecule = transmute(potential_molecule)
    print(i, len(potential_molecule))


