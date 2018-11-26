package utils

import (
	"testing"
	"fmt"
)


type Config struct {
	Addr string `usage:"connection addr" name:"a"`
	Port string `usage:"connection port" name:"p"`
	Debug bool `usage:"debug module" name:"d"`
	Int64 int64 `usage:"param for init64" name:"int64"`
	Float64 float64 `usage:"param for float64" name:"float64"`
	Uint uint `usage:"param for uint" name:"uint"`
	Uint64 uint64 `usage:"param for uint64" name:"uint64"`
}

func TestGenFlag(t *testing.T) {
	//config 为指针类型
	config := &Config{
		Addr: "127.0.0.1",
		Port: "9998",
		Debug: false,
	}
	//自动解析 parsed=true
	//如果还有其他command 需要搜东解析，这里填false
	//需要手动运行 flag.Parse()
	GenFlag(config,true)
	fmt.Println(*config)

	//GenFlag(config, false)
	//config_file := flag.String("c","cfg.json","configuration file path")
	//flag.Parse()
	//if *config_file != ""{
	//	fmt.Println(config_file)
	//}
}


