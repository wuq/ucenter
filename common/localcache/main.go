package memcache

import(
	"strconv"
//	"encoding/json"
//	"io/ioutil"


	"github.com/bradfitz/gomemcache/memcache"
    "github.com/micro/go-log"
)

var (
	err error
	Cli *memcache.Client
)



func init(){

	Cli = memcache.New("127.0.0.1:11211")
	if Cli == nil{
    	log.Fatal("mc conn error:", Cli)
	}
	Cli.MaxIdleConns = 10
}


func Set(k string, v interface{}) error{

	var s string 
	switch v.(type){
		case nil:
			return nil
		case int8:
			s = strconv.FormatInt(int64(v.(int8)), 10)
		case int16:
			s = strconv.FormatInt(int64(v.(int16)), 10)
		case int32:
			s = strconv.FormatInt(int64(v.(int32)), 10)
		case int64: 
			s = strconv.FormatInt(v.(int64), 10)
		case uint8:
			s = strconv.FormatUint(uint64(v.(uint8)), 10)
		case uint16:
			s = strconv.FormatUint(uint64(v.(uint16)), 10)
		case uint32:
			s = strconv.FormatUint(uint64(v.(uint32)), 10)
		case  uint64:
			s = strconv.FormatUint(v.(uint64), 10)
		case bool:
			s = strconv.FormatBool(v.(bool))
		case float32:
			s = strconv.FormatFloat(float64(v.(float32)), 'E', 3, 32)
		case float64:
			s = strconv.FormatFloat(v.(float64), 'E', 3, 64)
		case string:
			s = v.(string)
		default:
			return nil
	}
	return Cli.Set(&memcache.Item{Key: k, Value: []byte(s), Expiration: 5})
}


func Get(key string) (string, error){
	it, err := Cli.Get(key)
	if err != nil {
		return "", err
	}
    if string(it.Key) == key {
       return string(it.Value), nil
    }
    return "", nil
}

/**

//set key-value
Cli.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

//get key-value
it, _ := Cli.Get("foo")
if string(it.Key) == "foo" {
    fmt.Println("value is ", string(it.Value))
} else {
   fmt.Println("Get failed")
}


*/