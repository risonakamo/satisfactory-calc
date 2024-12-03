package satisfactory_calc

import (
	"satisfactory-calc/lib/factorylab"
	"testing"

	"github.com/kr/pretty"
)

func Test_recipesDictGen(t *testing.T) {
    facLabData:=factorylab.ReadFactoryLabJson("../../data/factorylab_data.json")
    result:=convertFacLabRecps(facLabData.Recipes,PresetExcludedProducers,PresetPreferredRecipes)
    result2:=groupRecipesIntoDict(result)

    pretty.Println(result2)
}

func Test_loadRecipes(t *testing.T) {
    result:=loadRecipesDict("../../data/factorylab_data.json")

    pretty.Println(result)
}