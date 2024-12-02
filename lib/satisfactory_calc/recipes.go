// components related to recipe data

package satisfactory_calc

import "satisfactory-calc/lib/factorylab"

// collection of recipes, sorted by item/recipe
// key: item name
// val: the recipe as dict of all its alternates
// essentially, a 2-level grouped dict
// {
//   [itemname]:{
//     [recipename]:<the recipe>
//   }
// }
type RecipesDict map[string]AlternatesDict

// collection of alternate recipes for 1 item
// key: alternate recipe name
// val: the recipe
type AlternatesDict map[string]ItemRecipe

// dict that defines inputs requirements.
// key: name of item. should exist in recipes dict
// val: amount required (items per minute)
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

    // items per minute
    Output float32
    Inputs InputsDict
}

// preset producers to exclude to prevent infinite loops. might figure out way to
// handle these latera
var PresetExcludedProducers []string=[]string{
    "converter",
    "packager",
}

// produce organised recipes dict from recipes list
func groupRecipesIntoDict(recps []ItemRecipe) RecipesDict {
    var result RecipesDict=RecipesDict{}

    var recp ItemRecipe
    for _,recp = range recps {
        var in bool
        _,in=result[recp.ItemName]

        if !in {
            result[recp.ItemName]=AlternatesDict{}
        }

        result[recp.ItemName][recp.RecipeName]=recp
    }

    return result
}

// get recipes dict from factory lab data json. uses preset producers for now.
func loadRecipesDict(path string) RecipesDict {
    var facLabData factorylab.FactorylabJson=factorylab.ReadFactoryLabJson(path)
    return groupRecipesIntoDict(
        convertFacLabRecps(facLabData.Recipes,PresetExcludedProducers),
    )
}