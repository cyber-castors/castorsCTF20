from gmpy2 import invert
import sympy

def gen_primes(n):
    return list(sympy.primerange(0, sympy.prime(n)+1))

def decrypt(m):
    n = len(bytes.fromhex(m))
    k = gen_primes(n)
    k_inv = [invert(i, 0x101) for i in k]
    dec = []
    for i, j in enumerate(range(0, len(m), 2)):
        temp = int(m[j:j+2], 16) * k_inv[i] % 0x101
        dec.append(int(temp))
    return bytes(dec)

if __name__ == '__main__':
    with open('enc.txt','r') as f:
        enc = f.read()
    print(len(enc))
    dec = decrypt(enc)
    print(dec)