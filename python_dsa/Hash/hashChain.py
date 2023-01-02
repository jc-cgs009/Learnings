from singlyLinkedList import singlyLinkedList

class hashChain:
    def __init__(self):
        self._hashsize = 10
        self._hashtable = [0] * self._hashsize
        for i in range(self._hashsize):
            self._hashtable[i] = singlyLinkedList()

    def hashCode(self, key):
        return key % self._hashsize

    def insert(self, e):
        i = self.hashCode(e)
        self._hashtable[i].insertSorted(e)

    def search(self, key):
        i = self.hashCode(key)
        return self._hashtable[i].search(key) != -1

    def display(self):
        for i in range(self._hashsize):
            print('[',i,']', end = " ")
            self._hashtable[i].display()
        print()

H = hashChain()
H.insert(54)
H.insert(78)
H.insert(64)
H.insert(92)
H.insert(34)
H.insert(86)
H.insert(28)
H.display()
print(H.search(74))
