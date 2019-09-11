package simple_variable

import (
	"encoding/json"
	simple_variable "github.com/wys4949/simple-variable"
)

type Maps struct {
}

func (m *Maps)Encode(mapInstances []map[string]string) (jsonStr string, err error)  {
	T := new(simple_variable.ToString)
	jsonStrs, err := json.Marshal(mapInstances)
	jsonStr = ""
	if err != nil {
		return
	}
	jsonStr = T.Str(jsonStrs)
	return
}
func (m *Maps)Decode(jsonInstance string) (mapInstance []map[string]string, err error)  {
	T := new(simple_variable.ToString)
	jsonStr := T.Byte(jsonInstance)
	err = json.Unmarshal(jsonStr, &mapInstance)
	return
}