package str

import (
	"time"
	"strconv"
)

func GetToken() string{
	now := time.Now().UnixNano()
	format := strconv.FormatInt(now,10)
	return Md5Encode(format)
}