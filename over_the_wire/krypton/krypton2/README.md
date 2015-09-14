
#Level 2

## Problem
```
OMQEMDUEQMEK
```

## Solution

Caesar

```python
class Caesar(object):
    
    def __init__(self, plain):
        self.plain = plain
        self.alphabet_L = list('abcdefghijklmnopqrstuvwxyz')
        self.alphabet_U = list('ABCDEFGHIJKLMNOPQRSTUVWXYZ')
        
    def rot(self, c):
        if c in self.alphabet_L:
            return self.alphabet_L[(self.alphabet_L.index(c)+self.index)%(len(self.alphabet_U))]
        if c in self.alphabet_U:
            return self.alphabet_U[(self.alphabet_U.index(c)+self.index)%(len(self.alphabet_L))]
        
        return c
        
    def encode(self, rot):
        self.index = rot
        l = list(self.plain)
        m = map(self.rot,l)    
        return ''.join(e for e in m)

def main():
    m = "OMQEMDUEQMEK"
    for x in range(1, 26):
        print str(x) + ": " + Caesar(m).encode(x) +"\n"

main()
```

## Flag

```
CAESARISEASY
```
