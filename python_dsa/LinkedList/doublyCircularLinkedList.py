class _Node:
    __slots__ = "_element", "_next", "_prev"
    def __init__(self,elememt, next, prev):
        self._element = elememt
        self._next = next
        self._prev = prev

class doublyCircularLinkedList:
    def __init__(self):
        self._head = None
        self._tail = None
        self._size = 0

    def __len__(self):
        return self._size

    def isempty(self):
        return self._size == 0

    def addlast(self, e):
        new = _Node(e, None,None)
        if self.isempty():
            self._head = new
            self._tail = new
        else:
            new._prev = self._tail
            self._tail._next = new
            self._tail = new
        self._head._prev = self._tail
        self._tail._next = self._head
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
        self._head._prev = self._tail
        self._tail._next = self._head
        self._size += 1
    
    def addany(self, e, position):
        if position <= 0:
            self.addfirst(e)
        elif position >= len(self):
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

    def removelast(self):
        r = None
        if self.isempty():
            print("Linked List is empty!...")
        else:
            r = self._tail._element
            self._tail = self._tail._prev
            self._tail._next = self._head
            self._head._prev = self._tail
            self._size -= 1
        
        if self.isempty():
            self._head = None
            self._tail = None
        return r

    def removeany(self, position):
        r = None
        if position < 0 or position >= len(self):
            print("List index out of range!...")
        elif position == 0:
            r = self.removefirst()
        elif position == len(self)-1:
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

    def removefirst(self):
        r = None
        if self.isempty():
            print("List is empty!...")
        else:
            r = self._head._element
            self._head = self._head._next
            self._head._prev = self._tail
            self._tail._next = self._head
            self._size -= 1 

        if self.isempty():
            self._head = None
            self._tail = None

        return r

    def display(self):
        p = self._head
        i = 0
        while i < self.__len__():
            print(p._element, end="-->")
            p = p._next
            i += 1
        print()

    def displayreverse(self):
        p = self._tail
        i = 0
        while i < self.__len__():
            print(p._element, end="<--")
            p = p._prev
            i += 1
        print()
    

dcl = doublyCircularLinkedList()
dcl.addlast(10)
dcl.addlast(20)
dcl.addlast(30)
dcl.display()
dcl.addfirst(40)
dcl.addfirst(50)
dcl.addfirst(60)
dcl.display()
# dcl.addany(100, len(dcl))
dcl.addany(100, 1)
dcl.display()
dcl.displayreverse()
print(dcl.removelast())
dcl.display()
dcl.displayreverse()
print(dcl.removefirst())
dcl.display()
dcl.displayreverse()
# print(dcl.removeany(len(dcl)-1))
print(dcl.removeany(2))
dcl.display()
dcl.displayreverse()
print(dcl._head._prev)
print(dcl._tail)
print("*"*10)
print(dcl._tail._next)
print(dcl._head)
        