import socketserver
import random
from threading import Timer
from secret import FLAG

# Globals
BANNER = b"""
------------------Welcome to Beginner Arithmetics!------------------
To get the flag you'll have to solve a series of arithmetic challenges.
The problems may be addition, substraction, multiplication or integer division.
Numbers can range from 1 to 9. Easy right?
You'll have very little time to answer so do your best!
Hit <enter> when ready.
"""

int2str = [
    ('0', 'zero'),
    ('1', 'one'),
    ('2', 'two'),
    ('3', 'three'),
    ('4', 'four'),
    ('5', 'five'),
    ('6', 'six'),
    ('7', 'seven'),
    ('8', 'eight'),
    ('9', 'nine')]

sym2str = [
    ('+', 'plus'),
    ('-', 'minus'),
    ('*', 'multiplied-by'),
    ('//', 'divided-by')]

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

def level1(a, b, op):
    return " ".join([int2str[a][0], sym2str[op][0], int2str[b][0]])

def level2(a, b, op):
    r = random.randint(0,1)
    return " ".join([int2str[a][r], sym2str[op][0], int2str[b][r]])

def level3(a, b, op):
    r = random.randint(0,1)
    return " ".join([int2str[a][r], sym2str[op][r], int2str[b][r]])

def level4(a, b, op):
    r1, r2, r3 = [random.randint(0,1) for i in range(3)]
    return " ".join([int2str[a][r1], sym2str[op][r2], int2str[b][r3]])

def gen_chal(n):
    a, b, op = random.randint(1,9), random.randint(1,9), random.randint(0,3)

    if   op == 0: ans = a + b
    elif op == 1: ans = a - b
    elif op == 2: ans = a * b
    elif op == 3: ans = a // b

    if   n >= 75: eq = level4(a, b, op)
    elif n >= 50: eq = level3(a, b, op)
    elif n >= 25: eq = level2(a, b, op)
    elif n >=  0: eq = level1(a, b, op)

    return eq, str(ans)

def challenge(req):
    for n in range(100):
        eq, ans = gen_chal(n)

        t = Question_Timer(1)
        req.sendall(b"What is " + eq.encode() + b" ?\n")
        res = req.recv(1024).strip()

        if not t.time_expired and ans == res.decode():
            req.sendall(b"Correct answer!\n")
        elif ans == res.decode():
            req.sendall(b"Too slow!")
            exit(0)
        else:
            req.sendall(b"Wrong answer! The correct answer was " + str(ans).encode())
            exit(0)
    else:
        req.sendall(b"Wow! You're fast! Here's your flag: " + FLAG)
        exit(0)

class TaskHandler(socketserver.BaseRequestHandler):
    def handle(self):
        self.request.sendall(BANNER)
        self.request.recv(1024).strip()
        challenge(self.request)

if __name__ == '__main__':
    socketserver.ThreadingTCPServer.allow_reuse_address = True
    server = socketserver.ThreadingTCPServer(('0.0.0.0', 8080), TaskHandler)
    server.serve_forever()
