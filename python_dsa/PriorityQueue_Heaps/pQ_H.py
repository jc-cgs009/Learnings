class Heap:
    def __init__(self):
        self._maxsize = 10
        self._data = [-1] * self._maxsize
        self._csize = 0

    def __len__(self):
        return len(self._data)

    def isempty(self):
        return self._csize == 0

    def insert(self, e):
        if self._csize == self._maxsize:
            print("No Space!...")
            return
        self._csize += 1
        hi = self._csize
        while hi > 1 and e > self._data[hi//2]:
            self._data[hi] = self._data[hi//2]
            hi = hi//2
        self._data[hi] = e

    def max(self):
        if self.isempty():
            print("Heap is empty!...")
            return
        return self._data[1]

    def deletemax(self):
        if self.isempty():
            print("Heap is empty!...")
            return
        e = self._data[1]
        self._data[1] = self._data[self._csize]
        self._data[self._csize] = -1
        self._csize -= 1
        i = 1
        j = i * 2
        while j <= self._csize:
            if self._data[j] < self._data[j + 1]:
                j = j + 1
            if self._data[i] < self._data[j]:
                temp = self._data[i]
                self._data[i] = self._data[j]
                self._data[j] = temp
                i = j
                j = i * 2
            else:
                break
        return e


if __name__ == "__main__":
    H = Heap()
    H.insert(20)
    H.insert(14)
    H.insert(2)
    H.insert(15)
    H.insert(10)
    H.insert(21)
    print(H._data)
    print(H.max())
    print(H.deletemax())
    print(H._data)
    print(H.max())


