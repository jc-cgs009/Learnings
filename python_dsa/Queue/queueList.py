class queueList:
    def __init__(self):
        self._queue = []

    def __len__(self):
        return len(self._queue)

    def isempty(self):
        return len(self) == 0

    def enqueue(self, e):
        self._queue.append(e)

    def dequeue(self):
        r = None
        if self.isempty():
            print("Queue is empty!...")
        else:
            r = self._queue.pop(0)

        return r

    def first(self):
        r = None
        if self.isempty():
            print("Queue is empty!...")
        else:
            r = self._queue[0]
        return r

    def display(self):
        for i in self._queue:
            print(i, end = "<--")
        print()

ql = queueList()
ql.enqueue(10)
ql.enqueue(20)
ql.enqueue(30)
ql.enqueue(40)
ql.enqueue(50)
print(ql.first())
ql.display()
print(ql.dequeue())
print(ql.dequeue())
ql.display()


