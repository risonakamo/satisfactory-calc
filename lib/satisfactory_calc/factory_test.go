package satisfactory_calc

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_factoryCreate(t *testing.T) {
    data:=loadRecipesDict("../../data/factorylab_data.json")

    fac:=createFactory(data["heavy-modular-frame"]["Heavy Modular Frame"])
    pp.Println(fac)

    if fac.TotalOutput!=2 {
        t.Errorf("wrong result")
    }

    fac2:=scaleFactory(fac,2)
    pp.Println(fac2)

    if fac2.TotalOutput!=2 {
        t.Errorf("wrong result")
    }

    fac3:=scaleFactory(fac,4)
    pp.Println(fac3)

    if fac3.TotalOutput!=4 || fac3.BuilderCount!=2 {
        t.Errorf("wrong result")
    }
}

// func Test_factoryConstruct(t *testing.T) {
//     data:=loadRecipesDict("../../data/factorylab_data.json")
//     fac:=createFactory(data["heavy-modular-frame"]["Heavy Modular Frame"])
//     fac=constructFactory(fac,data)

//     pp.Println(fac)
// }

func Test_factoryConstruct2(t *testing.T) {
    data:=loadRecipesDict("../../data/factorylab_data.json")
    fac:=createFactory(data["heavy-modular-frame"]["Heavy Modular Frame"])
    fac,e:=constructFactory2(fac,data,[]string{
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

    if e!=nil {
        t.Error(e)
    }

    pp.SetDefaultMaxDepth(-1)
    pp.BufferFoldThreshold=999999
    // pp.Default.SetColoringEnabled(false)
    // pp.Println(fac)
    longPrintFactory(fac,0)
}