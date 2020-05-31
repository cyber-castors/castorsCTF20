from PIL import Image
from numpy import array

img = array(Image.open('chal.png'))
h, w, z = img.shape

dec = []

for x in range(h):
    for y in range(w):
        pixel = img[x, y]
        for p in range(3):
            dec.append(format(pixel[p],'08b')[0])
        
msg = bytes.fromhex(format(int("00"+"".join(dec), 2), 'x'))

with open('flag.txt', 'wb') as f:
    f.write(msg)
