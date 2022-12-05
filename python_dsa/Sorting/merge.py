def mergeSort(l,left,right):
    if left < right:
        mid = (left+right)//2
        mergeSort(l,left,mid)
        mergeSort(l,mid+1,right)
        merge(l,left,mid,right)

def merge(l,left,mid,right):
    i = left
    j = mid+1
    k = left
    l1 = [0] * (right+1)
    while i <= mid and j <= right:
        if l[i] < l[j]:
            l1[k] = l[i]
            i = i + 1
        else:
            l1[k] = l[j]
            j = j + 1
        k = k + 1

    while i <= mid:
        l1[k] = l[i]
        i += 1
        k += 1

    while j <= right:
        l1[k] = l[j]
        j += 1
        k += 1

    for x in range(left,right+1):
        l[x] = l1[x]

l = [3,5,8,9,6,2]
print("Original list: ", l)
mergeSort(l,0,(len(l)-1))
print("Sorted list: ",l)