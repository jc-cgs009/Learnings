class _Node:
    __slots__ = "_element", "_next", "_prev"
    def __init__(self, element, next, prev):
        self._element = element
        self._next = next
        self._prev = prev

class doublyLinkedList:
    def __init__(self):
        self._head = None
        self._tail = None
        self._size = 0

    def __len__(self):
        return self._size
    
    def isempty(self):
        return self._size == 0

    def addlast(self, e):
        new = _Node(e, None, None)
        if self.isempty():
            self._head = new
            self._tail = new
        else:
            self._tail._next = new
            new._prev = self._tail
            self._tail = new
        self._size += 1
    
    def addfirst(self, e):
        new = _Node(e, None, None)
        if self.isempty():
            self._head = new
            self._tail = new
        else:
            new._next = self._head
            self._head._prev = new
            self._head = new
        self._size += 1

    def addany(self, e, position):
        if position <= 0:
            self.addfirst(e)
        elif position >= self.__len__():
            self.addlast(e)
        else:
            new = _Node(e, None, None)
            p = self._head
            i = 1
            while i < position:
                p = p._next
                i += 1
            new._next = p._next
            p._next._prev = new
            new._prev = p
            p._next = new
            
            self._size += 1

    def removefirst(self):
        r = None
        if self.isempty():
            print("Linked list is empty!..")
        else:
            r = self._head._element
            self._head = self._head._next
            if self._head == None:
                self._tail = None
            else:
                self._head._prev = None
            self._size -= 1
    
        return r

    def removelast(self):
        r = None 
        if self.isempty():
            print("Linked list is empty!..")
        else:
            r = self._tail._element
            self._tail = self._tail._prev
            if self._tail == None:
                self._head = None
            else:
                self._tail._next = None
            self._size -= 1 
        return r
    
    def removeany(self, position):
        r = None
        if position < 0 or position >= self.__len__():
            print("List index out of range!..")
        elif position == 0:
            r = self.removefirst()
        elif position == self.__len__()-1:
            r = self.removelast()
        else:
            p = self._head
            i = 1
            while i < position:
                p = p._next
                i += 1
            r = p._next._element
            p._next = p._next._next
            p._next._prev = p
            self._size -= 1
        return r


    def display(self):
        p = self._head
        while p:
            print(p._element, end = "-->")
            p = p._next
        print()
    
    def displayReverse(self):
        p = self._tail
        while p:
            print(p._element, end="<--")
            p = p._prev
        print()

    def search(self, key):
        p = self._head
        index = 0
        while p:
            if p._element == key:
                return index
            p = p._next
            index += 1
        return -1


dl = doublyLinkedList()
dl.addlast(10)
dl.addlast(20)
dl.addlast(30)
dl.display()
dl.addfirst(40)
dl.addfirst(50)
dl.display()
dl.addany(100, 3)
dl.display()
print(dl.search(50))
print(dl.removefirst())
dl.display()
print(dl.removelast())
dl.display()
print(dl.__len__())
print(dl.removeany(2))
dl.display()
print(dl._head._prev)
print(dl._tail._next)
dl.displayReverse()