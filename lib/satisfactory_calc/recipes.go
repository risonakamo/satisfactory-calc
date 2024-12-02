// components related to recipe data

package satisfactory_calc

// collection of recipes, sorted by item/recipe
// key: item name
// val: the recipe as dict of all its alternates
type RecipesDict map[string]AlternatesDict

// collection of alternate recipes for 1 item
// key: alternate recipe name
// val: the recipe
type AlternatesDict map[string]ItemRecipe

// dict that defines inputs requirements.
// key: name of item. should exist in recipes dict
// val: amount required
type InputsDict map[string]float32

// a single item's recipe specifications.
// there can be multiple recipes for the same item - all recipes
// will share the ItemName, but RecipeName will be different.
// recipe name might also be shared as some recipes produce multiple items.
// thus, to uniquely identify a recipe, must use both itemname and recipe.
type ItemRecipe struct {
    // name of item
    ItemName string
    // name of recipe
    RecipeName string

    Output float32
    Inputs InputsDict
}