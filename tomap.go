package simple_variable

import (
	"database/sql"
	"errors"
	"reflect"
)

type ToMap struct {

}
func (T *ToMap)Struct2Map(obj interface{}) map[string]string {
	S := new(ToString)
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

func (T *ToMap)Struct2MapInterface(obj interface{}) map[string]interface{} {
	S := new(ToString)
	obj_v := reflect.ValueOf(obj)
	v := obj_v.Elem()
	typeOfType := v.Type()
	var data = make(map[string]interface{})
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
func (T *ToMap) SliceMapColumn(slices []map[string]interface{}, arg string) (results map[string]map[string]interface{} ){
	TS := new(ToString)
	results = map[string]map[string]interface{}{}
	if len(slices) == 0 {
		return
	}
	//var ks string
	for _,v := range slices{
		_,ok := v[arg]
		if !ok{
			continue
		}
		results[TS.Str(v[arg])] = v
	}
	return
}


//切片转 指定key map
func (T *ToMap) SliceMapColumns(slices []map[string]interface{}, args ...string) (results map[string]map[string]interface{} ){
	TS := new(ToString)
	results = map[string]map[string]interface{}{}
	var argKey string
	if len(slices) == 0 {
		return
	}
	//var ks string
	for _,v := range slices{
		argKey = ""
		for _, arg := range args {

			keyS,ok := v[arg]
			if !ok{
				continue
			}
			key := TS.Str(keyS)
			argKey += key +"-"

		}
		results[argKey] = v
	}
	return
}

//切片转 指定key map
func (T *ToMap) StrMapColumn(slices []map[string]string, arg string) (results map[string]map[string]string ){
	//S := new(ToString)
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
	//S := new(ToString)
	results = map[string]string{}
	//var ks string
	for k,v := range arg{
		results[v] = k
	}
	return
}

//flip

//判断是否在map中
func (T *ToMap) InMapInterface(args map[string]interface{} ,l string)(value  bool){
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

//判断是否在map中
func (T *ToMap) InMapKeyInterface(args map[string]interface{} ,l string)(value  bool){
	//TS := new(ToString)
	if len(args) == 0 {
		return
	}
	value = false
	_,ok := args[l]
	if ok {
		value = true
	}
	return
}

//判断是否在map中
func (T *ToMap) InMap(args map[string]string ,l string)(value  bool){
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
func (T *ToMap) ValueToString(args map[string]interface{} )(value  map[string]string){
	ToString := new(ToString)
	value = make(map[string]string)
	for k, val := range args {
		switch val.(type) {
		case string:
			value[k] = val.(string)
		default:
			value[k] = ToString.Str(val)
		}
	}
	return
}





func (T *ToMap) RowsToMap(rows *sql.Rows)(mal []map[string]interface{},errs error)  {
	defer rows.Close()
	errs = nil
	cols, _ := rows.Columns()
	ToString := new(ToString)
	mal =  make([]map[string]interface{},0)
	if rows == nil{
		errs = errors.New("空的不能转换")
	}

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			 errs = err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = ToString.Str(*val)
		}
		mal = append(mal, m)
		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
	}
	return

}



func (T *ToMap)SliceToMaps(mapInstances  []interface{}) (mapInstance map[string]interface{}, err error)  {
	err = nil
	ToString := new(ToString)
	mapInstance = make(map[string]interface{})
	if len(mapInstances) == 0 {
		return
	}
	for k,v := range mapInstances{
		mapInstance[ToString.Str(k)]=v
	}
	return
}



func (T *ToMap)SliceToMapsInterface(mapInstances []map[string]interface{}) (mapInstance map[string]map[string]interface{}, err error)  {
	err = nil
	ToString := new(ToString)
	mapInstance = make(map[string]map[string]interface{})
	if len(mapInstances) == 0 {
		return
	}
	for k,v := range mapInstances{
		mapInstance[ToString.Str(k)]=v
	}
	return
}
func (T *ToMap)SliceToMapsString(mapInstances []map[string]interface{}) (mapInstance  map[string]map[string]string, err error)  {
	err = nil
	ToString := new(ToString)
	mapInstance = make(map[string]map[string]string)
	if len(mapInstances) == 0 {
		return
	}
	for k,v := range mapInstances{
		for key,val := range v{
			mapInstance[ToString.Str(k)][key]= ToString.Str(val)
		}
	}
	return
}




func (T *ToMap) ColumnOne(args map[string]map[string]interface{} ,l string)(value  map[string]interface{}){
	value = make(map[string]interface{})
	if len(args) == 0 {
		return
	}
	for k, val := range args {
		value[k] = val[l]
	}
	return
}


func (T *ToMap) Join(args map[string]interface{} ,l string)(value  string){

	TS := new(ToString)
	if len(args) == 0 {
		return
	}
	i := 0
	for _, val := range args {
		if i >0 {
			value += l
		}
		value += TS.Str(val)
		i ++
	}
	return
}


