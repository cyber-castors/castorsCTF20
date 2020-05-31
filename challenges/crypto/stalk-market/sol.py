from pwn import remote

HOST = ('chals20.cybercastors.com', 14423)

sbox = [92, 74, 18, 190, 162, 125, 45, 159, 217, 153, 167, 179, 221, 151, 140, 100, 227, 83, 8, 4, 80, 75, 107, 85, 104, 216, 53, 90, 136, 133, 40, 20, 94, 32, 237, 103, 29, 175, 127, 172, 79, 5, 13, 177, 123, 128, 99, 203, 0, 198, 67, 117, 61, 152, 207, 220, 9, 232, 229, 120, 48, 246, 238, 210, 143, 7, 33, 87, 165, 111, 97, 135, 240, 113, 149, 105, 193, 130, 254, 234, 6, 76, 63, 19, 3, 206, 108, 251, 54, 102, 235, 126, 219, 228, 141, 72, 114, 161, 110, 252, 241, 231, 21, 226, 22, 194, 197, 145, 39, 192, 95, 245, 89, 91, 81, 189, 171, 122, 243, 225, 191, 78, 139, 148, 242, 43, 168, 38, 42, 112, 184, 37, 68, 244, 223, 124, 218, 101, 214, 58, 213, 34, 204, 66, 201, 180, 64, 144, 147, 255, 202, 199, 47, 196, 36, 188, 169, 186, 1, 224, 166, 10, 170, 195, 25, 71, 215, 52, 15, 142, 93, 178, 174, 182, 131, 248, 26, 14, 163, 11, 236, 205, 27, 119, 82, 70, 35, 23, 88, 154, 222, 239, 209, 208, 41, 212, 84, 176, 2, 134, 230, 51, 211, 106, 155, 185, 253, 247, 158, 56, 73, 118, 187, 250, 160, 55, 57, 16, 17, 157, 62, 65, 31, 181, 164, 121, 156, 77, 132, 200, 138, 69, 60, 50, 183, 59, 116, 28, 96, 115, 46, 24, 44, 98, 233, 137, 109, 49, 30, 173, 146, 150, 129, 12, 86, 249]
p = [8, 6, 5, 11, 14, 7, 4, 0, 9, 1, 13, 10, 2, 3, 15, 12]
sbox_inv = [sbox.index(i) for i,j in enumerate(sbox)]
p_inv = [p.index(i) for i,j in enumerate(p)]
round = 8

def start(r):
    r.recvuntil(b'Prophet!\n\n')

def pad(s):
    if len(s) % 16 == 0:
        return s
    else:
        pad_b = 16 - len(s) % 16
        return s + bytes([pad_b]) * pad_b

def repeated_xor(p, k):
    return bytearray([p[i] ^ k[i] for i in range(len(p))])

def group(s):
    return [s[i * 16: (i + 1) * 16] for i in range(len(s) // 16)]

def rev_hash(data, key):
    data = bytes.fromhex(data)
    for _ in range(round):
        temp = bytearray(16)
        for i in range(len(data)):
            temp[p_inv[i]] = data[i]
        for i in range(len(data)):
            temp[i] = sbox_inv[temp[i]]
        temp = repeated_xor(temp, key)
        data = temp
    return data.hex()

def get_target(r, hash, price):
    return rev_hash(hash, pad(f"mon-am-{price}".encode()))

def get_prices(hashes, target, d, highest):
    for hash in hashes:
        for day in d.keys():
            for time in d[day].keys():
                for n in range(20, 601):
                    temp = rev_hash(hash, pad(f"{day}-{time}-{n}".encode()))
                    if temp == target:
                        if n > highest[1]:
                            highest = ("-".join([day, time]), n)
                        d[day][time] = n
                        break
    return d, highest

def solve(r, prices):
    hashes = r.recvuntil(b'\n\n').strip().decode().split()[-12:]
    prices['mon']['am'] = int(r.recvline().strip().decode().split()[3])
    highest = ('mon-am', prices['mon']['am'])
    target = get_target(r, hashes[0], prices['mon']['am'])

    print(f"[+] Matched {highest[0]}-{highest[1]} to {target}")
    prices, highest = get_prices(hashes, target, prices, highest)
    print(f"[+] Highest price is {highest[1]} on {highest[0]}\n")
    
    r.recvuntil(b'week: ')
    r.sendline(highest[0].encode())
    r.recvline()

if __name__ == '__main__':
    r = remote(*HOST)
    start(r)
    for _ in range(20):
        prices = {"mon": {"am": 0, "pm": 0},"tue": {"am": 0, "pm": 0},"wed": {"am": 0, "pm": 0},"thu": {"am": 0, "pm": 0},"fri": {"am": 0, "pm": 0},"sat": {"am": 0, "pm": 0}}
        solve(r, prices)
    else:
        print(r.recv(1024))
