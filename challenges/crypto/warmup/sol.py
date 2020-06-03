from Crypto.PublicKey import RSA
from Crypto.Util.number import bytes_to_long, long_to_bytes
from gmpy2 import isqrt, invert

flag = b'castorsCTF{n0th1ng_l1k3_pr1m3_numb3r5_t0_w4rm_up_7h3_3ng1n3s}'
print(int.from_bytes(flag, 'big'))
with open('rsa.key','r') as f:
    data = f.read()

key = RSA.import_key(data)
n, e, d, p, q = key.n, key.e, key.d, key.p, key.q

if q > p: p, q = q, p

# Given
h = 2 * (p**2 + q**2)
a = (p**2 - q**2) // 2

# Found with Given
p_sol = isqrt(a + h // 4)
q_sol = isqrt(p_sol**2 - 2*a)
n_sol = p_sol*q_sol
d_sol = invert(e, (p_sol-1)*(q_sol-1))

# Assert
assert n == n_sol
assert d == d_sol
assert p == p_sol
assert q == q_sol

# Gen chal
c = pow(bytes_to_long(flag), e, n_sol)
print(c)
print(long_to_bytes(pow(c,d,n)))

w_data = f"a=p+q\nb=p-q\nc={h}\nA={a}\ne=0x10001\nenc={c}"
with open('chal.txt','w') as f:
    f.write(w_data)
