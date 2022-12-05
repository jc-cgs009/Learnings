class _Node:
    __slots__ = "_element", "_next"
    def __init__(self, element, next):
        self._element = element
        self._next = next

class singlyLinkedList:
    def __init__(self):
        self._head = None
        self._tail = None
        self._size = 0

    def __len__(self):
        return self._size
    
    def isempty(self):
        return self._size == 0

    def addlast(self, e):
        new = _Node(e, None)
        if self.isempty():
            self._head = new
            self._tail = new
        else:
            self._tail._next = new
            self._tail = new
        self._size += 1
    
    def addfirst(self, e):
        new = _Node(e, None)
        if self.isempty():
            self._head = new
            self._tail = new
        else:
            new._next = self._head
            self._head = new
        self._size += 1

    def addany(self, e, position):
        if position <= 0:
            self.addfirst(e)
        elif position >= self.__len__():
            self.addlast(e)
        else:
            new = _Node(e, None)
            p = self._head
            i = 1
            while i < position:
                p = p._next
                i += 1
            new._next = p._next
            p._next = new
            self._size += 1

    def removefirst(self):
        r = None
        if self.isempty():
            print("Linked list is empty!..")
        else:
            r = self._head._element
            self._head = self._head._next
            self._size -= 1

        if self._head == None:
            self._tail = None
        
        return r

    def removelast(self):
        r = None 
        if self.isempty():
            print("Linked list is empty!..")
        else:
            p = None
            i = 0
            while i < self.__len__()-1:
                if p:
                    p = p._next
                else:
                    p = self._head
                i += 1
            r = self._tail._element
            self._tail = p

            if self._tail != None:
                self._tail._next = None
            else:
                self._head = None
            
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
        
        return r


    def display(self):
        p = self._head
        while p:
            print(p._element, end = "-->")
            p = p._next
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


sl = singlyLinkedList()
sl.addlast(10)
sl.addlast(20)
sl.addlast(30)
sl.display()
sl.addfirst(40)
sl.addfirst(50)
sl.display()
sl.addany(100, 3)
sl.display()
print(sl.search(50))
print(sl.removefirst())
sl.display()
print(sl.removelast())
sl.display()
print(sl.__len__())
print(sl.removeany(2))
sl.display()