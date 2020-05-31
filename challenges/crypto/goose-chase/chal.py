from PIL import Image
from numpy import *
from random import randint

img1 = array(Image.open('goose_flag.png'))
img2 = array(Image.open('goose_stole_the_key.png'))
h, w, z = img1.shape

img3 = zeros(shape=(h, w, 3))

for x in range(h):
    for y in range(w):
        pixel1 = img1[x, y]
        pixel2 = img2[x, y]
        pixel3 = img3[x, y]
        for p in range(3):
            pixel3[p] = pixel1[p] ^ pixel2[p]

new_img = Image.fromarray((img3 * 255).astype(uint8))
new_img.save('chal.png')

# img = zeros(shape=img2.shape)
# for x in range(h):
#     for y in range(w):
#         p = img[x, y]
#         for i in range(3):
#             p[i] = randint(0, 255)
# print(img)
# new_img = Image.fromarray((img * 255).astype(uint8))
# new_img.save('noise.png')