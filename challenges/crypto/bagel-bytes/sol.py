from pwn import remote
from binascii import hexlify, unhexlify
import string

HOST = ('chals20.cybercastors.com', 14420)
sigma = string.printable

def bake_flag(r):
    r.recvuntil(b"choice: ")
    r.sendline(b'2')
    r.recvuntil(b"> ")

def get_pad(m, s):
    return '*' * (s - len(m) - 1)

def get_res(r, p, pos=False):
    r.sendline(p.encode())
    l = r.recvuntil(b"bytes:\n")
    if pos:
        print("\n{}".format(l.decode().strip('\n')))
    return unhexlify(r.recvline().strip())[s-16:s]

def solve(r, m, s):
    bake_flag(r)
    payload = get_pad(m, s)
    target = get_res(r, payload, False)
    for ch in sigma:
        bake_flag(r)
        payload = get_pad(m, s) + m + ch
        res = get_res(r, payload)
        if res == target:
            return m + ch
    else:
        print("Mmmm crispy bagels!")
        exit(0)
        
if __name__ == "__main__":
    r = remote(*HOST)
    m = ''
    s = 48
    while True:
        m = solve(r, m, s)
        print("Found: {}".format(m))