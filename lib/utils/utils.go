// random util funcs

package utils

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

	for _ = range amount {
		result+=str
	}

	return result
}