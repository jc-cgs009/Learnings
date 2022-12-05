def quickSort(l, low, high):
    if low < high:
        pi = partition(l , low, high)
        quickSort(l , low, pi-1)
        quickSort(l, pi+1, high)

def partition(l, low, high):
    pivot = l[low]
    i = low+1
    j = high
    while True:
        while i <= j and l[i] <= pivot:
            i += 1
        
        while i <= j and l[j] > pivot:
            j -= 1

        if i <= j:
            l[i], l[j] = l[j], l[i]
        else:
            break
    l[low], l[j] = l[j], l[low]
    return j

l = [3,5,8,9,6,2]
print("Original list: ", l)
quickSort(l, 0, len(l)-1)
print("Sorted list: ",l)
