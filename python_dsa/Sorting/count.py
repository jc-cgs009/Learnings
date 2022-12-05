def countSort(l):
    n = len(l)
    maxnum = max(l)
    carray = [0] * (maxnum+1)
    for i in range(n):
        carray[l[i]] = carray[l[i]]+1
    
    i,j = 0, 0
    while j < maxnum+1:
        if carray[j] > 0:
            l[i] = j
            carray[j] -= 1
            i += 1
        else:
            j += 1



l = [3,5,8,9,6,2,3,5,5]
print("Original list: ", l)
countSort(l)
print("Sorted list: ", l)
