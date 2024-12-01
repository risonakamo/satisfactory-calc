// functions dealing with converting factorylab data to
// our own data structs

package satisfactory_calc

import "satisfactory-calc/lib/factorylab"

// given a number of items produced per N seconds, compute
// items produced per minute.
// can also be used for inputs per minute:
// output amount -> input amount
// seconds to produce -> seconds per consumption
func itemsPerMinute(
    outputAmount int,
    secondsToProduce int,
) float32 {
    return float32(outputAmount)/float32(secondsToProduce)/60
}

// convert fac lab recipe to our item recipe
func facLabRecpToItemRecp(facLabRecp factorylab.Recipe) ItemRecipe {
    return ItemRecipe{
        ItemName:facLabRecp.Id,
        RecipeName:facLabRecp.Name,
        Output:itemsPerMinute(
            facLabRecp.Out[facLabRecp.Id],
            facLabRecp.Time,
        ),
        Inputs:facLabInputsToInputsDict(facLabRecp.In,facLabRecp.Time),
    }
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