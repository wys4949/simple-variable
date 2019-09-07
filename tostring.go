package simple_variable

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ToSring struct {

}

//获取类型
func (T *ToSring) TypeOf(v interface{}) (t string ){
	t = "nil"
	if v != nil{
		t = reflect.TypeOf(v).String()
	}
	return  t
}

//转str
func (T *ToSring) Str(v interface{}) (returnVar  string) {
	// inter.(type)
	varType := T.TypeOf(v)
	switch varType {
	case "nil":
		returnVar = ""
	case "uint":
		returnVar =strconv.Itoa( T.Int( v.(uint)))
	case "uint8":
		returnVar =strconv.Itoa( T.Int( v.(uint8)))
	case "uint32":
		returnVar =strconv.Itoa( T.Int( v.(uint32)))
	case "uint64":
		returnVar =strconv.Itoa( T.Int( v.(uint64)))
	case "int":
		returnVar = strconv.Itoa(v.(int))
	case "int64":
		returnVar = strconv.FormatInt(v.(int64), 10)
	case "int32":
		returnVar = strconv.FormatInt(int64(v.(int32)), 10)
	case "int16":
		returnVar = strconv.FormatInt(int64(v.(int16)), 10)
	case "string":
		returnVar = v.(string)
	case "float64":
		returnVar = strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case "[]uint8":
		returnVar = string(v.([]byte)[:])
	case "time.Time":
		returnVar =  v.(time.Time).Format("2006-01-02 15:04:05")
	case  "bool":
		if v.(bool) {
			returnVar = ""
		}else{
			returnVar = "1"

		}
	}
	return
}

func (T *ToSring) StrAdd(args ...interface{}) (str string) {
	str = ""
	for _, arg := range args {

		switch T.TypeOf(arg) {
		case "nil":
			str = str + ""
		case "uint":
			str = str + strconv.Itoa(T.Int(arg.(uint)))
		case "uint8":
			str = str + strconv.Itoa(T.Int(arg.(uint8)))
		case "uint32":
			str = str + strconv.Itoa(T.Int(arg.(uint32)))
		case "uint64":
			str = str + strconv.Itoa(T.Int(arg.(uint64)))
		case "int":
			str = str + strconv.Itoa(arg.(int))
		case "string":
			str = str + arg.(string)
		case "int64":
			str = str + strconv.FormatInt(arg.(int64), 10)
		case "int32":
			str = str + strconv.FormatInt(int64(arg.(int32)), 10)
		case "int16":
			str = str + strconv.FormatInt(int64(arg.(int16)), 10)
		case "float64":
			str = str +  strconv.FormatFloat(arg.(float64), 'f', -1, 64)
		case "float32":
			str = str +  strconv.FormatFloat(arg.(float64), 'f', -1, 32)
		case "[]uint8":
			str =str + string(arg.([]byte)[:])
		case  "bool":
			if arg.(bool) {
				str =str + ""
			}else{
				str =str + "1"

			}
		}
	}

	return
}


func (T *ToSring) StrSubtract(args ...interface{}) (str string) {
	str = ""
	tmp := ""

	for k, arg := range args {
		switch T.TypeOf(arg) {
		case "nil":
			tmp =  ""
		case "uint":
			tmp =  strconv.Itoa(T.Int(arg.(uint)))
		case "uint8":
			tmp =   strconv.Itoa(T.Int(arg.(uint8)))
		case "uint32":
			tmp =  strconv.Itoa(T.Int(arg.(uint32)))
		case "uint64":
			tmp =  strconv.Itoa(T.Int(arg.(uint64)))
		case "int":
			tmp = strconv.Itoa(arg.(int))
		case "string":
			tmp = arg.(string)
		case "int64":
			tmp = strconv.FormatInt(arg.(int64), 10)
		case "int32":
			tmp = strconv.FormatInt(arg.(int64), 10)
		case "int16":
			tmp = strconv.FormatInt(int64(arg.(int16)), 10)
		case "float64":
			tmp =  strconv.FormatFloat(arg.(float64), 'f', 2, 64)
		case "float32":
			tmp =  strconv.FormatFloat(arg.(float64), 'f', -1, 32)
		case "[]uint8":
			tmp = string(arg.([]byte)[:])
		case  "bool":
			if arg.(bool) {
				tmp = "true"
			}else{
				tmp = "false"

			}
		}
		if k == 0 {
			str = tmp
		}else{
			str = strings.Replace(str , tmp, "", 1)
		}
	}
	return
}



//转Bool
func (T *ToSring) Bool(v interface{}) (returnVar bool) {
	varType := T.TypeOf(v)
	switch varType {
	case "nil":
		returnVar =  false
	case "uint":
		if int( v.(uint)) == 0{
			returnVar =  false
		}else{
			returnVar = true
		}
	case "uint8":
		if int( v.(uint8)) == 0{
			returnVar =  false
		}else{
			returnVar = true
		}
	case "uint32":
		if int( v.(uint32)) == 0{
			returnVar =  false
		}else{
			returnVar = true
		}
	case "uint64":
		if int( v.(uint64)) == 0{
			returnVar =  false
		}else{
			returnVar = true
		}
	case "int":
		if int( v.(int)) == 0{
			returnVar =  false
		}else{
			returnVar = true
		}
	case "int64":
		if v.(int64) == 0{
			returnVar =  false
		}else{
			returnVar = true
		}
	case "int32":
		if v.(int32) == 0{
			returnVar =  false
		}else{
			returnVar = true
		}
	case "int16":
		if v.(int16) == 0{
			returnVar =  false
		}else{
			returnVar = true
		}
	case "string":

		if v.(string) == "FALSE" || v.(string) == "false" || v.(string) == "0" || v.(string) ==""{
			returnVar =  false
		}else{
			returnVar = true
		}

	case "float64":
		if v.(float64) == 0{
			returnVar =  false
		}else{
			returnVar = true
		}
	case "float32":
			if v.(float32) == 0{
				returnVar =  false
			}else{
				returnVar = true
			}
	case  "bool":
		if v.(bool) == false {
			returnVar = false
		}else{
			returnVar = true
		}
	}
	return
}


//转str
func (T *ToSring) Int(v interface{}) (returnVar int) {
	varType := T.TypeOf(v)
	switch varType {
	case "nil":
		returnVar =  0
	case "uint":
		returnVar = int( v.(uint))
	case "uint8":
		returnVar= int( v.(uint8))
	case "uint32":
		returnVar= int( v.(uint32))
	case "uint64":
		returnVar= int( v.(uint64))
	case "int":
		returnVar = v.(int)
	case "int64":
		returnVar,_ = strconv.Atoi(strconv.FormatInt(v.(int64), 10))
	case "int16":
		returnVar = int(v.(int16))
	case "int32":
		returnVar = int(v.(int32))
	case "string":
		returnVar,_ = strconv.Atoi(v.(string))
	case "float64":
		returnVar, _ =    strconv.Atoi( strconv.FormatFloat(v.(float64), 'f', -1, 64))
	case "float32":
		returnVar, _ =  strconv.Atoi( strconv.FormatFloat(v.(float64), 'f', -1, 32))
	case  "bool":
		if v.(bool) == false {
			returnVar = 0
		}else{
			returnVar = 1
		}
	}
	return
}


//转str
func (T *ToSring) Int64(v interface{}) (returnVar int64) {
	varType := T.TypeOf(v)
	switch varType {
	case "nil":
		returnVar = int64(0)
	case "int":
		returnVar =  int64(v.(int))
	case "int64":
		returnVar,_ = v.(int64)
	case "int16":
		returnVar = int64(v.(int16))
	case "int32":
		returnVar = int64(v.(int32))
	case "string":
		returnVar,_ = strconv.ParseInt(v.(string), 10, 64)
	case "float64":
		returnVar =  int64( math.Floor( v.(float64)))
	case "float32":
		returnVar = int64( math.Floor( v.(float64)))
	case  "bool":
		if v.(bool) == false {
			returnVar =int64(0)
		}else{
			returnVar = int64(1)
		}
	}
	return
}

//转str
func (T *ToSring) Int16(v interface{}) (returnVar int16) {
	varType := T.TypeOf(v)
	switch varType {
	case "nil":
		returnVar = int16(0)
	case "int":
		returnVar =  int16(v.(int))
	case "int16":
		returnVar,_ = v.(int16)
	case "int32":
		returnVar =  int16(v.(int32))
	case "int64":
		returnVar = int16(v.(int64))
	case "string":
		rVTmp,_ := strconv.Atoi(v.(string))
		returnVar = int16( rVTmp)
	case "float64":
		returnVar =  int16( math.Floor( v.(float64)))
	case "float32":
		returnVar = int16( math.Floor( v.(float64)))
	case  "bool":
		if v.(bool) == false {
			returnVar =int16(0)
		}else{
			returnVar = int16(1)
		}
	}
	return
}
/*

	case "int16":
		returnVar = int(v.(int16))
	case "int32":
		returnVar = int(v.(int32))
 */
//转float64
func (T *ToSring) Float64(v interface{}) (returnVar float64) {
	varType := T.TypeOf(v)
	switch varType {
	case "nil":
		returnVar,_ =  strconv.ParseFloat(T.Str(0), 64)
	case "int":
		returnVar,_ =  strconv.ParseFloat(T.Str(v.(int)), 64)
	case "int64":
		returnVar,_ =strconv.ParseFloat(T.Str(v.(int64)), 64)
	case "int32":
		returnVar,_ =strconv.ParseFloat(T.Str(v.(int32)), 64)
	case "int16":
		returnVar,_ =strconv.ParseFloat(T.Str(v.(int16)), 64)
	case "string":
		returnVar,_ =  strconv.ParseFloat(v.(string), 64)
	case "float64":
		returnVar = v.(float64)
	case "float32":
		returnVar,_ = strconv.ParseFloat(T.Str(v.(float32)), 64)
	case  "bool":
		if v.(bool) == false {
			returnVar,_ =  strconv.ParseFloat(T.Str(0), 64)
		}else{
			returnVar,_ =  strconv.ParseFloat(T.Str(1), 64)
		}
	}
	return
}

func (T *ToSring)StrRound(str string, prec int, round bool) float64 {
	f,_ := strconv.ParseFloat(str,64)
	return T.Precision(f,prec,round)
}
func (T *ToSring)Round(f float64, prec int, round bool) float64 {
	return T.Precision(f,prec,round)
}
func (T *ToSring)Precision(f float64, prec int, round bool) float64 {
	pow10_n := math.Pow10(prec)
	if round {
		return math.Trunc((f + 0.5/pow10_n)*pow10_n)/pow10_n
	}
	return math.Trunc((f)*pow10_n) / pow10_n
}

/*

	case "int16":
		returnVar = int(v.(int16))
	case "int32":
		returnVar = int(v.(int32))
*/

//加
func (T *ToSring) MathAdd(args ...interface{}) (num float64) {
	num = 0.00
	var tmp float64

	for k, arg := range args {
		switch T.TypeOf(arg) {
		case "nil":
			tmp =  T.Float64(0)
		case "uint":
			tmp =  T.Float64(arg.(uint))
		case "uint8":
			tmp =  T.Float64(arg.(uint8))
		case "uint32":
			tmp =  T.Float64(arg.(uint32))
		case "uint64":
			tmp =  T.Float64(arg.(uint64))
		case "int":
			tmp =  T.Float64(arg.(int))
		case "string":
			tmp =  T.Float64(arg.(string))
		case "int64":
			tmp =  T.Float64(arg.(int64))
		case "int32":
			tmp =  T.Float64(arg.(int32))
		case "int16":
			tmp =  T.Float64(arg.(int16))
		case "float64":
			tmp =  T.Float64(arg.(float64))
		case "float32":
			tmp =  T.Float64(arg.(float32))
		}
		if k == 0 {
			num = tmp
		}else{
			num = num + tmp
		}
	}
	return
}

/*

	case "int16":
		returnVar = int(v.(int16))
	case "int32":
		returnVar = int(v.(int32))
*/
//减
func (T *ToSring) MathSubtract(args ...interface{}) (num float64) {
	num = 0.00
	var tmp float64

	for k, arg := range args {
		switch T.TypeOf(arg) {
		case "nil":
			tmp =  T.Float64(0)
		case "uint":
			tmp =  T.Float64(arg.(uint))
		case "uint8":
			tmp =  T.Float64(arg.(uint8))
		case "uint32":
			tmp =  T.Float64(arg.(uint32))
		case "uint64":
			tmp =  T.Float64(arg.(uint64))
		case "int":
			tmp =  T.Float64(arg.(int))
		case "string":
			tmp =  T.Float64(arg.(string))
		case "int64":
			tmp =  T.Float64(arg.(int64))
		case "int32":
			tmp =  T.Float64(arg.(int32))
		case "int16":
			tmp =  T.Float64(arg.(int16))
		case "float64":
			tmp =  T.Float64(arg.(float64))
		case "float32":
			tmp =  T.Float64(arg.(float32))
		}
		if k == 0 {
			num = tmp
		}else{
			num = num - tmp
		}
	}
	return
}

// 乘

func (T *ToSring) MathMultiply(args ...interface{}) (num float64) {
	num = 0.00
	var tmp float64

	for k, arg := range args {
		switch T.TypeOf(arg) {
		case "nil":
			tmp =  T.Float64(0)
		case "uint":
			tmp =  T.Float64(arg.(uint))
		case "uint8":
			tmp =  T.Float64(arg.(uint8))
		case "uint32":
			tmp =  T.Float64(arg.(uint32))
		case "uint64":
			tmp =  T.Float64(arg.(uint64))
		case "int":
			tmp =  T.Float64(arg.(int))
		case "string":
			tmp =  T.Float64(arg.(string))
		case "int64":
			tmp =  T.Float64(arg.(int64))
		case "int32":
			tmp =  T.Float64(arg.(int32))
		case "int16":
			tmp =  T.Float64(arg.(int16))
		case "float64":
			tmp =  T.Float64(arg.(float64))
		case "float32":
			tmp =  T.Float64(arg.(float32))
		}
		if k == 0 {
			num = tmp
		}else{
			num = num * tmp
		}
	}
	return
}

// 除divide
func (T *ToSring) MathDivide(args ...interface{}) (num float64) {
	num = 0.00
	var tmp float64

	for k, arg := range args {
		switch T.TypeOf(arg) {
		case "nil":
			tmp =  T.Float64(0)
		case "uint":
			tmp =  T.Float64(arg.(uint))
		case "uint8":
			tmp =  T.Float64(arg.(uint8))
		case "uint32":
			tmp =  T.Float64(arg.(uint32))
		case "uint64":
			tmp =  T.Float64(arg.(uint64))
		case "int":
			tmp =  T.Float64(arg.(int))
		case "string":
			tmp =  T.Float64(arg.(string))
		case "int64":
			tmp =  T.Float64(arg.(int64))
		case "int32":
			tmp =  T.Float64(arg.(int32))
		case "int16":
			tmp =  T.Float64(arg.(int16))
		case "float64":
			tmp =  T.Float64(arg.(float64))
		case "float32":
			tmp =  T.Float64(arg.(float32))
		}
		if k == 0 {
			num = tmp
		}else{
			num = num / tmp
		}
	}
	return
}



// 幂运算
func (T *ToSring) MathPow(args ...interface{}) (num float64) {
	num = 0.00
	var tmp float64

	for k, arg := range args {
		switch T.TypeOf(arg) {
		case "nil":
			tmp =  T.Float64(0)
		case "uint":
			tmp =  T.Float64(arg.(uint))
		case "uint8":
			tmp =  T.Float64(arg.(uint8))
		case "uint32":
			tmp =  T.Float64(arg.(uint32))
		case "uint64":
			tmp =  T.Float64(arg.(uint64))
		case "int":
			tmp =  T.Float64(arg.(int))
		case "string":
			tmp =  T.Float64(arg.(string))
		case "int64":
			tmp =  T.Float64(arg.(int64))
		case "int32":
			tmp =  T.Float64(arg.(int32))
		case "int16":
			tmp =  T.Float64(arg.(int16))
		case "float64":
			tmp =  T.Float64(arg.(float64))
		case "float32":
			tmp =  T.Float64(arg.(float32))
		}
		if k == 0 {
			num = tmp
		}else{
			num = math.Pow(num, tmp)
		}
	}
	return
}




// 取余数
func (T *ToSring) MathRemainder(args ...interface{}) (num int) {
	num = 0
	var tmp int

	for k, arg := range args {
		switch T.TypeOf(arg) {
		case "nil":
			tmp =  T.Int(0)
		case "uint":
			tmp =  T.Int(arg.(uint))
		case "uint8":
			tmp =  T.Int(arg.(uint8))
		case "uint32":
			tmp =  T.Int(arg.(uint32))
		case "uint64":
			tmp =  T.Int(arg.(uint64))
		case "int":
			tmp =  T.Int(arg.(int))
		case "string":
			tmp =  T.Int(arg.(string))
		case "int64":
			tmp =  T.Int(arg.(int64))
		case "int32":
			tmp =  T.Int(arg.(int32))
		case "int16":
			tmp =  T.Int(arg.(int16))
		case "float64":
			tmp =  T.Int(arg.(float64))
		case "float32":
			tmp =  T.Int(arg.(float32))
		}
		if k == 0 {
			num = tmp
		}else{
			num =  num % tmp
		}
	}
	return
}

func (T *ToSring) Byte(arg interface{}) (bytes []byte) {
	tmp := ""
	switch T.TypeOf(arg) {
	case "nil":
		tmp =  T.Str(0)
	case "uint":
		tmp =  T.Str(arg.(uint))
	case "uint8":
		tmp =  T.Str(arg.(uint8))
	case "uint32":
		tmp =  T.Str(arg.(uint32))
	case "uint64":
		tmp =  T.Str(arg.(uint64))
	case "int":
		tmp =  T.Str(arg.(int))
	case "string":
		tmp =  T.Str(arg.(string))
	case "int64":
		tmp =  T.Str(arg.(int64))
	case "int32":
		tmp =  T.Str(arg.(int32))
	case "int16":
		tmp =  T.Str(arg.(int16))
	case "float64":
		tmp =  T.Str(arg.(float64))
	case "float32":
		tmp =  T.Str(arg.(float32))
	case "bool":
		if arg.(bool) == false {
			tmp = ""
		}else if  arg.(bool) == true{
			tmp = "1"
		}
	}
	bytes  = []byte(tmp)
	return
}

func (T *ToSring) Uint(arg interface{}) (uints uint) {
	str := T.Str(arg)
	_ = T.StrToUint(str, &uints)
	return
}
func (T *ToSring) Uint8(arg interface{}) (uints uint8) {
	str := T.Str(arg)
	_ = T.StrToUint(str, &uints)
	return
}
func (T *ToSring) Uint16(arg interface{}) (uints uint16) {
	str := T.Str(arg)
	_ = T.StrToUint(str, &uints)
	return
}
func (T *ToSring) Uint32(arg interface{}) (uints uint32) {
	str := T.Str(arg)
	_ = T.StrToUint(str, &uints)
	return
}

func (T *ToSring)StrToUint(strNumber string, value interface{}) (err error) {
	var number interface{}
	number, err = strconv.ParseUint(strNumber, 10, 64)
	switch v := number.(type) {
	case uint64:
		switch d := value.(type) {
		case *uint64:
			*d = v
		case *uint:
			*d = uint(v)
		case *uint16:
			*d = uint16(v)
		case *uint32:
			*d = uint32(v)
		case *uint8:
			*d = uint8(v)
		}
	}
	return
}


func (T *ToSring) Time(arg interface{}) (times time.Time) {
	tmp := ""
	switch T.TypeOf(arg) {
	case "nil":
		tmp =  T.Str(0)
	case "uint":
		tmp =  T.Str(arg.(uint))
	case "uint8":
		tmp =  T.Str(arg.(uint8))
	case "uint32":
		tmp =  T.Str(arg.(uint32))
	case "uint64":
		tmp =  T.Str(arg.(uint64))
	case "int":
		tmp =  T.Str(arg.(int))
	case "string":
		tmp =  T.Str(arg.(string))
	case "int64":
		tmp =  T.Str(arg.(int64))
	case "int32":
		tmp =  T.Str(arg.(int32))
	case "int16":
		tmp =  T.Str(arg.(int16))
	case "float64":
		tmp =  T.Str(arg.(float64))
	case "float32":
		tmp =  T.Str(arg.(float32))
	case "bool":
		if arg.(bool) == false {
			tmp = ""
		}else if  arg.(bool) == true{
			tmp = "1"
		}
	}
	times,_  = time.Parse("2006-01-02 15:04:05",tmp)
	return
}



//弱类型转换 喵
func (T *ToSring) IsSame(args ...interface{}) (isSame bool) {
	isSame = false
	tmp :=""
	tmpMaster :=""

	var tmpF string
	var tmpFP string
	var tmpFv float64
	var tmpFPv float64


	defer func() {
		err :=recover()
		if err != nil{
			isSame = false
			return
		}
	}()

	for k, arg := range args {
		if k > 0 {
			tmpFP = tmpF
			tmpFPv = tmpFv
		}
		switch T.TypeOf(arg) {
		case "nil":
			tmp =  ""
			tmpF = "str"
			tmpFv = 1

		case "uint":
			tmp =  T.Str(arg.(uint))
			tmpF = "uint"
			tmpFv = T.Float64(arg.(uint))

		case "uint8":
			tmp =  T.Str(arg.(uint8))
			tmpF = "uint8"
			tmpFv = T.Float64(arg.(uint8))

		case "uint32":
			tmp =  T.Str(arg.(uint32))
			tmpF = "uint32"
			tmpFv = T.Float64(arg.(uint32))

		case "uint64":
			tmp =  T.Str(arg.(uint64))
			tmpF = "uint64"
			tmpFv = T.Float64(arg.(uint64))

		case "int":
			tmp =  T.Str(arg.(int))
			tmpF = "int"
			tmpFv = T.Float64(arg.(int))

		case "string":
			tmp =  T.Str(arg.(string))
			tmpF = "string"
			tmpFv = T.Float64(arg.(string))

		case "int64":
			tmp =  T.Str(arg.(int64))
			tmpF = "int64"
			tmpFv = T.Float64(arg.(int64))

		case "int32":
			tmp =  T.Str(arg.(int32))
			tmpF = "int32"
			tmpFv = T.Float64(arg.(int32))

		case "float64":
			tmp =  T.Str(arg.(float64))
			tmpF = "float64"
			tmpFv = T.Float64(arg.(float64))

		case "float32":
			tmp =  T.Str(arg.(float32))
			tmpF = "float32"
			tmpFv = T.Float64(arg.(float32))

		case "bool":
			tmpF = "bool"
			tmpFv = T.Float64(arg.(bool))

			if arg.(bool) == false{
				tmp = ""
			}else{
				tmp = "1"
			}
		}
		if k == 0 {
			tmpMaster = tmp
		}else if tmpF == "float64" || tmpFP == "float64"||tmpF == "int" || tmpFP == "int"||
			tmpF == "float32" || tmpFP == "float32"||tmpF == "int64" || tmpFP == "int64"||tmpF == "int32" || tmpFP == "int32"||tmpF == "bool" || tmpFP == "bool"{
			if T.Float64(tmp) == T.Float64(tmpMaster){
				isSame = true
			}else if T.Float64(tmp) != T.Float64(tmpMaster){
				isSame = false
				break
			}
		}else if tmpFv != 0 || tmpFPv != 0{
			if T.Float64(tmp) == T.Float64(tmpMaster){
				isSame = true
			}else if T.Float64(tmp) != T.Float64(tmpMaster){
				isSame = false
				break
			}
		}else if tmp == tmpMaster{
			isSame = true
		}else if tmp != tmpMaster{
			isSame = false
			break
		}
	}
	return
}



func (T *ToSring) Split(arg string ,l string)(value  []string){
	value = make([]string,0)
	if arg == "" {
		return
	}
	value = strings.Split(arg, l)
	return
}


func (T *ToSring)Json(mapInstances interface{}) (jsonStr string, err error)  {
	jsonStrs, err := json.Marshal(mapInstances)
	jsonStr = ""
	if err != nil {
		return
	}
	jsonStr = T.Str(jsonStrs)
	return
}



func (T *ToSring)JsonToMaps(mapInstances interface{}) (mapInstance []map[string]string, err error)  {
	jsonStrs, err := json.Marshal(mapInstances)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonStrs, &mapInstance)
	return
}
func (T *ToSring)JsonToInterface(mapInstances interface{}) (mapInstance []map[string]interface{}, err error)  {
	jsonStr := T.Byte(mapInstances)
	err = json.Unmarshal(jsonStr, &mapInstance)
	return
}


func (T *ToSring)JsonToMap(mapInstances interface{}) (mapInstance map[string]string, err error)  {
	jsonStrs, err := json.Marshal(mapInstances)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonStrs, &mapInstance)
	return
}

//