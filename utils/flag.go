package utils

import (
	"reflect"
	"unsafe"
	"flag"
	"strings"
)
//传入格式
//usage 对应flag 参数中的usage，
//name  对应flag 参数中的name

//type Config struct {
//	Addr string `usage:"connection addr" name:"a"`
//	Port string `usage:"connection port" name:"p"`
//	Debug bool `usage:"debug module" name:"d"`
//	Int64 int64 `usage:"param for init64" name:"int64"`
//	Float64 float64 `usage:"param for float64" name:"float64"`
//	Uint uint `usage:"param for uint" name:"uint"`
//	Uint64 uint64 `usage:"param for uint64" name:"uint64"`
//}

const(
	USAGE = "usage"
	NAME = "name"
)

// parsed 为 true 时自动解析flag，否则需要手动解析
func GenFlag(config interface{},parsed bool)  {
	var valueOfConfig reflect.Value
	var typeOfConfig reflect.Type

	if reflect.TypeOf(config).Kind() != reflect.Ptr{
		panic("解析对象必须为指针类型")
	}else {
		valueOfConfig = reflect.ValueOf(config).Elem()
		typeOfConfig = reflect.TypeOf(config).Elem()
	}

	for i := 0; i< typeOfConfig.NumField();i++{

		tag := typeOfConfig.Field(i).Tag
		var name,usage string

		name,ok := tag.Lookup(NAME)
		if !ok {
			name = strings.ToLower(typeOfConfig.Field(i).Name)
		}
		usage,ok = tag.Lookup(USAGE)
		if !ok{
			usage = strings.ToLower(typeOfConfig.Field(i).Name)
		}
		
		unPointer := unsafe.Pointer(valueOfConfig.Field(i).Addr().Pointer())
		value := valueOfConfig.Field(i)

		switch typeOfConfig.Field(i).Type.Kind(){
		case reflect.String:
			flag.StringVar((*string)(unPointer),name,value.String(),usage)
		case reflect.Int:
			flag.IntVar((*int)(unPointer),name,int(value.Int()),usage)
		case reflect.Int64:
			flag.Int64Var((*int64)(unPointer),name,value.Int(),usage)
		case reflect.Bool:
			flag.BoolVar((*bool)(unPointer),name,value.Bool(),usage)
		case reflect.Float64:
			flag.Float64Var((*float64)(unPointer),name,value.Float(),usage)
		case reflect.Uint:
			flag.UintVar((*uint)(unPointer),name,uint(value.Uint()),usage)
		case reflect.Uint64:
			flag.Uint64Var((*uint64)(unPointer),name,value.Uint(),usage)
		}
	}
	if parsed{
		flag.Parse()
	}
}