import gmpy2
import sympy

FLAG = b'castorsCTF{1f_y0u_g07_th1s_w1th0u7_4ny_h1n7s_r3sp3c7}'
print(len(FLAG))

def gen_primes(n):
    return list(sympy.primerange(0, sympy.prime(n)+1))

def encrypt(m):
    k = gen_primes(len(m))
    return "".join([format(b * k[i] % 0x101, '02x') for i, b in enumerate(m)])

if __name__ == '__main__':
    enc = encrypt(FLAG)
    with open('enc.txt','w') as f:
        f.write(enc)