package satisfactory_calc

import (
	"satisfactory-calc/lib/factorylab"
	"testing"

	"github.com/k0kubun/pp/v3"
	"github.com/kr/pretty"
)

func Test_recipesDictGen(t *testing.T) {
    facLabData:=factorylab.ReadFactoryLabJson("../../data/factorylab_data.json")
    result:=convertFacLabRecps(facLabData.Recipes,PresetExcludedProducers,
        PresetPreferredRecipes,PresetExcludedRecipes)
    result2:=groupRecipesIntoDict(result)

    pretty.Println(result2)
}

func Test_loadRecipes(t *testing.T) {
    result:=LoadRecipesDict("../../data/factorylab_data.json")

    pretty.Println(result)
}

func Test_mergeInputs(t *testing.T) {
    var inputs1 InputsDict=InputsDict{
        "asda":3,
        "b":4,
        "d":90,
    }

    var inputs2 InputsDict=InputsDict{
        "asda":10,
        "b":1,
        "c":111,
    }

    pp.Println(mergeInputDict(inputs1,inputs2))
}

func Test_printAlternates(t *testing.T) {
    data:=LoadRecipesDict("../../data/factorylab_data.json")

    recipes:=data["heavy-modular-frame"]

    PrintAlternatesDict(recipes,20)
}

func Test_scaleAlternates(t *testing.T) {
    data:=LoadRecipesDict("../../data/factorylab_data.json")

    recipes:=data["heavy-modular-frame"]

    result:=ScaleAlternatesDict(recipes,20)

    PrintAlternatesDict(result,20)
}