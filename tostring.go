package simple_variable

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ToString struct {

}

//获取类型
func (T *ToString) TypeOf(v interface{}) (t string ){
	t = "nil"
	if v != nil{
		t = reflect.TypeOf(v).String()
	}
	return  t
}
func (T *ToString)Shell(s string) (string, error){
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()


	return out.String(), err
}

//转str
func (T *ToString) Str(v interface{}) (returnVar  string) {
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
	case "[]byte":
		returnVar = T.ByteString(v.([]byte))
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

func (T *ToString) StrAdd(args ...interface{}) (str string) {
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
// String 将 `[]byte` 转换为 `string`
func (T *ToString)ByteString(b []byte) string {
	for idx, c := range b {
		if c == 0 {
			return string(b[:idx])
		}
	}
	return string(b)
}

// StringWithoutZero 将 `[]byte` 转换为 `string`
func (T *ToString)ByteStringWithoutZero(b []byte) string {
	s := make([]rune, len(b))
	offset := 0
	for i, c := range b {
		if c == 0 {
			offset++
		} else {
			s[i-offset] = rune(c)
		}
	}
	return string(s[:len(b)-offset-1])
}

func (T *ToString) StrSubtract(args ...interface{}) (str string) {
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
func (T *ToString) Bool(v interface{}) (returnVar bool) {
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
func (T *ToString) Int(v interface{}) (returnVar int) {
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
		tmp,_ :=strconv.ParseFloat(v.(string), 64)
		returnVar, _ =    strconv.Atoi( fmt.Sprintf("%1.0f", tmp))
		//returnVar,_ = strconv.Atoi(v.(string))
	case "float64":
		//returnVar, _ =    strconv.Atoi( strconv.FormatFloat(v.(float64), 'f', -1, 64))
		returnVar, _ =    strconv.Atoi( fmt.Sprintf("%1.0f", v.(float64)))

	case "float32":
		//returnVar, _ =  strconv.Atoi( strconv.FormatFloat(v.(float64), 'f', -1, 32))
		returnVar, _ =    strconv.Atoi( fmt.Sprintf("%1.0f", v.(float32)))

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
func (T *ToString) Int64(v interface{}) (returnVar int64) {
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
		returnVar = int64( T.Int(v.(string)))
	case "float64":
		//returnVar =  int64( math.Floor( v.(float64)))

		returnVar =  int64( T.Int(v.(float64)))

	case "float32":
		//returnVar = int64( math.Floor( v.(float64)))
		returnVar =  int64( T.Int(v.(float32)))

	case  "bool":
		if v.(bool) == false {
			returnVar =int64(0)
		}else{
			returnVar = int64(1)
		}
	}
	return
}

//Int16
func (T *ToString) Int16(v interface{}) (returnVar int16) {
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
		returnVar = int16( T.Int(v.(string)))

	case "float64":
		//returnVar =  int16( math.Floor( v.(float64)))
		returnVar =  int16( T.Int(v.(float64)))
	case "float32":
		//returnVar = int16( math.Floor( v.(float64)))
		returnVar =  int16( T.Int(v.(float64)))

	case  "bool":
		if v.(bool) == false {
			returnVar =int16(0)
		}else{
			returnVar = int16(1)
		}
	}
	return
}


//Int8
func (T *ToString) Int8(v interface{}) (returnVar int8) {
	varType := T.TypeOf(v)
	switch varType {
	case "nil":
		returnVar = int8(0)
	case "int":
		returnVar =  int8(v.(int))
	case "int8":
		returnVar,_ = v.(int8)
	case "int32":
		returnVar =  int8(v.(int32))
	case "int64":
		returnVar = int8(v.(int64))
	case "string":
		returnVar = int8( T.Int(v.(string)))
	case "float64":
		returnVar = int8( T.Int(v.(float64)))
	case "float32":
		returnVar = int8( T.Int(v.(float32)))
	case  "bool":
		if v.(bool) == false {
			returnVar =int8(0)
		}else{
			returnVar = int8(1)
		}
	}
	return
}

//Int32
func (T *ToString) Int32(v interface{}) (returnVar int32) {
	varType := T.TypeOf(v)
	switch varType {
	case "nil":
		returnVar = int32(0)
	case "int":
		returnVar =  int32(v.(int))
	case "int16":
		returnVar = int32(v.(int16))
	case "int32":
		returnVar =  int32(v.(int32))
	case "int64":
		returnVar = int32(v.(int64))
	case "string":
		returnVar = int32( T.Int(v.(string)))

	case "float64":
		returnVar = int32( T.Int(v.(float64)))
	case "float32":
		returnVar = int32( T.Int(v.(float32)))
	case  "bool":
		if v.(bool) == false {
			returnVar =int32(0)
		}else{
			returnVar = int32(1)
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
func (T *ToString) Float64(v interface{}) (returnVar float64) {
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

func (T *ToString)StrRound(str string, prec int, round bool) float64 {
	f,_ := strconv.ParseFloat(str,64)
	return T.Precision(f,prec,round)
}
func (T *ToString)Round(f float64, prec int, round bool) float64 {
	return T.Precision(f,prec,round)
}
func (T *ToString)Precision(f float64, prec int, round bool) float64 {
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
func (T *ToString) MathAdd(args ...interface{}) (num float64) {
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
func (T *ToString) MathSubtract(args ...interface{}) (num float64) {
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

func (T *ToString) MathMultiply(args ...interface{}) (num float64) {
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
func (T *ToString) MathDivide(args ...interface{}) (num float64) {
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
func (T *ToString) MathPow(args ...interface{}) (num float64) {
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
func (T *ToString) MathRemainder(args ...interface{}) (num int) {
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

func (T *ToString) Byte(arg interface{}) (bytes []byte) {
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
	case "[]byte":
		tmp =   T.Str(arg.(byte))
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

func (T *ToString) Uint(arg interface{}) (uints uint) {
	str := T.Str(arg)
	_ = T.StrToUint(str, &uints)
	return
}
func (T *ToString) Uint8(arg interface{}) (uints uint8) {
	str := T.Str(arg)
	_ = T.StrToUint(str, &uints)
	return
}
func (T *ToString) Uint16(arg interface{}) (uints uint16) {
	str := T.Str(arg)
	_ = T.StrToUint(str, &uints)
	return
}
func (T *ToString) Uint32(arg interface{}) (uints uint32) {
	str := T.Str(arg)
	_ = T.StrToUint(str, &uints)
	return
}

func (T *ToString)StrToUint(strNumber string, value interface{}) (err error) {
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


func (T *ToString) Time(arg interface{}) (times time.Time) {
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
	local, _ := time.LoadLocation("Local")
	times, _ = time.ParseInLocation("2006-01-02 15:04:05", tmp, local)
	return
}


func (T *ToString) TimeV2(arg interface{}) (times time.Time) {
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
	local, _ := time.LoadLocation("Local")
	times, _ = time.ParseInLocation("20060102", tmp, local)
	return
}

//弱类型转换 喵
func (T *ToString) IsSame(args ...interface{}) (isSame bool) {
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



func (T *ToString) Split(arg string ,l string)(value  []string){
	value = make([]string,0)
	if arg == "" {
		return
	}
	value = strings.Split(arg, l)
	return
}


func (T *ToString)Json(mapInstances interface{}) (jsonStr string, err error)  {
	jsonStrs, err := json.Marshal(mapInstances)
	jsonStr = ""
	if err != nil {
		return
	}
	jsonStr = T.Str(jsonStrs)
	return
}
func (T *ToString)MapJson(mapInstances map[string]interface{}) (jsonStr string, err error)  {
	jsonStrs, err := json.Marshal(mapInstances)
	jsonStr = ""
	if err != nil {
		return
	}
	jsonStr = T.Str(jsonStrs)
	return
}


func (T *ToString)JsonToMaps(mapInstances interface{}) (mapInstance []map[string]string, err error)  {
	jsonStrs, err := json.Marshal(mapInstances)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonStrs, &mapInstance)
	return
}
func (T *ToString)JsonToInterface(mapInstances interface{}) (mapInstance []map[string]interface{}, err error)  {
	jsonStr := T.Byte(mapInstances)
	err = json.Unmarshal(jsonStr, &mapInstance)
	return
}
func (T *ToString)JsonToInterfaceFirst(mapInstances interface{}) (mapInstance map[string]interface{}, err error)  {
	jsonStr := T.Byte(mapInstances)
	err = json.Unmarshal(jsonStr, &mapInstance)
	return
}

func (T *ToString)JsonToInterfaceAll(mapInstances interface{}) (mapInstance interface{}, err error)  {
	jsonStr := T.Byte(mapInstances)
	err = json.Unmarshal(jsonStr, &mapInstance)
	return
}


func (T *ToString)JsonToMap(mapInstances interface{}) (mapInstance map[string]string, err error)  {
	jsonStrs, err := json.Marshal(mapInstances)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonStrs, &mapInstance)
	return
}

//



func (T *ToString)RandString(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)

}


func (T *ToString)CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func (T *ToString)SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

//filterArray
// snake string, XxYy to xx_yy , XxYY to xx_yy
func (T *ToString)FilterArray(arr map[string]interface{}, keys []string, target string  ) (arrs map[string]interface{},err error) {
	err = nil
	ret :=  map[string]interface{}{}
	arrs = arr
	var condition bool
	if len(keys) == 0 {
		err = errors.New("过滤字段不能为空")
		return
	}
	for _,v :=  range keys{
		switch target {
		case "empty":
			if v == ""{
				condition = false
			}else{
				condition = true
			}
		default:
			condition = true
		}
		if condition{
			ret[v] = arr[v]
		}
	}
	arrs = ret
	return
}

func (T *ToString)TimeStamp(datetime string) (timestamp string) {
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestampInt64 := tmp.Unix()
	timestamp= T.Str(timestampInt64)    //转化为时间戳 类型是int64
	return
}
func (T *ToString)SliceToString(Slices []interface{}) (Strs []string) {
	//日期转化为时间戳
	Strs = make([]string,0)
	for _,strategyGroupsAccountStr := range Slices{
		Strs = append(Strs, T.Str(strategyGroupsAccountStr))
	}
	return
}

func (T *ToString)RegBetween(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}

func (T *ToString) HashHmacBySha256(str string, key string)(retStr string , err error)  {
	keys := []byte(key)
	mac := hmac.New(sha256.New, keys)
	mac.Write([]byte(str))
	retStr = T.Str(mac.Sum(nil))
	return
}

func (T *ToString)GetDayDate(date string,day int) (ltDate string){
	dateto := T.TimeV2(date)
	h := T.Str(day * 24)
	d ,_:= time.ParseDuration("+"+h+"h")
	etm := dateto.Add(d)
	ltDate = etm.Format("20060102")
	return ltDate
}
//获取周一0分的时间戳
func (T *ToString)GetMondayTime() (totime int64) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	totime = weekStart.Unix()
	return
}
//获取指定时间戳 到周一 0分的位数
func (T *ToString)GetNowtoMondayTime(NowTime int64,Num int) ( location int) {

	totime := T.GetMondayTime()
	location = T.Int(T.Int(NowTime - totime)/(60*Num))
	return
}
