package satisfactory_calc

import (
	"satisfactory-calc/lib/utils"

	"k8s.io/apimachinery/pkg/util/sets"
)

// calculate total resource use of a factory given target resources
// output is inputs dict which doubles as outputs dict.
func calculateResourceUse(fac Factory,targetResources sets.Set[string]) InputsDict {
    var result InputsDict=InputsDict{}

    if targetResources.Has(fac.ItemName) {
        result[fac.ItemName]=fac.TotalOutput
    }

    var subFactDict FactorybyRecipe
    for _,subFactDict = range fac.SubFactories {
        var subFactory Factory=utils.GetDictFirstItem(subFactDict)
        result=mergeInputDict(
            result,
            calculateResourceUse(subFactory,targetResources),
        )
    }

    return result
}