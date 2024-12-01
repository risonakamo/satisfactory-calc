package factorylab

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_readJson(t *testing.T) {
    result:=readFactoryLabJson("../../data/factorylab_data.json")

    pp.Print(result)
}