import socketserver
import random
from threading import Timer
from secret import FLAG

BANNER = b"""
 ______ __      ______  ______       ______  ______  _____   ______    
/\  ___/\ \    /\  __ \/\  ___\     /\  ___\/\  __ \/\  __-./\  ___\   
\ \  __\ \ \___\ \  __ \ \ \__ \    \ \ \__ \ \ \/\ \ \ \/\ \ \___  \  
 \ \_\  \ \_____\ \_\ \_\ \_____\    \ \_____\ \_____\ \____-\/\_____\ 
  \/_/   \/_____/\/_/\/_/\/_____/     \/_____/\/_____/\/____/ \/_____/ 
                                                                                                                                                             
"""
MESSAGE = b"""
We have a small problem...
The flag gods are trying to send us a message, but our transmitter isn't calibrated to
decode it! If you can find the hamming distance for the following messages we may be
able to calibrate the transmitter in time. Entering the wrong distance will lock the
machine. Good luck, we'll only have 20 seconds!
Hit <enter> when ready.
"""

articles = ("a", "the")
adj = ("adorable", "clueless", "dirty", "odd", "fast")
nouns = ("puppy", "dragon", "rabbit", "girl", "monkey", "boy")
verbs = ("runs", "hits", "jumps", "drives", "sings to") 
adv = ("crazily.", "dutifully.", "foolishly.", "merrily.", "occasionally.")
sentence = [articles, adj, nouns, verbs, articles, nouns, adv]

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

def gen_rand_sentence():
    return " ".join(random.choice(word) for word in sentence).capitalize().encode()

def add_noise(m, state):
    ret = []
    for b in format(int.from_bytes(m,'big'),'b'):
        r = random.randint(1, 100)
        if r > state:
            ret.append(str(int(b) ^ 1))
        else:
            ret.append(b)
    return int("".join(ret),2).to_bytes(len(m),'big')

def gen_chal(state):
    target = gen_rand_sentence()
    corrupt = add_noise(target, state)
    return target, corrupt, str(format(int.from_bytes(target,'big') ^ int.from_bytes(corrupt,'big'),'b').count('1'))

def challenge(req):
    start, end = 20, 100
    t = Question_Timer(90)
    for state in range(start, end):
        target, corrupt, dist = gen_chal(state)

        req.sendall(f"The machine is currently {state}% calibrated.\n".encode())
        req.sendall(b"Transmitted message: " + target + b"\n")
        req.sendall(b"Received message: " + corrupt.hex().encode() + b"\n")
        req.sendall(b"Enter hamming distance: ")
        res = req.recv(1024).strip().decode()

        if not t.time_expired and dist == res:
            req.sendall(b"Correct answer!\n")
        elif dist == res:
            req.sendall(b"Time's up. Guess we'll never see the flag.")
            exit(0)
        else:
            req.sendall(f"The machine locked up. The correct distance was: {dist}".encode())
            exit(0)
    else:
        req.sendall(f"Machine calibration successful! Here's the transmitted flag: {FLAG}".encode())
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