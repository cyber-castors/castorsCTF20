import random

articles = ("a", "the")
adj = ("adorable", "clueless", "dirty", "odd", "fast")
nouns = ("puppy", "car", "rabbit", "girl", "monkey", "boy")
verbs = ("runs", "hits", "jumps", "drives", "sings to") 
adv = ("crazily.", "dutifully.", "foolishly.", "merrily.", "occasionally.")
sentence = [articles, adj, nouns, verbs, articles, nouns, adv]

def gen_rand_sentence():
    return " ".join(random.choice(word) for word in sentence).capitalize().encode()

def add_noise(m):
    ret = []
    for b in format(int.from_bytes(m,'big'),'b'):
        r = random.randint(1, 100)
        if r == 1:
            ret.append(str(int(b) ^ r))
        else:
            ret.append(b)
    return int("".join(ret),2).to_bytes(len(m),'big')

def gen_chal():
    target = gen_rand_sentence()
    corrupt = add_noise(target)
    return target, corrupt, format(int.from_bytes(target,'big') ^ int.from_bytes(corrupt,'big'),'b').count('1')

t, c, d = gen_chal()

print(len(t), len(c), d)
print(t,c)
