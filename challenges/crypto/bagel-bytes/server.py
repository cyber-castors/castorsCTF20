from Crypto.Cipher import AES
import socketserver
from secret import FLAG, KEY

BANNER = b"""
 ______  ______  ______  ______  __           ______  __  __  ______ ______  ______    
/\  == \/\  __ \/\  ___\/\  ___\/\ \         /\  == \/\ \_\ \/\__  _/\  ___\/\  ___\   
\ \  __<\ \  __ \ \ \__ \ \  __\\\ \ \____    \ \  __<\ \____ \/_/\ \\\ \  __\\\ \___  \  
 \ \_____\ \_\ \_\ \_____\ \_____\ \_____\    \ \_____\/\_____\ \ \_\\\ \_____\/\_____\ 
  \/_____/\/_/\/_/\/_____/\/_____/\/_____/     \/_____/\/_____/  \/_/ \/_____/\/_____/ 
                                                                                       """
MESSAGE = b"""
Welcome to Bagel Bytes!
Our ovens are known for baking ExtraCrispyBagels!
We bake 16 bagels per rack and the last rack is always full!
Today's special is flag bagels!
"""

MENU = b"""
Select:
    1) Bake your own bagels!
    2) Bake the flag!
Your choice: """

aes = AES.new(KEY, AES.MODE_ECB)

def pad(s):
    if len(s) % 16 == 0:
        return s
    else:
        pad_b = 16 - len(s) % 16
        return s + bytes([pad_b]) * pad_b

def bake_your_own(s, c):
    return c.encrypt(pad(s))

def bake_flag(s, c):
    return len(s), c.encrypt(pad(s + FLAG))

def challenge(req):
    while True:
        req.sendall(MENU)
        choice = req.recv(1024).strip()
        if choice == b'1':
            req.sendall(b'Add your bagels:\n> ')
            inp = req.recv(1024).strip()
            enc = bake_your_own(inp, aes)
            req.sendall(b"\nThank you for baking with us! Here are your baked bytes:\n" + enc.hex().encode() + b"\n")
        elif choice == b'2':
            req.sendall(b"Add your bagels:\n> ")
            inp = req.recv(1024).strip()
            n, enc = bake_flag(inp, aes)
            req.sendall(b"\nThank you for baking with us! Here are your baked bytes:\n" + enc.hex().encode() + b"\n")
        else:
            req.sendall(b"\nThat's not on the menu!\n")

class TaskHandler(socketserver.BaseRequestHandler):
    def handle(self):
        self.request.sendall(BANNER)
        self.request.sendall(MESSAGE)
        challenge(self.request)

if __name__ == '__main__':
    socketserver.ThreadingTCPServer.allow_reuse_address = True
    server = socketserver.ThreadingTCPServer(('0.0.0.0', 8080), TaskHandler)
    server.serve_forever()