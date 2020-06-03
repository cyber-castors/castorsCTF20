import random

def xor(a, b):
    try:
        assert len(a) == len(b)
        return [str(a[i] ^ b[i]) for i in range(len(a))]
    except:
        print("Unequal lengths.")

timestamp = 1590782400
random.seed(timestamp)
seed = [random.randint(0, 255) for i in range(42)]
flag = b'castors{d0n7_f0rg37_t0_73nd_t0_y0ur_s33ds}'

water = xor(seed,flag)

with open("water.txt", "w") as f:
    f.write(" ".join(water))