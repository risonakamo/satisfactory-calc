// random util funcs

package utils

import (
	"os"
	"path/filepath"
)

// get 1st item in a dict
func GetDictFirstItem[KeyT comparable,ValT any](dict map[KeyT]ValT) ValT {
	var key KeyT
	for key = range dict {
		return dict[key]
	}

	panic("empty")
}

// generate a string by duplicating the selected string a number of times
func DuplicateString(str string,amount int) string {
	var result string=""

	for range amount {
		result+=str
	}

	return result
}

// give folder location of the exe that calls this func
func GetHereDirExe() string {
    var exePath string
    var e error
    exePath,e=os.Executable()

    if e!=nil {
        panic(e)
    }

    return filepath.Dir(exePath)
}