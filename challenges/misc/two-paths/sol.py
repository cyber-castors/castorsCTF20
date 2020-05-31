import string
from collections import Counter

etaoin = " ontehfxjrpzsua={ckgby}dwqlvim.!,'\n"
enc = "_♓♒❌✖⏫💯🔴➿🔁♑📶♊🈲♉={♈🆔🌀♌🔟}⛎🚺♐♏➗⏺Ⓜ.!,'\n"
print(enc.encode())


with open('emoji.txt', encoding='utf-8') as f:
    data = f.read()

ucode = Counter(data)
ucode = {k: v for k, v in sorted(ucode.items(), reverse=True, key=lambda item: item[1])}
# print("".join(ucode))

sub = dict(zip(enc, etaoin))

for ch in data:
    print(sub[ch],end='')