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

// flat list of all recipes
var RecipesList []ItemRecipe=[]ItemRecipe{
    {
        ItemName: "heavy modular frame",
        RecipeName: "regular",
        Output: 2,
        Inputs: InputsDict{
            "modular frame":10,
            "steel pipe":40,
            "encased industrial beam":10,
            "screw":240,
        },
    },
    {
        ItemName: "heavy modular frame",
        RecipeName: "heavy encased frame",
        Output: 2.8125,
        Inputs: InputsDict{
            "modular frame":8,
            "steel pipe":33.75,
            "encased industrial beam":9.375,
            "concrete":20.625,
        },
    },
    {
        ItemName: "heavy modular frame",
        RecipeName: "heavy flexible frame",
        Output: 3.75,
        Inputs: InputsDict{
            "modular frame":18.75,
            "rubber":75,
            "encased industrial beam":11.25,
            "screw":390,
        },
    },
    {
        ItemName: "modular frame",
        RecipeName: "regular",
        Output: 2,
        Inputs: InputsDict{
            "reinforced iron plate":3,
            "iron rod":12,
        },
    },
    {
        ItemName: "modular frame",
        RecipeName: "bolted frame",
        Output: 5,
        Inputs: InputsDict{
            "reinforced iron plate":7.5,
            "screw":140,
        },
    },
    {
        ItemName: "modular frame",
        RecipeName: "steeled frame",
        Output: 3,
        Inputs: InputsDict{
            "reinforced iron plate":2,
            "steel pipe":10,
        },
    },
}

// recipes collection as dict
// todo: generate this
var Recipes RecipesDict=RecipesDict{}