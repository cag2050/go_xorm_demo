package common

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/log"
	"github.com/cag2050/go_xorm_demo/utildb"
	"xorm.io/xorm"
	xormLog "xorm.io/xorm/log"
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

	logger := log.StandardLogger()
	engine.SetLogger(utildb.NewXormLogger("", logger))

	engine.ShowSQL(true)
	engine.SetLogLevel(xormLog.LOG_DEBUG)

	if pingErr := engine.Ping(); pingErr != nil {
		return nil, pingErr
	}

	return engine, nil
}
