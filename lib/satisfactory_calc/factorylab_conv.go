// functions dealing with converting factorylab data to
// our own data structs

package satisfactory_calc

import (
	"satisfactory-calc/lib/factorylab"
	"slices"
)

// declare for certain items a single perferred recipe.
// when outputting recipes for a item that has a preferred recipe, only output
// the recipe that matches the preferred recipe. this causes the item with a preferred
// recipe to only have 1 recipe (the one that was marked preferred)
type PreferredRecpsDict map[string]string

// convert list of fac lab recps to our item recps.
// set excluded producers to exlcude recipes that use certain producers
func convertFacLabRecps(
    recps []factorylab.Recipe,
    excludedProducers []string,
    preferredRecps PreferredRecpsDict,
) []ItemRecipe {
    var result []ItemRecipe

    var recp factorylab.Recipe
    for _,recp = range recps {
        // skip this recp if it is produced by one of the excluded producers
        if slices.ContainsFunc(recp.Producers,func(producer string) bool {
            return slices.Contains(excludedProducers,producer)
        }) {
            continue
        }

        // parse the factory lab recp into item recipes
        var madeRecps []ItemRecipe=facLabRecpToItemRecp(recp)

        // for each newly made item recipe, check if the item it is producing has a preferred
        // recipe. if it does, only push onto the result if it is the preferred recipe.
        var aMadeRecipe ItemRecipe
        for _,aMadeRecipe = range madeRecps {
            var preferredRecipeName string
            var hasAPreferredRecipe bool
            preferredRecipeName,hasAPreferredRecipe=preferredRecps[aMadeRecipe.ItemName]

            // don't push into result recipes if it is not a preferred recipe
            if hasAPreferredRecipe && preferredRecipeName!=aMadeRecipe.RecipeName {
                continue
            }

            result=append(result,aMadeRecipe)
        }
    }

    return result
}

// given a number of items produced per N seconds, compute
// items produced per minute.
// can also be used for inputs per minute:
// output amount -> input amount
// seconds to produce -> seconds per consumption
func itemsPerMinute(
    outputAmount int,
    secondsToProduce int,
) float32 {
    return float32(outputAmount)/(float32(secondsToProduce)/60)
}

// convert fac lab recipe to our item recipe. factorylab recp can produce more than
// 1 recp, as our recipes are always for a single item, while faclab recipes
// match the game which can produce multiple outputs
func facLabRecpToItemRecp(facLabRecp factorylab.Recipe) []ItemRecipe {
    var result []ItemRecipe

    var outputItemName string
    var outputAmount int
    for outputItemName,outputAmount = range facLabRecp.Out {
        result=append(result,ItemRecipe{
            ItemName:outputItemName,
            RecipeName:facLabRecp.Name,
            Producers: facLabRecp.Producers,
            Output:itemsPerMinute(
                outputAmount,
                facLabRecp.Time,
            ),
            Inputs:facLabInputsToInputsDict(facLabRecp.In,facLabRecp.Time),
        })
    }

    return result
}

// convert fac lab inputs dict to our inputs dict
func facLabInputsToInputsDict(
    inputs factorylab.InputsDict,
    secondsPerInput int,
) InputsDict {
    var result InputsDict=InputsDict{}

    var item string
    var neededInput int
    for item,neededInput = range inputs {
        result[item]=itemsPerMinute(
            neededInput,
            secondsPerInput,
        )
    }

    return result
}