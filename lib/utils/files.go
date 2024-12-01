// general files read/write utils

package utils

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/rs/zerolog/log"
)

// read json file and return result
func ReadJson[DataT any](filename string) (DataT,error) {
	var data []byte
	var e error
	data,e=os.ReadFile(filename)

	if errors.Is(e,fs.ErrNotExist) {
		log.Info().Msgf("file not found: %s",filename)
		var def DataT
		return def,e
	}

	if e!=nil {
		var def DataT
		return def,e
	}

	var parsedData DataT
	json.Unmarshal(data,&parsedData)

	return parsedData,nil
}