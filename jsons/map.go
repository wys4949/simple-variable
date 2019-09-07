package simple_variable

import (
	"encoding/json"
	simple_variable "github.com/wys4949/simple-variable"
)

type Map struct {
}
func (m *Map)Encode(mapInstance map[string]string) (jsonStr string, err error)  {
	T := new(simple_variable.ToSring)
	jsonStrs, err := json.Marshal(mapInstance)
	jsonStr = ""
	if err != nil {
		return
	}
	jsonStr = T.Str(jsonStrs)
	return
}

func (m *Map)Decode(jsonInstance string) (mapInstance map[string]string, err error)  {
	T := new(simple_variable.ToSring)
	jsonStr := T.Byte(jsonInstance)
	err = json.Unmarshal([]byte(jsonStr), &mapInstance)
	return
}