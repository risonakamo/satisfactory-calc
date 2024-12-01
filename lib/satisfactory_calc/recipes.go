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

// a single item's recipe specifications
type ItemRecipe struct {
    // name of item
    ItemName string
    // name of recipe
    RecipeName string

    Output float32
    Inputs InputsDict
}