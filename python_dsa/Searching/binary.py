def binarySearch(l, key):
    s = 0
    e = len(l)-1
    while s <= e:
        mid = (s+e)//2
        if l[mid] == key:
            return mid
        elif key < l[mid]:
            e = mid-1
        elif key > l[mid]:
            s = mid+1
    return -1

l = [15,21,47,84,96]
found = binarySearch(l, 47)
print(f"Key found at index {found}")