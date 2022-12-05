def bubbleSort(l):
    n = len(l)
    for passes in range(n-1, 0, -1):
        for i in range(passes):
            if l[i] > l[i+1]:
                temp = l[i]
                l[i] = l[i+1]
                l[i+1] = temp

l = [3,5,8,9,6,2]
print("Original list: ", l)
bubbleSort(l)
print("Sorted list: ", l) 