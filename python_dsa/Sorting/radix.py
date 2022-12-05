def radixSort(l):
    n = len(l)
    maxnum = max(l)
    digits = len(str(maxnum))
    bins = [[]] * 10

    for i in range(digits):
        for j in range(n):
            e = ((l[j]//10**i)%10)
            if len(bins[e]) > 0:
                bins[e].append(l[j])
            else:
                bins[e] = [l[j]]

        k = 0

        for x in range(10):
            if len(bins[x]) > 0:
                for y in range(len(bins[x])):
                    l[k] = bins[x].pop(0)
                    k += 1
l = [63,250,835,947,651,28]
print("Original list: ", l)
radixSort(l)
print("Sorted list: ", l)     