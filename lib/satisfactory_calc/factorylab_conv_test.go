package satisfactory_calc

import (
	"satisfactory-calc/lib/factorylab"
	"testing"

	"github.com/kr/pretty"
)

func Test_recpConv(t *testing.T) {
    facLabData:=factorylab.ReadFactoryLabJson("../../data/factorylab_data.json")

    result:=convertFacLabRecps(facLabData.Recipes,PresetExcludedProducers)

    // pp.Print(result)
    pretty.Println(result)
}