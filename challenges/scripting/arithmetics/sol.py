from pwn import remote

HOST = ('chals20.cybercastors.com', 14429)

n2s = {
    '1': 'one',
    '2': 'two',
    '3': 'three',
    '4': 'four',
    '5': 'five',
    '6': 'six',
    '7': 'seven',
    '8': 'eight',
    '9': 'nine',}
s2n = {v: k for k, v in n2s.items()}

def start(r):
    r.recv(4096)
    r.sendline(b'')

def solve(r):
    res = r.recvline().rstrip(b'\n')
    print(res)
    a, op, b = res.decode().split()[-4:-1]
    ans = parse(a, b, op)
    r.sendline(ans)
    print(r.recvline().rstrip(b'\n'))

def parse(a, b, op):
    try:
        n2s[a]
        a = int(a)
    except:
        a = int(s2n[a])
    try:
        n2s[b]
        b = int(b)
    except:
        b = int(s2n[b])

    if   op in ['+', 'plus']:           ans = a + b
    elif op in ['-', 'minus']:          ans = a - b
    elif op in ['*', 'multiplied-by']:  ans = a * b
    elif op in ['//', 'divided-by']:    ans = a // b

    return str(ans).encode()

if __name__ == '__main__':
    r = remote(*HOST)
    start(r)
    for i in range(100):
        solve(r)
    else:
        print(r.recv(1024))
