this file transferred into issues

# needs
- when create a cli, would like the following:
    - cli keeps calling func to create a factory, but prompts the user when reach a decision on which recipe to pick
    - user then picks a recipe
    - not sure if possible, but see if can have it print out the current progress of the factory build and what resources are already planned to be used
- for printing, would like the following:
    - [x] better way to distinguish different levels, maybe with number in front of the level line?
    - [x] colour to distinguish clockrate and production amount
    - raw resources needs to be printed and highlighted
        - lets see how it works without explicitly pointing out raw resources. might not be too bad
    - [ ] push all the production numbers to be aligned. will need to pre-generate strings to figure out the longest one
- [ ] calculate total raw resources from factory