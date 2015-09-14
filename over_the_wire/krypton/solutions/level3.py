'''
http://www.simonsingh.net/The_Black_Chamber/hintsandtips.html

ENGLISH
Order Of Frequency Of Single Letters
E T A O I N S H R D L U
Order Of Frequency Of Digraphs
th er on an re he in ed nd ha at en es of or nt ea ti to it st io le is ou ar as de rt ve
Order Of Frequency Of Trigraphs
the and tha ent ion tio for nde has nce edt tis oft sth men
Order Of Frequency Of Most Common Doubles
ss ee tt ff ll mm oo
Order Of Frequency Of Initial Letters
T O A W B C D S F M R H I Y E G L N P U J K
Order Of Frequency Of Final Letters
E S T D N R Y F L O G H A K M P U W
One-Letter Words
a, I
Most Frequent Two-Letter Words
of, to, in, it, is, be, as, at, so, we, he, by, or, on, do, if, me, my, up, an, go, no, us, am
Most Frequent Three-Letter Words
the, and, for, are, but, not, you, all, any, can, had, her, was, one, our, out, day, get, has, him, his, how, man, new, now, old, see, two, way, who, boy, did, its, let, put, say, she, too, use
Most Frequent Four-Letter Words
that, with, have, this, will, your, from, they, know, want, been, good, much, some, time

'''

class Tr(object):
    
    def __init__(self, plain):
        self.plain = plain
        self.alphabet = list('ABCDEFGHIJKLMNOPQRSTUVWXYZ')
        self.key =      list('BOIHGKNQVTWYURXZAJEMSLDFPC')
        self.stat = [0] * 26
        
    def rot(self, c):
        if c in self.alphabet:
            return self.key[self.alphabet.index(c)]
        
        return c
        
    def encode(self):
        l = list(self.plain)
        m = map(self.rot,l)
        return ''.join(e for e in m)
        
    def statstics(self, input):
        l = list(input)
        for x in l:
            if x in self.alphabet:
                i = self.alphabet.index(x)
                self.stat[i] = self.stat[i] + 1

    def print_stat(self):
        for i in range(0,26):
            print self.alphabet[i] + "(" + str(self.stat[i]) + ")\t => " + self.key[i] 

    def print_none(self):
        for x in self.alphabet:
            if not x in self.key:
                print x

def main():
    problem = "KSVVW BGSJD SVSIS VXBMN YQUUK BNWCU ANMJS"
    t = Tr(problem)
    
    f = open("over_the_wire/krypton/krypton3/found1", "r")
    for line in f.readlines():
        t.statstics(line)
        #print line
        #print Tr(line).encode()
    f.close()
    
    f = open("over_the_wire/krypton/krypton3/found2", "r")
    for line in f.readlines():
        t.statstics(line)
        #print line
        #print Tr(line).encode()
    f.close()
    
    f = open("over_the_wire/krypton/krypton3/found3", "r")
    for line in f.readlines():
        t.statstics(line)
        #print line
        #print Tr(line).encode()
    f.close()
    
    #t.print_stat()
    #t.print_none()
    
    print problem
    print t.encode()

main()