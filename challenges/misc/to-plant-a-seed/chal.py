import random

def xor(a, b):
    try:
        assert len(a) == len(b)
        return [a[i] ^ b[i] for i in range(len(a))]
    except:
        print("Unequal lengths.")

day = random.randint(1, 365)
random.seed(day)
seed = [random.randint(0, 255) for i in range(42)]
flag = b'castors{d0n7_f0rg37_t0_73nd_t0_y0ur_s33ds}'
print(len(flag))


a,b = [1,2], [1,2]
water = xor(seed,flag)
print(bytes(water))