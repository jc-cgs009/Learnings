from pQ_H import Heap

def heapSort(A):
    n = len(A)
    H = Heap()
    for i in range(n):
        H.insert(A[i])
    
    k = n - 1
    for i in range(H._csize):
        A[k] = H.deletemax()
        k -= 1

A = [20, 14, 2, 15, 10, 21]
print("Original Array : ", A)
heapSort(A)
print("Sorted Array : ", A)
