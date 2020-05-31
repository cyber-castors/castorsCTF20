from pwn import remote
from base64 import b64decode

# HOST = ('localhost', 8080)
HOST = ('chals20.cybercastors.com', 14430)

def start(r):
    r.recvuntil(b'ready.\n')
    r.sendline(b'')

def solve(r):
    ct = r.recvline().rstrip(b'\n')
    for base in [2, 8, 16]:
        ct = "".join([chr(int(i, base)) for i in ct.split()])
    ans = b64decode(ct)
    r.sendline(ans)
    print(b'[+] Sending: ' + ans)
    print(r.recvline().rstrip(b'\n'))

if __name__ == '__main__':
    r = remote(*HOST)
    start(r)
    for i in range(50):
        solve(r)
    else:
        print(r.recv(1024))