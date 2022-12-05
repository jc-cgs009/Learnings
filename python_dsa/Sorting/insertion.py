def insertionSort(l):
    n = len(l)
    for i in range(n):
        cvalue = l[i]
        position = i
        while position > 0 and l[position-1] > cvalue:
            l[position] = l[position - 1]
            position = position - 1
        l[position] = cvalue

l = [3,5,8,9,6,2]
print("Original list: ", l)
insertionSort(l)
print("Sorted list: ", l) 