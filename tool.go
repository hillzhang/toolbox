package utils
import (
	uuid "log_collection/utils/satori/go.uuid"
	"crypto/md5"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

type Res struct {
	Area string `json:"area"`
	Location string `json:"location"`
}
type Response struct {
	ResultCode string `json:"resultcode"`
	Reason string `json:"reason"`
	Result *Res `json:"result"`
}
func Md5Encode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateUUID() string {
	sig := uuid.NewV1().String()
	sig = strings.Replace(sig, "-", "", -1)
	return sig
}

func CheckIpLocation(ip string) string{
	defer func(){
		if r := recover(); r != nil{
			log.Println("get ip location error")
		}
	}()
	response, err := http.Get(fmt.Sprintf("http://apis.juhe.cn/ip/ip2addr?ip=%s&key=116199a801a395a94d79fcd7e0cd9eb9",ip))
	if err != nil {
		return " "
	}
	b, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var res Response
	err = json.Unmarshal([]byte(string(b)),&res)
	if err != nil {
		return " "
	}
	return res.Result.Area
}

