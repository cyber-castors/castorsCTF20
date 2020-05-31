import socketserver
from collections import OrderedDict
from textwrap import wrap
import numpy as np
import string
from secret import FLAG, KEYS

BANNER = b"""
 __    __  ______  ______  __  ______       ______  ______  __  __  ______  ______  __           ______  __  __  ______    
/\ "-./  \/\  __ \/\  ___\/\ \/\  ___\     /\  ___\/\  ___\/\ \_\ \/\  __ \/\  __ \/\ \         /\  == \/\ \/\ \/\  ___\   
\ \ \-./\ \ \  __ \ \ \__ \ \ \ \ \____    \ \___  \ \ \___\ \  __ \ \ \/\ \ \ \/\ \ \ \____    \ \  __<\ \ \_\ \ \___  \  
 \ \_\ \ \_\ \_\ \_\ \_____\ \_\ \_____\    \/\_____\ \_____\ \_\ \_\ \_____\ \_____\ \_____\    \ \_____\ \_____\/\_____\ 
  \/_/  \/_/\/_/\/_/\/_____/\/_/\/_____/     \/_____/\/_____/\/_/\/_/\/_____/\/_____/\/_____/     \/_____/\/_____/\/_____/ 
                                                                                                                           
"""
MESSAGE = b"""
All right kids!
The Magic Flag Bus is about to leave!
Make sure to fill each row as you step in.
Wait, what! All the kids got moved around.
I *knew* I should've stayed home today!
Nevermind... Load your own bus and let's leave.
"""
MENU = b"""
Select:
    1) Load magic school bus
    2) View magic flag bus
Your choice: """
PAD='X'

def elim_key_redundancy(keys):
    return [''.join(OrderedDict.fromkeys(key)) for key in keys]

def filter(msg):
    sigma = string.punctuation
    for ch in sigma:
        msg = msg.replace(ch, '')
    return ''.join(msg.split())

def pad(msg, key):
    if len(msg) % len(key) != 0:
        msg +=PAD * (len(key) - len(msg) % len(key))
    return msg

def encrypt(msg,keys):
    keys = elim_key_redundancy(keys)
    for key in keys:
        if key == '':
            break
        key_ord = sorted(key)
        tp_key=[key.index(i) for i in key_ord]
        tp_msg = np.array([list(i) for i in wrap(pad(filter(msg), key),len(key))])
        enc_msg = ''
        for i in tp_key:
            enc_msg += ''.join(tp_msg[:,i])
        msg = enc_msg
    return enc_msg.replace(PAD,'').upper()

def challenge(req):
    while True:
        req.sendall(MENU)
        choice = req.recv(1024).strip()
        if choice == b'1':
            req.sendall(b"\nWho's riding the bus?: ")
            inp = req.recv(1024).strip().decode()
            ct = encrypt(inp, KEYS)
            req.sendall(f"Bus seating: {ct}\n".encode())
        elif choice == b'2':
            ct = encrypt(FLAG, KEYS)
            req.sendall(f"\nFlag bus seating: {ct}\n".encode())

class TaskHandler(socketserver.BaseRequestHandler):
    def handle(self):
        self.request.sendall(BANNER)
        self.request.sendall(MESSAGE)
        challenge(self.request)

if __name__ == '__main__':
    socketserver.ThreadingTCPServer.allow_reuse_address = True
    server = socketserver.ThreadingTCPServer(('0.0.0.0', 8080), TaskHandler)
    server.serve_forever()
    