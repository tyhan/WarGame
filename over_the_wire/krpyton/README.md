##WriteUp
#Over the wire
[Over the wire](http://overthewire.org/wargames/)

##Krpyton
[krpyton](http://overthewire.org/wargames/krypton/)

### Level 0

####Problem

decode base64:
<code>S1JZUFRPTklTR1JFQVQ=</code>

#### Solution

level0.py

<code>#!/usr/local/bin/python
print "S1JZUFRPTklTR1JFQVQ=".decode("base64")
</code>

#### Flag

<code> KRYPTONISGREAT </code>

### Level 1

#### Problem



<code>YRIRY GJB CNFFJBEQ EBGGRA</code>

#### Solution

In README
"It is 'encrypted' using a simple rotation called ROT13."

<code>#!/usr/local/bin/python
"YRIRY GJB CNFFJBEQ EBGGRA".encode('ROT13')
</code>

#### Flag

<code>LEVEL TWO PASSWORD ROTTEN</code>


