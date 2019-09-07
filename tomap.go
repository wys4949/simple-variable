package simple_variable

import "reflect"

type ToMap struct {

}
func (T *ToMap)Struct2Map(obj interface{}) map[string]string {
	S := new(ToSring)
	obj_v := reflect.ValueOf(obj)
	v := obj_v.Elem()
	typeOfType := v.Type()
	var data = make(map[string]string)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		data[typeOfType.Field(i).Name] = S.Str(field.Interface())
	}
	return data
}





func (T *ToMap)Struct2Maps(objs []map[string]string) (newobjs []map[string]string) {
	newobjs = make([]map[string]string,0)
	ToSlice :=new(ToSlice)
	val, ok := ToSlice.CreateAnyTypeSlice(objs)
	if !ok {
		return
	}

	if len(val) == 0 {
		return
	}
	for _,obj := range val{
		newobjs = append(newobjs, T.Struct2Map(obj))
	}
	return
}
//切片转 指定key map
func (T *ToMap) MapColumn(slices []map[interface{}]interface{}, arg string) (results map[interface{}]map[interface{}]interface{} ){
	//S := new(ToSring)
	results = map[interface{}]map[interface{}]interface{}{}
	if len(slices) == 0 {
		return
	}
	//var ks string
	for _,v := range slices{
		_,ok := v[arg]
		if !ok{
			continue
		}
		results[v[arg]] = v
	}
	return
}
//切片转 指定key map
func (T *ToMap) StrMapColumn(slices []map[string]string, arg string) (results map[string]map[string]string ){
	//S := new(ToSring)
	results = map[string]map[string]string{}
	if len(slices) == 0 {
		return
	}
	//var ks string
	for _,v := range slices{
		_,ok := v[arg]
		if !ok{
			continue
		}
		results[v[arg]] = v
	}
	return
}

//reverse

//切片转 指定key map
func (T *ToMap) Reverse(arg map[string]string) (results map[string]string ){
	//S := new(ToSring)
	results = map[string]string{}
	//var ks string
	for k,v := range arg{
		results[v] = k
	}
	return
}