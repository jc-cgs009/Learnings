class _Node:
    __slots__ = "_element", "_next"
    def __init__(self, element, next):
        self._element = element
        self._next = next

class dequeueLinkedList:
    def __init__(self):
        self._front = None
        self._rare = None
        self._size = 0

    def __len__(self):
        return self._size

    def isempty(self):
        return self._size == 0

    def addlast(self, e):
        new = _Node(e, None)
        if self.isempty():
            self._front = new
            self._rare = new
        else:
            self._rare._next = new
            self._rare = new
        self._size += 1

    def addfirst(self, e):
        new = _Node(e, None)
        if self.isempty():
            self._front = new
            self._rare = new
        else:
            new._next = self._front
            self._front = new
        self._size += 1

    def removefirst(self):
        r = None
        if self.isempty():
            print("DEQueue is empty!...")
        else:
            r = self._front._element
            self._front = self._front._next
            self._size -= 1
        
        if self.isempty():
            self._front = None
            self._rare = None
        
        return r

    def removelast(self):
        r = None
        if self.isempty():
            print("DEQueue is empty!...")
        else:
            p = self._front
            i = 1
            while i < len(self)-1:
                p = p._next
                i += 1
            r = self._rare._element
            self._rare = p
            self._rare._next = None
            self._size -= 1
        
        if self.isempty():
            self._front = None
            self._rare = None
        
        return r

    def first(self):
        e = None
        if self.isempty():
            print("DEQueue is empty!...")
        else:
            e = self._front._element
        return e

    def last(self):
        e = None
        if self.isempty():
            print("DEQueue is empty!...")
        else:
            e = self._rare._element
        return e

    def display(self):
        p = self._front
        while p:
            print(p._element, end = "<-->")
            p = p._next
        print()

dqll = dequeueLinkedList()
dqll.addfirst(10)
print(dqll.first())
print(dqll.last())
dqll.display()
# print(dqll.removefirst())
print(dqll.removelast())
dqll.display()
dqll.addfirst(10)
dqll.addfirst(20)
dqll.addfirst(30)
dqll.display()
dqll.addlast(15)
dqll.addlast(25)
dqll.addlast(35)
dqll.display()
print(dqll.first())
print(dqll.last())
dqll.display()
print(dqll.removefirst())
print(dqll.removelast())
dqll.display()


