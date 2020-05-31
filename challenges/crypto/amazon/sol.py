import sympy

def gen_primes(n):
    return list(sympy.primerange(0, sympy.prime(n)+1))

def decrypt(m):
    k = 2 # First prime
    n = 0 # current key index
    c = 0 # current index
    dec = []
    for i in range(len(enc)):
        temp = int(enc[c:i+1]) / k
        if temp == int(temp) and temp > 31:
            k = sympy.nextprime(k)
            c = i + 1
            dec.append(int(temp))
    return bytes(dec)


if __name__ == '__main__':
    with open('enc.txt','r') as f:
        enc = f.read()
    dec = decrypt(enc)
    print(dec)