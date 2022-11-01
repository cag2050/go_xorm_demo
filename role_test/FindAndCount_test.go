package role_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	"testing"
)

func TestFindAndCount(t *testing.T) {
	name := "role"
	var pageSize int64 = 2
	var pageIndex int64 = 1

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	roles := make([]model.Role, 0)
	//count, err1 := engine.Limit(pageSize, (pageIndex-1)*pageSize).FindAndCount(&roles)
	// 返回总记录条数的同时，进行分页
	session := engine.Where("1=1")
	if len(name) > 0 {
		session = session.Where("name like ?", "%"+name+"%")
	}
	//count, err1 := engine.Where("name like ?", "%"+name+"%").Limit(pageSize, (pageIndex-1)*pageSize).FindAndCount(&roles)
	count, err1 := session.Limit(int(pageSize), int((pageIndex-1)*pageSize)).FindAndCount(&roles)
	if err1 != nil {
		fmt.Println(fmt.Sprintf("%+v", err1))
	}
	fmt.Println(fmt.Sprintf("%+v", count))
	fmt.Println(fmt.Sprintf("%+v", roles))
}
