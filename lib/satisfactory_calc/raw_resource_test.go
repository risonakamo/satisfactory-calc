package satisfactory_calc

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/sets"
)

func Test_calcResource(t *testing.T) {
	data := loadRecipesDict("../../data/factorylab_data.json")
	fac := createFactory(data["heavy-modular-frame"]["Heavy Modular Frame"])
	fac, e := constructFactory2(fac, data, []string{
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


    result:=calculateResourceUse(fac,sets.New([]string{
        "iron-ore",
        "coal",
        "limestone",
    }...))

    printInputsDict(result)
}