from pwn import *

flag = [0x64, 0x35, 0x68, 0x35, 0x64, 0x37, 0x33, 0x7a, 0x38, 0x6b, 0x33, 0x37, 0x6b, 0x72, 0x67, 0x7a]

def main():
    password = ''
    # TODO Change ip address
    conn = remote("localhost", 8080)
    output = rev_add(rev_rot(flag))

    for c in output:
        password += chr(c)

    print(f'Sending password: {password}')

    conn.sendline(password)
    conn.recvline()
    result = conn.recv(4096)

    print(result)

    conn.close()

def rev_add(flag):
    for i, _ in enumerate(flag):
        flag[i] -= 2
    return flag

def rev_rot(flag):
    for i, c in enumerate(flag):
        if c >= 0x61 and c <= 0x7A:
            flag[i] = (((flag[i] - 0x61) - 0xA) % 26) + 0x61
    return flag


if __name__ == '__main__':
    main()