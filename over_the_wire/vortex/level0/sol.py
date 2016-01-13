import pwn
import re

addr = "vortex.labs.overthewire.org"
port = "5842"

r = pwn.remote(addr,port)

num = []
for i in range(4):
    num += [pwn.u32(r.recvn(4))]

r.send(pwn.p32(sum(num) & 0xFFFFFFFF))

rcv = r.recvall()
passwd =  re.findall("Password: (.*)", rcv)

print passwd[0]