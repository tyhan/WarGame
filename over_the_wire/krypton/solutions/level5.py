from pygenere import Vigenere, VigCrack

        
def main():
    problem = "BELOSZ"
    
    f = open("over_the_wire/krypton/krypton5/found1", "r")
    key1 = VigCrack(Vigenere(f.readline())).crack_codeword()
    print key1
    f.close()
    
    f = open("over_the_wire/krypton/krypton5/found2", "r")
    key2 = VigCrack(Vigenere(f.readline())).crack_codeword()
    print key2
    f.close()
    
    print Vigenere(problem).decipher(key1)
    print Vigenere(problem).decipher(key2)

main()