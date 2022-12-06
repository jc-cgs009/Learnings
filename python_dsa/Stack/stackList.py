class stackList:
    def __init__(self):
        self._stack = []

    def length(self):
        return len(self._stack)
    
    def isempty(self):
        return len(self._stack)

    def push(self, e):
        self._stack.append(e)

    def pop(self):
        e = None
        if self.isempty():
            print("Stack is empty!...")
        else:
            e = self._stack.pop()
        return e

    def top(self):
        return self._stack[-1]

    def display(self):
        for i in self._stack:
            print(i, end = "-->")
        print()

sl = stackList()
sl.push(10)
sl.push(20)
sl.push(30)
sl.push(40)
sl.push(50)
sl.display()
print(sl.top())
print(sl.pop())
sl.display()
print(sl.pop())
sl.display()


