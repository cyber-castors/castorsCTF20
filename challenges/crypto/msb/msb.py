from PIL import Image
from numpy import array
from lyrics import lyrics, flag

def pad(m):
    if len(m) % 8 == 0:
        return m
    else:
        return len(m) % 8 * "0" + m


hidden = lyrics * 4 + flag + lyrics * 4
msg = list("0"+format(int.from_bytes(hidden, 'big'), 'b'))

img = array(Image.open('jigglypuff.png'))
h, w, z = img.shape

for x in range(0, h, 5):
    for y in range(w):
        pixel = img[x, y]
        for p in range(3):
            try:
                pixel[p] = (int(msg.pop(0)) * 128) + (int(format(pixel[p] << 1,'08b')[-8:], 2) >> 1)
            except IndexError:
                enc = Image.fromarray(img)
                enc.save('chal.png')
                exit(0)