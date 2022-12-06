class _Node:
    __slots__ = "_element", "_next"
    def __init__(self, element, next):
        self._element = element
        self._next = next

class stackLinkedList:
    def __init__(self):
        self._top = None
        self._size = 0

    def __len__(self):
        return self._size

    def isempty(self):
        return len(self) == 0

    def push(self, e):
        new = _Node(e, None)
        if self.isempty():
            self._top = new
        else:
            new._next = self._top
            self._top = new
        self._size += 1

    def pop(self):
        if self.isempty():
            print("Stack is empty!...")
        else:
            r = self._top._element
            self._top = self._top._next
            self._size -= 1

        return r

    def top(self):
        return self._top._element

    def display(self):
        p = self._top
        while p:
            print(p._element, end = "-->")
            p = p._next
        print()

sll = stackLinkedList()
sll.push(10)
sll.push(20)
sll.push(30)
sll.push(40)
sll.push(50)
sll.push(60)
print(sll.top())
sll.display()
print(sll.pop())
sll.display()

