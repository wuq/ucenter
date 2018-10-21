package redis

import(
	"hash/crc32"
	"encoding/json"
	"io/ioutil"

	"github.com/micro/go-log"
	"github.com/gomodule/redigo/redis"
)

var (
	conf map[string]interface{}
	hostLen int8
	redisClient  []*redis.Pool
	err error
)

func init(){

	bytes, err := ioutil.ReadFile("./config/redis.dev.json")
    if err != nil {
        log.Fatal("Read redis config: ", err.Error())
    }
 
    if err := json.Unmarshal(bytes, &conf); err != nil {
        log.Fatal("Unmarshal redis config: ", err.Error())
    }

	hostLen     = len(conf["host"].([]interface{}))
	redisClient = make([]*redis.Pool, hostLen)

	for k,v := range conf["host"].([]interface{}){
		redisClient[k] = &redis.Pool{
			MaxIdle: conf["maxidle"],
			MaxActive: conf["maxactive"],
			IdleTimeout: conf["timeout"],
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", v)
				if err != nil {
					return nil, err
				}
				c.Do("AUTH", conf["auth"])
				return c, nil
			},
		}
	}
}


//执行命令
func CallDo(cmd string, args ...interface{}) (interface{}, error){
	k := getConn(cmd)
	//拿取链接
	rc := redisClient[k].Get()
	defer rc.Close()

	return rc.Do(cmd, args...)
}


//获取节点
func getConn(cmd string) int{
	return int(crc32.ChecksumIEEE([]byte(cmd))) % hostLen
}