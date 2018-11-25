package convert

import "encoding/json"

func ToJsonString(obj interface{})(string,error){
	res,err := json.Marshal(obj)
	if err != nil {
		return "",err
	}
	return string(res),nil
}

func ToJsonByte(obj interface{})([]byte,error){
	res,err := json.Marshal(obj)
	if err != nil {
		return nil,nil
	}
	return res,nil
}

