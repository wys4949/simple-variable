package simple_variable

import (
	"encoding/json"
	simple_variable "github.com/wys4949/simple-variable"
)

type Slices struct {
}

func (m *Slices)Encode(mapInstances []interface{}) (jsonStr string, err error)  {
	T := new(simple_variable.ToSring)
	jsonStrs, err := json.Marshal(mapInstances)
	jsonStr = ""
	if err != nil {
		return
	}
	jsonStr = T.Str(jsonStrs)
	return
}
func (m *Slices)Decode(jsonInstance string) (mapInstance []string, err error)  {
	T := new(simple_variable.ToSring)
	jsonStr := T.Byte(jsonInstance)
	err = json.Unmarshal(jsonStr, &mapInstance)
	return
}