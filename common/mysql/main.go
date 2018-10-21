package mysql

import(
	"fmt"
	"encoding/json"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/micro/go-log"
)

const SUB_TABLE_NUM uint64 = 1
const MC_NULL_VALUE string = "nullStr"

var (
	conf map[string]interface{}
	DB *xorm.EngineGroup
	conns []string
	err error
)

/**
 *	1.xorm.RandomPolicy()随机访问负载均衡,
 *	2.xorm.WeightRandomPolicy([]int{2, 3,4})权重随机负载均衡
 *	3.xorm.RoundRobinPolicy() 轮询访问负载均衡
 *	4.xorm.WeightRoundRobinPolicy([]int{2, 3,4}) 权重轮训负载均衡
 *	5.xorm.LeastConnPolicy()最小连接数负载均衡
 */


func init(){

	bytes, err := ioutil.ReadFile("./config/mysql.dev.json")
    if err != nil {
        log.Fatal("Read mysql config: ", err.Error())
    }
 
    if err := json.Unmarshal(bytes, &conf); err != nil {
        log.Fatal("Unmarshal mysql config: ", err.Error())
    }

	tpl := fmt.Sprintf("%s:%s@tcp(%%s)/%s?charset=%s", conf["user"], conf["pass"], conf["dbname"], conf["charset"])

	conns = make([]string, 0)
	conns = append(conns, fmt.Sprintf(tpl, conf["masterHost"]))
	for _, v := range conf["slaveHost"].([]interface{}){
		conns = append(conns, fmt.Sprintf(tpl, v))
	}

    DB, err = xorm.NewEngineGroup("mysql", conns, xorm.RoundRobinPolicy())
    if err != nil{
    	log.Log("mysql conn error:", err)
    }

}
