import socketserver
from secret import FLAG

def pad(s):
    if len(s) > len(FLAG):
        return ((len(s)//len(FLAG) + 1) * FLAG)[:len(s)]
    else:
        return FLAG[:len(s)]

def otp(s):
    otp_s = pad(s)
    return bytes([byte ^ s[i] for i, byte in enumerate(otp_s)]).replace(b'\x00', b'')

def challenge(req):
    while True:
        req.sendall(b"> ")
        inp = req.recv(1024).rstrip(b"\n")
        enc = otp(inp)
        req.sendall(str(enc).encode() + b"\n")
        if inp == FLAG:
            exit(0)

class TaskHandler(socketserver.BaseRequestHandler):
    def handle(self):
        challenge(self.request)

if __name__ == '__main__':
    socketserver.ThreadingTCPServer.allow_reuse_address = True
    server = socketserver.ThreadingTCPServer(('0.0.0.0', 8080), TaskHandler)
    server.serve_forever()
