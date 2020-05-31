import sympy

FLAG = b'castorsCTF{N0_End_T0d4y_F0r_L0v3_I5_X3n0n}'
print(len(FLAG))

def gen_primes(n):
    return list(sympy.primerange(0, sympy.prime(n)+1))

def encrypt(m):
    k = gen_primes(len(m))
    return "".join([str(b * k[i]) for i, b in enumerate(m)])

if __name__ == '__main__':
    enc = encrypt(FLAG)
    with open('enc.txt','w') as f:
        f.write(enc)