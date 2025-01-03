package satisfactory_calc

import (
	"maps"
	"satisfactory-calc/lib/utils"
	"slices"

	"k8s.io/apimachinery/pkg/util/sets"
)

// default raw resources to use for resource calculation
var DefaultRawResources sets.Set[string]=sets.New([]string{
    "iron-ore",
    "coal",
    "limestone",
    "water",
    "copper-ore",
    "crude-oil",
    "raw-quartz",
    "caterium-ore",
}...)

// calculate total resource use of a factory given target resources
// output is inputs dict which doubles as outputs dict.
func CalculateResourceUse(fac Factory,targetResources sets.Set[string]) InputsDict {
    var result InputsDict=InputsDict{}

    if targetResources.Has(fac.ItemName) {
        result[fac.ItemName]=fac.TotalOutput
    }

    var subFactDict FactorybyRecipe
    for _,subFactDict = range fac.SubFactories {
        var subFactory Factory=utils.GetDictFirstItem(subFactDict)
        result=mergeInputDict(
            result,
            CalculateResourceUse(subFactory,targetResources),
        )
    }

    return result
}

// convert a recipes dict to raw resource dict, which would count every resource
// as a raw resource
func RecipesDictToRawResourceSet(recipes RecipesDict) sets.Set[string] {
    return sets.New(
        slices.Collect(
            maps.Keys(recipes),
        )...,
    )
}