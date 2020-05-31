import base64
import string

flag = "eHpzdG9yc1hXQXtpYl80cjFuMmgxNDY1bl80MXloMF82Ml95MDQ0MHJfNGQxbl9iNXVyMn0="
customMapping = b"ZYXFGABHOPCDEQRSTUVWIJKLMNabczyxjklmdefghinopqrstuvw5670123489+/_{}"
normalMapping = b"ABCDEFGHIJKLMNOPQRSTUVWZYXzyxabcdefghijklmnopqrstuvw0123456789+/_{}"

decoded_flag = base64.b64decode(flag) 

translation_table = decoded_flag.maketrans(customMapping, normalMapping)

print(decoded_flag.translate(translation_table))

