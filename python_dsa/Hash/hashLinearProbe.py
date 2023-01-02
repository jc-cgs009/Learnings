class hashLinearProbe:
    def __init__(self):
        self._hashsize = 10
        self._hashtable = [0] * self._hashsize

    def hashcode(self, key):
        return key % self._hashsize

    def lprobe(self, key):
        i = self.hashcode(key)
        j = 0

        while self._hashtable[(i + j) % self._hashsize] != 0:
            j += 1

        return (i + j) % self._hashsize

    def insert(self, e):
        i = self.hashcode(e)
        if self._hashtable[i] == 0:
            self._hashtable[i] = e
        else:
            i = self.lprobe(e)
            self._hashtable[i] = e

    def search(self, key):
        i = self.hashcode(key)
        j = 0

        while self._hashtable[(i + j) % self._hashsize] != key:
            if self._hashtable[(i + j) % self._hashsize] == 0:
                return False
            j += 1

        return True

    def display(self):
        print(self._hashtable)

H =hashLinearProbe()
H.insert(54)
H.insert(78)
H.insert(64)
H.insert(92)
H.insert(34)
H.insert(28)
H.insert(86)
H.display()
print(H.search(74))