// components related to factory struct

package satisfactory_calc

import (
	"errors"
	"fmt"
	"math"
	"satisfactory-calc/lib/utils"
	"slices"

	"github.com/fatih/color"
	"github.com/k0kubun/pp/v3"
)

// factories all for a single item grouped by recipe name.
// all factories in this dict should produce the same itema
// key: recipe name
// val: factory producing an item following the keyed recipe
type FactorybyRecipe map[string]Factory

// collection of factories to match an inputs dict.
// key: name of item that should also exist in inputs dict
// val: dict of factories (grouped by recipe) which produce for the corresponding
// inputs dict
type SubFactoriesDict map[string]FactorybyRecipe

// defines a single factory which builds 1 thing. factory consists
// of multiple builders
type Factory struct {
    // name of item. should correspond with an item name in
    // the recipes dict
    ItemName string
    RecipeName string

    // number of builders that make up this factory
    BuilderCount int
    // clock speed to set each builder in this factory to. float percent
    ClockRate float32

    // amount each builder unit in this factory should be making
    // items per minute
    OutputPerBuilder float32
    // total output produced by this entire factory using all of its
    // builder units.
    // items per minute
    TotalOutput float32

    // requirements for each single builder in this factory.
    InputsPerBuilder InputsDict

    // total requirements of this factory. calculated by multiplying
    // inputs per builder by builder count.
    TotalInputs InputsDict

    // sub factories that feed into this factory.
    // these sub factories should be properly scaled to meet this factory's specified
    // total inputs
    SubFactories SubFactoriesDict

    Recipe ItemRecipe
}

// initialise a factory for the target recipe. factory will be 1 builder.
// sub factories are not yet calculated
func createFactory(recp ItemRecipe) Factory {
    return Factory{
        ItemName: recp.ItemName,
        RecipeName: recp.RecipeName,
        BuilderCount: 1,
        ClockRate: 1,
        OutputPerBuilder: recp.Output,
        TotalOutput: recp.Output,
        InputsPerBuilder: recp.Inputs,
        TotalInputs: recp.Inputs,
        SubFactories: SubFactoriesDict{},
        Recipe: recp,
    }
}

// modify a factory so it produces the target output amount.
// modifies several fields, but does NOT affect subfactories.
// returns the new modified factory
func scaleFactory(factory Factory,targetOutputAmount float32) Factory {
    // need this many factories at 100% speed
    var floatFactoriesNeeded float32=targetOutputAmount/factory.Recipe.Output

    // round up to get factories needed if we're going to underclock the factories (also can't
    // have partial factories)
    var factoriesNeeded int=int(math.Ceil(float64(floatFactoriesNeeded)))

    // this is the output we need out of each factory if we use `factoriesNeeded` factories
    var buildOutputNeededPerBuilder float32=targetOutputAmount/float32(factoriesNeeded)

    // this is the clock rate to set each of the sub-factories to. calculated as the needed output
    // just calculated is always a partial of the original 100% output
    var builderClockRate float32=buildOutputNeededPerBuilder/factory.Recipe.Output

    return Factory{
        ItemName: factory.ItemName,
        RecipeName: factory.RecipeName,
        BuilderCount: factoriesNeeded,
        ClockRate: builderClockRate,
        OutputPerBuilder: buildOutputNeededPerBuilder,
        TotalOutput: targetOutputAmount,
        InputsPerBuilder: scaleInputsToClockrate(
            factory.Recipe.Inputs,
            builderClockRate,
            1,
        ),
        TotalInputs: scaleInputsToClockrate(
            factory.Recipe.Inputs,
            builderClockRate,
            factoriesNeeded,
        ),
        SubFactories: factory.SubFactories,
        Recipe: factory.Recipe,
    }
}

// scale inputs dict to a certain clock rate
func scaleInputsToClockrate(inputs InputsDict,clockRate float32,builders int) InputsDict {
    var result InputsDict=InputsDict{}

    var item string
    var amount float32
    for item,amount = range inputs {
        result[item]=amount*clockRate*float32(builders)
    }

    return result
}

// calculate all sub factories for a given factory.
// the factory should be scaled to the desired amount before
// calling this function
func dep_constructFactory(fact Factory,recps RecipesDict) Factory {
    var subFactories SubFactoriesDict=SubFactoriesDict{}

    var item string
    var neededAmount float32
    for item,neededAmount = range fact.TotalInputs {
        var alternateRecps AlternatesDict=recps[item]

        subFactories[item]=FactorybyRecipe{}

        var subRecipeName string
        var subRecipe ItemRecipe
        for subRecipeName,subRecipe = range alternateRecps {
            fmt.Println("calculating",subRecipeName)
            subFactories[item][subRecipeName]=dep_constructFactory(
                scaleFactory(
                    createFactory(subRecipe),
                    neededAmount,
                ),
                recps,
            )
        }
    }

    fact.SubFactories=subFactories

    return fact
}

// construct a factory, filling out the subfactories.
// give list of recipes to use. if reach an item where a recp isn't selected, will fail
// and report the issue
func constructFactory2(
    fact Factory,
    recps RecipesDict,
    recpsSelection []string,
) (Factory,error) {
    var subFactories SubFactoriesDict=SubFactoriesDict{}

    var item string
    var neededAmount float32
    for item,neededAmount = range fact.TotalInputs {
        // recps for the current item
        var alternateRecps AlternatesDict=recps[item]

        var foundRecp ItemRecipe=getRecpFromSelections(
            alternateRecps,
            recpsSelection,
        )

        if len(foundRecp.ItemName)==0 {
            fmt.Println("failed to find recipe for:",item)
            fmt.Println("need to create:",neededAmount)
            fmt.Println("available recipes:")
            pp.Println(alternateRecps)
            return Factory{},errors.New("missing recipe")
        }

        var madeFactory Factory
        var e error
        madeFactory,e=constructFactory2(
            scaleFactory(
                createFactory(foundRecp),
                neededAmount,
            ),
            recps,
            recpsSelection,
        )

        if e!=nil {
            return Factory{},e
        }

        subFactories[item]=FactorybyRecipe{
            foundRecp.RecipeName:madeFactory,
        }
    }

    fact.SubFactories=subFactories

    return fact,nil
}

// given alternates dict and selection of recps, get the 1st recipe
// that can be found in the recps selection.
// if can't find, returned recipe will have empty name
func getRecpFromSelections(
    recipesDict AlternatesDict,
    recpsSelection []string,
) ItemRecipe {
    var recpName string
    for recpName = range recipesDict {
        if slices.Contains(recpsSelection,recpName) {
            return recipesDict[recpName]
        }
    }

    return ItemRecipe{}
}

// format print single factory
func printFactory(fac Factory) {
    fmt.Printf("%dx %s (%s) @ %s -> %dx%s = %s\n",
        fac.BuilderCount,
        fac.ItemName,
        fac.RecipeName,
        color.YellowString("%.2f",fac.ClockRate),
        fac.BuilderCount,
        color.CyanString("%.2f",fac.OutputPerBuilder),
        color.HiGreenString("%.2f",fac.TotalOutput),
    )
}

// print factory and all subfactories
func longPrintFactory(
    fac Factory,
    indentLevel int,
    indentSize int,
) {
    var indentStr string=utils.DuplicateString(" ",indentLevel*indentSize)
    fmt.Printf("[ ] %s%s: ",
        indentStr,
        color.RedString("%d",indentLevel+1),
    )
    printFactory(fac)

    var subFactoryDict FactorybyRecipe
    for _,subFactoryDict = range fac.SubFactories {
        var subFactory Factory=utils.GetDictFirstItem(subFactoryDict)
        longPrintFactory(subFactory,indentLevel+1,indentSize)
    }
}