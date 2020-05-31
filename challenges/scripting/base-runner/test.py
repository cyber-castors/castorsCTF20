import random, string
from base64 import b64encode

def gen_random_word(length):
   letters = string.ascii_lowercase
   return ''.join(random.choice(letters) for i in range(length))

def gen_chal():
    target = b'castorsCTF{' + gen_random_word(random.randint(10,20)).encode() + b'}'
    enc = b64encode(target)
    for base in ['02x', 'o', '08b']:
        enc = " ".join([format(i, base) for i in enc]).encode()
    
    return target, enc

if __name__ == '__main__':
    target, enc = gen_chal()
    print(target)
    print(enc)
    # for i in range(10):
    #     print(gen_random_word(random.randint(10, 20)))