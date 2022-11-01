package user_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	"testing"
)

func TestUnscopedGet(t *testing.T) {
	// 测试的id值
	id := 6

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	user := new(model.User)

	// 注意：Unscoped().Get(user) 和 Get(user) 的区别
	// [SQL] SELECT `id`, `name`, `account`, `password`, `mobile`, `state`, `created_at`, `updated_at`, `deleted_at` FROM `exporter_user` WHERE `id`=? LIMIT 1 []interface {}{6}
	b1, getErr1 := engine.ID(id).Unscoped().Get(user)
	if getErr1 != nil {
		fmt.Println(fmt.Sprintf("%+v", getErr1))
	}
	fmt.Println(fmt.Sprintf("%+v", b1))
	fmt.Println(fmt.Sprintf("%+v", user))
}
