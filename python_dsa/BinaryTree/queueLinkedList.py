class _Node:
    __slots__ = "_element", "_next"
    def __init__(self, element, next):
        self._element = element
        self._next = next

class queueLinkedList:
    def __init__(self):
        self._front = None
        self._rare = None
        self._size = 0

    def __len__(self):
        return self._size

    def isempty(self):
        return len(self) == 0

    def enqueue(self, e):
        new = _Node(e, None)
        if self._front == None:
            self._front = new
            self._rare = new
        else:
            self._rare._next = new
            self._rare = new
        self._size += 1

    def dequeue(self):
        r = None
        if self.isempty():
            print("Queue is empty!...")
        else:
            r = self._front._element
            self._front = self._front._next
            self._size -= 1
        return r

    def first(self):
        e = None
        if self.isempty():
            print("Queue is empty!...")
        else:
            e  = self._front._element
        return e

    def display(self):
        p = self._front
        while p:
            print(p._element, end ="<--" )
            p = p._next
        print()

if __name__ == "__main__":
    qll = queueLinkedList()
    qll.enqueue(10)
    qll.enqueue(20)
    qll.enqueue(30)
    qll.enqueue(40)
    qll.enqueue(50)
    print(qll.first())
    qll.display()
    print(qll.dequeue())
    print(qll.dequeue())
    qll.display()
