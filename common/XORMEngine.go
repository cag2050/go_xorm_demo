package common

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func CreateXORMEngine() (*xorm.Engine, error) {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v",
		"root",
		"123456",
		"10.1.106.161",
		"3306",
		"openser",
		"utf8")
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return nil, fmt.Errorf("create failed, error:%v, dataSourceName:%v", err, dataSourceName)
	}
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	engine.SetLogLevel(core.LOG_DEBUG)
	return engine, nil
}
