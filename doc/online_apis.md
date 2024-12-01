online apis to get json data

# links
- https://satisfactory.wiki.gg/wiki/Online_tools
- https://api.calculatory.ovh/data/items
- https://api.calculatory.ovh/data/recipes
- https://factoriolab.github.io/data/sfy/data.json

# requirements
To fill out our items list, we need, for each item+recipe combo:

- output amount
- item inputs and amount per input
- name of item
- name of recipe

Each recipe should be able to link with it's "original" item.

# ovh calc data analysis
- does not provide name of recp
- need to calculate inputs/min

# factorylab
- can identify recipe's link with original item with the output item
- need to calculate inputs/min
- has name of alternate recps

# conclusion
factorylab seems to have what we would need. should be able to create converter func that converts the json into our own json, then don't need to continuously run this func.