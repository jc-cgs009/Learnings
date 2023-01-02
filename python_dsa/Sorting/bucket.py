def bucketSort(A):
    n = len(A)
    maxelement = max(A)
    l = []
    buckets = [l] * 10
    for i in range(n):
        j = int(n * A[i] / (maxelement + 1))
        if len(buckets[j]) == 0:
            buckets[j] = [A[i]]
        else:
            buckets[j].append(A[i])

    for i in range(10):
        insertionSort(buckets[i])

    k = 0
    for i in range(10):
        for j in range(len(buckets[i])):
            A[k] = buckets[i].pop(0)
            k += 1

def insertionSort(l):
    n = len(l)
    for i in range(n):
        cvalue = l[i]
        position = i
        while position > 0 and l[position-1] > cvalue:
            l[position] = l[position - 1]
            position = position - 1
        l[position] = cvalue

A = [3,5,8,9,6,2]
print("Original list: ", A)
bucketSort(A)
print("Sorted list: ", A)