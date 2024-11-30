online apis to get json data

# links
- https://satisfactory.wiki.gg/wiki/Online_tools
- https://api.calculatory.ovh/data/items
- https://api.calculatory.ovh/data/recipes

# ovh calc data analysis
These jsons should provide recipe information for any given item, but doesn't include the name of the recipe.

Also, it includes the items made and time in seconds to make, so it doesn't have items per minute, but that can be calculated.

To fill out our items list, we need, for each item+recipe combo:

- output amount
- item inputs and amount per input
- name of item
- name of recipe

The data has all of this except for recipe name