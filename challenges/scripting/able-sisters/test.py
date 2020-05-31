import json
from random import choice
from catalog import items
from villagers import residents

cart = {"accessories": {}, "bottoms": {}, "dresses": {}, "hats": {}, "shoes": {}, "socks": {}, "tops": {}}
total = 0

for _ in range(5):
    category = choice(list(items.keys()))
    item = choice(list(items[category].keys()))
    details = items[category][item]

    cart[category][item] = details
    total += details['priceBuy']
    print(category, item)

print(cart)
print(total)