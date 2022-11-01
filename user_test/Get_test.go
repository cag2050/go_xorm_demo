package user_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	"testing"
)

func TestGet(t *testing.T) {
	// 测试的id值
	id := 6

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	user := new(model.User)

	// 注意：Unscoped().Get(user) 和 Get(user) 的区别
	// [SQL] SELECT `id`, `name`, `account`, `password`, `mobile`, `state`, `created_at`, `updated_at`,
	//`deleted_at` FROM `exporter_user` WHERE (`deleted_at` IS NULL OR `deleted_at`=?) AND `id`=? LIMIT 1 []interface {}{"0001-01-01 00:00:00", 6}
	b1, getErr1 := engine.ID(id).Get(user)
	if getErr1 != nil {
		fmt.Println(fmt.Sprintf("%+v", getErr1))
	}
	fmt.Println(fmt.Sprintf("%+v", b1))
	fmt.Println(fmt.Sprintf("%+v", user))
}
