def binaryRecursiveSearch(l, key,s , e):
    if s > e:
        return -1
    else:
        mid = (s+e)//2
        if l[mid] == key:
            return mid
        elif key < l[mid]:
            return binaryRecursiveSearch(l, key, s, mid-1)
        elif key > l[mid]:
            return binaryRecursiveSearch(l, key, mid+1, e)

l = [15,21,47,84,96]
found = binaryRecursiveSearch(l, 470, 0, len(l)-1)
print(f"Key found at index {found}")