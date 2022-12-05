class _Node:
    __slots__ = "_element", "_next"
    def __init__(self, element, next):
        self._element = element
        self._next = next

class circularLinkedList:
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
        else:
            self._tail._next = new
        self._tail = new
        self._tail._next = self._head
        self._size += 1
    
    def addfirst(self, e):
        new = _Node(e, None)
        if self.isempty():
            self._head = new
            self._tail = new
        else:
            new._next = self._head
            self._head = new
        self._tail._next = self._head
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
            self._tail._next = self._head
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
                self._tail._next = self._head
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
            self._size -= 1
        return r


    def display(self):
        p = self._head
        i = 0
        while i < self.__len__():
            print(p._element, end = "-->")
            p = p._next
            i += 1
        print()

    def search(self, key):
        p = self._head
        index = 0
        i = 0
        while i < self.__len__():
            if p._element == key:
                return index
            p = p._next
            index += 1
            i += 1
        return -1


cl = circularLinkedList()
cl.addlast(10)
cl.addlast(20)
cl.addlast(30)
cl.display()
cl.addfirst(40)
cl.addfirst(50)
cl.display()
cl.addany(100, 3)
cl.display()
print(cl.search(50))
print(cl.removefirst())
cl.display()
print(cl.removelast())
cl.display()
print(cl.__len__())
print(cl.removeany(2))
cl.display()
print(cl._head)
print(cl._tail._next)