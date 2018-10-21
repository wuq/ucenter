package model

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"

    "github.com/micro/go-log"
  _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"

	L1 "tenno.ucenter/common/memcache"
)

var err error
var db *xorm.Engine
var conf map[string]interface{}


const (
	//当前最大ID号
	KEY_MC_TICKET_MAXID string = "key_mc_user_maxid"
)



type Ticket struct{
	Id uint64
}


func init(){
	
	bytes, err := ioutil.ReadFile("./config/mysql.dev.json")
    if err != nil {
        log.Fatal("Read mysql config: ", err.Error())
    }
 
    if err := json.Unmarshal(bytes, &conf); err != nil {
        log.Fatal("Unmarshal mysql config: ", err.Error())
    }

	tpl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", conf["user"], conf["pass"], conf["masterHost"], "ticket", conf["charset"])
	log.Log(tpl)
    db, err = xorm.NewEngine("mysql", tpl)
    if err != nil {
    	fmt.Println(err)
    }
}


//取一个新ID
func (t *Ticket)GetId() error{
	log.Log("model Ticket.GetId request")

	_, err := db.Exec("REPLACE INTO ticket(stag) VALUE ('a')")
	if err != nil {
		return err
	}
	results, err := db.Query("SELECT LAST_INSERT_ID()")
	if err != nil {
		return err
	}

	newId := string(results[0]["LAST_INSERT_ID()"])
	id,err := strconv.ParseUint(newId, 10, 64)
	if err != nil{
		return err
	}

	L1.Set(KEY_MC_TICKET_MAXID, id, 0)
	t.Id = id

	return err
}
