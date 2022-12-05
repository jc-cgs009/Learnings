def shellSort(l):
    n = len(l)
    gap = n//2
    while gap > 0:
        i = gap
        while i < n:
            gvalue = l[i]
            j = i - gap
            while j >= 0 and l[j] > gvalue:
                l[j+gap] = l[j]
                j = j - gap
            l[j+gap] = gvalue
            i += 1
        gap = gap//2

l = [3,5,8,9,6,2]
print("Original list: ", l)
shellSort(l)
print("Sorted list: ", l) 