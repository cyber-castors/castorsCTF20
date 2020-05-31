# This challenge doesn't need a solver.
# Find the length of the pad and xor the input with the output.

from pwn import *
import string

HOST = ('chals20.cybercastors.com', 14422)
sigma = string.printable
target = b'CASTORSctf[K\x13\x13P\x7fY\x10UR\x7fK\x13Y\x15\x7fS\x13CR\x13\x17\x7f\x14ND\x7fD\x10N\x17\x7fR\x13US\x13\x7fTH\x13M\x01]'

def get(r, m):
  r.recvuntil(b'> ')
  r.sendline(m)
  return r.recvline().decode().rstrip('\n')

def main():
  r = remote(*HOST)
  m = ''
  for i in range(len(m), len(target)):
    for ch in sigma:
      payload = m + ch
      res = get(r, payload)
      if res == "b''":
        m += ch
        print('Found: {0}'.format(m))
        break
    else:
      print('No match found')

main()