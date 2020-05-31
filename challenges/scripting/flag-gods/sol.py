from pwn import remote

HOST = ('chals20.cybercastors.com', 14431)

def start(r):
    r.recvuntil(b'ready.\n')
    r.sendline(b'')

def solve(r):
    r.recvline()
    target = r.recvline().strip().split(b' ', 2)[2]
    corrupt = r.recvline().strip().split(b' ', 2)[2]
    dist = str(format(int.from_bytes(target,'big') ^ int.from_bytes(bytes.fromhex(corrupt.decode()),'big'),'b').count('1'))
    r.recv(1024)
    r.send(dist.encode())
    print(r.recvline())

if __name__ == '__main__':
    r = remote(*HOST)
    start(r)
    for _ in range(80):
        solve(r)
    else:
        print(r.recv(1024))
