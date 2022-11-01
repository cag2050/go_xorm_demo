package user_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	"testing"
)

func TestUpdate(t *testing.T) {
	// 测试的值
	id := 7
	name := "myName7"

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	user := new(model.User)

	// SELECT * FROM user WHERE id = ?
	b1, getErr1 := engine.ID(id).Get(user)
	if getErr1 != nil {
		fmt.Println(fmt.Sprintf("%+v", getErr1))
	}
	fmt.Println(fmt.Sprintf("%+v", b1))
	fmt.Println(fmt.Sprintf("%+v", user))

	// 当传入的为结构体指针时，只有非空和0的field才会被作为更新的字段。https://xorm.io/zh/docs/chapter-06/readme/
	user.Name = name
	affected, updateErr := engine.ID(id).Update(user)
	if updateErr != nil {
		fmt.Println(fmt.Sprintf("%+v", updateErr))
	}
	fmt.Println(fmt.Sprintf("%+v", affected))

	// SELECT * FROM user WHERE id = ?
	b2, getErr2 := engine.ID(id).Get(user)
	if getErr1 != nil {
		fmt.Println(fmt.Sprintf("%+v", getErr2))
	}
	fmt.Println(fmt.Sprintf("%+v", b2))
	fmt.Println(fmt.Sprintf("%+v", user))
}
