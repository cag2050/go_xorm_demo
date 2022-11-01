package role_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	"testing"
)

func TestFind(t *testing.T) {
	//name := "role"
	pageSize := 5
	pageIndex := 2

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	roles := make([]model.Role, 0)
	//if err := engine.Where("name like ?", "%"+name+"%").Limit(pageSize, (pageIndex-1)*pageSize+1).Find(&roles); err != nil {
	if err := engine.Limit(pageSize, (pageIndex-1)*pageSize).Find(&roles); err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}
	fmt.Println(fmt.Sprintf("%+v", len(roles)))
	fmt.Println(fmt.Sprintf("%+v", roles))
}
