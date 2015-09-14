class Tr6(object):
    
    def __init__(self, plain):
        self.plain = plain
        self.alphabet = list('ABCDEFGHIJKLMNOPQRSTUVWXYZ')
        self.key =      [5, 17, 4, 10, 4, 24]
        
    def rot(self, c):
        if c in self.alphabet:
            return self.key[self.alphabet.index(c)]
        
        return c
        
    def encode(self):
        k = 0
        l = list(self.plain)
        m = []
        for c in l:
            if c in self.alphabet:
                c = self.alphabet[(self.alphabet.index(c) + self.key[k])%26]
                k = (k + 1) % 6
            m = m + [c]
        return ''.join(e for e in m)

    def decode(self):
        k = 0
        l = list(self.plain)
        m = []
        for c in l:
            if c in self.alphabet:
                c = self.alphabet[(self.alphabet.index(c) + (26 - self.key[k]))%26]
                k = (k + 1) % 6
            m = m + [c]
        return ''.join(e for e in m)
        
def main():
    problem = "HCIKV RJOX"
    t = Tr6(problem)
    
    f = open("over_the_wire/krypton/krypton4/found1", "r")
    print Tr6(f.readline()).decode()
    f.close()
    
    f = open("over_the_wire/krypton/krypton4/found2", "r")
    print Tr6(f.readline()).decode()
    f.close()
    
    print problem
    print t.decode()

main()