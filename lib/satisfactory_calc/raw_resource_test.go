package satisfactory_calc

import (
	"testing"

	"github.com/k0kubun/pp/v3"
	"k8s.io/apimachinery/pkg/util/sets"
)

func Test_calcResource(t *testing.T) {
	data := LoadRecipesDict("../../data/factorylab_data.json")
	fac := CreateFactory(data["heavy-modular-frame"]["Heavy Modular Frame"])
	fac, e := ConstructFactory2(fac, data, []string{
		"Steel Pipe",
		"Steel Ingot",
		"Modular Frame",
		"Reinforced Iron Plate",
		"Iron Plate",
		"Iron Rod",
		"Iron Ingot",
		"Iron Ore",
		"Screw",
		"Coal",
		"Encased Industrial Beam",
		"Concrete",
		"Steel Beam",
		"Limestone",
	})

	if e != nil {
		t.Error(e)
	}


    result:=CalculateResourceUse(fac,sets.New([]string{
        "iron-ore",
        "coal",
        "limestone",
    }...))

    PrintInputsDict(result)
}

func Test_allRawResources(t *testing.T) {
	data := LoadRecipesDict("../../data/factorylab_data.json")

	result:=RecipesDictToRawResourceSet(data)

	pp.Println(result)
}