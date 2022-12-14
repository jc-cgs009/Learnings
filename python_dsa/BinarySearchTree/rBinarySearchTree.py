class _Node:
    __slots__ = '_element', '_left', '_right'
    def __init__(self, element , left = None, right = None):
        self._element = element
        self._left = left
        self._right = right

class binarySearchTree:
    def __init__(self):
        self._root = None

    def rinsert(self, troot, e):
        if troot:
            if e < troot._element:
                troot._left = self.rinsert(troot._left, e)
            elif e > troot._element:
                troot._right = self.rinsert(troot._right, e)
        else:
            n = _Node(e)
            troot = n
        return troot
    
    def rsearch(self, troot, key):
        if troot:
            if key == troot._element:
                return True
            elif key < troot._element:
                return self.rsearch(troot._left, key)
            elif key > troot._element:
                return self.rsearch(troot._right, key)
        return False

    def delete(self, e):
        p = self._root
        pp = None
        while p and p._element != e:
            pp = p
            if e < p._element:
                p = p._left
            else:
                p = p._right

        if not p:
            return False

        if p._left and p._right:
            s = p._left
            ps = p
            while s._right:
                ps = s
                s = s._right
            p._element = s._element
            p = s
            pp = ps

        c = None
        if p._left:
            c = p._left
        else:
            c = p._right

        if p == self._root:
            self._root = None
        else:
            if p == pp._left:
                pp._left = c
            else:
                pp._right = c

        


    def inorder(self, troot):
        if troot:
            self.inorder(troot._left)
            print(troot._element, end = ' ')
            self.inorder(troot._right)

    def preorder(self, troot):
        if troot:
            print(troot._element, end = ' ')
            self.preorder(troot._left)
            self.preorder(troot._right)

    def postorder(self, troot):
        if troot:
            self.postorder(troot._left)
            self.postorder(troot._right)
            print(troot._element, end = ' ')

    def levelorder(self, troot):
        Q = []
        t = troot
        print(t._element, end = ' ')
        Q.append(t)
        while Q:
            t = Q.pop(0)
            if t._left:
                print(t._left._element, end = ' ')
                Q.append(t._left)
            if t._right:
                print(t._right._element, end = ' ')
                Q.append(t._right)


B = binarySearchTree()
B._root = B.rinsert(B._root, 50)
B.rinsert(B._root, 30)
B.rinsert(B._root, 80)
B.rinsert(B._root, 10)
B.rinsert(B._root, 40)
B.rinsert(B._root, 60)
B.rinsert(B._root, 90)
B.rinsert(B._root, 10) # Binary search doesn't support duplicate. 
print("Inorder treaversal")
B.inorder(B._root)
print()
print("preorder treaversal")
B.preorder(B._root)
print()
print("postorder treaversal")
B.postorder(B._root)
print()
print("levelorder treaversal")
B.levelorder(B._root)
print()
print(B.rsearch(B._root, 40))
print(B.rsearch(B._root, 100))
B.delete(10)
print("Inorder traversal")
B.inorder(B._root)

