import socketserver
import random, string
from base64 import b64encode
from threading import Timer
from secret import FLAG

BANNER = b"""
 ______  ______  ______  ______       ______  __  __  __   __  __   __  ______  ______    
/\  == \/\  __ \/\  ___\/\  ___\     /\  == \/\ \/\ \/\ "-.\ \/\ "-.\ \/\  ___\/\  == \   
\ \  __<\ \  __ \ \___  \ \  __\     \ \  __<\ \ \_\ \ \ \-.  \ \ \-.  \ \  __\\\ \  __<   
 \ \_____\ \_\ \_\/\_____\ \_____\    \ \_\ \_\ \_____\ \_\\\ \_\ \_\\\ \_\ \_____\ \_\ \_\ 
  \/_____/\/_/\/_/\/_____/\/_____/     \/_/ /_/\/_____/\/_/ \/_/\/_/ \/_/\/_____/\/_/ /_/ 
                                                                                          
"""
MESSAGE = b"""
Listen up!
We're down 0 - 49 with 2 outs on the final inning but we got this!
Don't worry about getting hits, just run those bases as fast as you can.
They have The Flash fielding for their team!
Hit <enter> when ready.
"""

class Question_Timer(object):
    def __init__(self, interval):
        self.time_expired = False
        self.interval = interval
        self.start_timer()

    def timeout(self):
        self.time_expired = True

    def start_timer(self):
        t = Timer(self.interval, self.timeout)
        t.daemon == True
        t.start()

def gen_random_word(length):
   letters = string.ascii_letters
   return ''.join(random.choice(letters) for i in range(length))

def gen_chal():
    target = b'castorsCTF{' + gen_random_word(random.randint(10, 20)).encode() + b'}'
    enc = b64encode(target)
    for base in ['02x', 'o', '08b']:
        enc = " ".join([format(i, base) for i in enc]).encode()
    
    return enc, target

def challenge(req):
    for i in range(50):
        enc, ans = gen_chal()

        t = Question_Timer(3)
        req.sendall(enc + b"\n")
        res = req.recv(1024).strip()

        if not t.time_expired and ans == res:
            req.sendall(b"Correct answer!\n")
        elif ans == res:
            req.sendall(b"Ah too slow! Gotta beat The Flash!")
            exit(0)
        else:
            req.sendall(b"3 Outs... there's always next season.")
            exit(0)
    else:
        req.sendall(b"Wow! The Flash is impressed! Here's your flag: " + FLAG)
        exit(0)

class TaskHandler(socketserver.BaseRequestHandler):
    def handle(self):
        self.request.sendall(BANNER)
        self.request.sendall(MESSAGE)
        self.request.recv(1024).strip()
        challenge(self.request)

if __name__ == '__main__':
    socketserver.ThreadingTCPServer.allow_reuse_address = True
    server = socketserver.ThreadingTCPServer(('0.0.0.0', 8080), TaskHandler)
    server.serve_forever()
