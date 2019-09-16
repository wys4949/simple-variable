package simple_variable

import (
	"reflect"
)

type ToSlice struct {

}


//切片map截取
func (T *ToSlice) Test( args ...string) (results bool ){
	results = true
	return
}

//切片map截取
func (T *ToSlice) Column(slices []map[string]string, args ...string) (results []map[string]string ){
	//S := new(ToString)
	results = make([]map[string]string,0)
	if len(slices) == 0 {
		return
	}
	var result  map[string]string
	//var ks string
	for _,v := range slices{
		result = map[string]string{}
		for _,arg := range args{
			//ks = S.Str(k)
			_, ok:= v[arg]
			if !ok{
				continue
			}
			result[arg] = v[arg]
		}
		results = append(results, result)
	}
	return
}


// 判断是否为slcie数据
func (T *ToSlice) IsSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}




// interface{}转为 []interface{}
func  (T *ToSlice) CreateAnyTypeSlice(slice interface{}) ([]interface{}, bool) {
	val, ok := T.IsSlice(slice)

	if !ok {
		return nil, false
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out, true
}

func (T *ToSlice) ColumnOne(args []map[string]string ,l string)(value  []string){
	value = make([]string,0)
	if len(args) == 0 {
		return
	}
	for _, val := range args {
		value = append(value, val[l])
	}
	return
}


func (T *ToSlice) Join(args []string ,l string)(value  string){

	//TS := new(ToString)
	if len(args) == 0 {
		return
	}
	for k, val := range args {
		if k >0 {
			value += l
		}
		value += val
	}
	return
}


//判断是否在切片中
func (T *ToSlice) InSlice(args []string ,l string)(value  bool){
	//TS := new(ToString)
	if len(args) == 0 {
		return
	}
	value = false
	for _, val := range args {
		if val == l{
			value = true
			break
		}
	}
	return
}



func (T *ToSlice) CamelString(args []string )(value  []string){

	TS := new(ToString)
	value = make([]string, 0)
	if len(args) == 0 {
		return
	}
	for _, val := range args {
		value = append(value, TS.CamelString(val))
	}
	return
}



func (T *ToSlice) SnakeString(args []string )(value  []string){

	TS := new(ToString)
	value = make([]string, 0)
	if len(args) == 0 {
		return
	}
	for _, val := range args {
		value = append(value, TS.SnakeString(val))
	}
	return
}




//


