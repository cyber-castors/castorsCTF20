import json

# json_files = ['accessories.json', 'bottoms.json', 'dresses.json', 'hats.json', 'shoes.json', 'socks.json', 'tops.json']

# parsed_data = {"accessories": {}, "bottoms": {}, "dresses": {}, "hats": {}, "shoes": {}, "socks": {}, "tops": {}}

# for i, item_type in enumerate(parsed_data.keys()):
#     print(item_type)
#     with open(json_files[i], 'r') as f:
#         data = json.loads(f.read())

#     temp = {}
#     for item in data.keys():
#         temp[item] = {k:v for (k,v) in data[item].items() if k in ['name', 'priceBuy']}
#     parsed_data[item_type] = temp

# with open('items.json', 'w') as f:
#     f.write(json.dumps(parsed_data))

# Villagers Parse

with open('villagers.json','r') as f:
    data = json.loads(f.read())

with open('villagers.py','w') as f:
    f.write(json.dumps(list(data.keys())))