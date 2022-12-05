def selectionSort(l):
    n = len(l)
    for i in range(n):
        position = i
        for j in range(i+1,n):
            if l[j] < l[position]:
                position = j
        temp = l[i]
        l[i] = l[position]
        l[position] = temp

l = [3,5,8,9,6,2]
print("Original list: ", l)
selectionSort(l)
print("Sorted list: ", l) 