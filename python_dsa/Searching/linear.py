def linearSearch(l, key):
    index = 0
    while index < len(l):
        if l[index] == key:
            return index
        index+=1
    return -1

l = [84,21,47,96,56]
found = linearSearch(l, 211)
print(f"Key found at index {found}")