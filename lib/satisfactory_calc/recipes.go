// components related to recipe data

package satisfactory_calc

import (
	"fmt"
	"satisfactory-calc/lib/factorylab"
)

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

    Producers []string

    // items per minute
    Output float32
    Inputs InputsDict
}

// preset producers to exclude to prevent infinite loops. might figure out way to
// handle these latera
var PresetExcludedProducers []string=[]string{
    "converter",
    "packager",
    "quantum-encoder",
    "blender",
}

var PresetPreferredRecipes PreferredRecpsDict=PreferredRecpsDict{

}

var PresetExcludedRecipes []string=[]string{

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

// get recipes dict from factory lab data json with some preset filtering options
func LoadRecipesDict(path string) RecipesDict {
    var facLabData factorylab.FactorylabJson=factorylab.ReadFactoryLabJson(path)
    return groupRecipesIntoDict(
        convertFacLabRecps(
            facLabData.Recipes,
            PresetExcludedProducers,
            PresetPreferredRecipes,
            PresetExcludedRecipes,
        ),
    )
}

// merge inputs dicts by adding all inputs together
func mergeInputDict(inputs1 InputsDict,inputs2 InputsDict) InputsDict {
    var result InputsDict=InputsDict{}

    // looping over input1, grab and add from input2. due to defaulting to 0
    // if it doesnt exist, this updates all using input1
    var item string
    var amount float32
    for item,amount = range inputs1 {
        result[item]=amount+inputs2[item]
    }

    // now need to catch items that was in input2 but not in 1
    var in bool
    for item,amount = range inputs2 {
        _,in=inputs1[item]

        // add all items that was not in inputs1
        if !in {
            result[item]=amount
        }
    }

    return result
}

// format print inputs dict
func printInputsDict(inputs InputsDict) {
    var item string
    var amount float32
    for item,amount = range inputs {
        fmt.Printf("%s: %.2f\n",
            item,
            amount,
        )
    }
}

// pretty print list of recipes
func PrintRecipesList(recipes []string) {
    var recipe string
    for _,recipe = range recipes {
        fmt.Printf("- %s\n",recipe)
    }
}