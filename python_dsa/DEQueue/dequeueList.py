class dequeueList:
    def __init__(self):
        self._data = []

    def __len__(self):
        return len(self._data)

    def isempty(self):
        return len(self) == 0

    def addfirst(self, e):
        self._data.insert(0, e)

    def addlast(self, e):
        self._data.append(e)

    def removefirst(self):
        r = None
        if self.isempty():
            print("DEQueue is empty!...")
        else:
            r = self._data.pop(0)
        return r

    def removelast(self):
        r = None
        if self.isempty():
            print("DEQueue is empty!...")
        else:
            r = self._data.pop()
        return r

    def first(self):
        e = None
        if self.isempty():
            print("DEQueue is empty!...")
        else:
            e = self._data[0]
        return e

    def last(self):
        e = None
        if self.isempty():
            print("DEQueue is empty!...")
        else:
            e = self._data[-1]
        return e

dql = dequeueList()
dql.addfirst(10)
dql.addfirst(20)
dql.addfirst(30)
print(dql._data)
dql.addlast(15)
dql.addlast(25)
dql.addlast(35)
print(dql._data)
print(dql.first())
print(dql.last())
print(dql.removefirst())
print(dql.removelast())
print(dql._data)
print(dql.first())
print(dql.last())



