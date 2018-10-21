package common

import(
//	"hash/crc32"
//	"io/ioutil"
	"encoding/json"

)

func Rsp(code int, msg error, data interface{}) string{
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"]  = msg.Error()
	m["data"] = data

	str, _ := json.Marshal(m)
	return string(str)
}