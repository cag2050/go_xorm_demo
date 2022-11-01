package user_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	"testing"
)

func TestDelete(t *testing.T) {
	// 测试的id值
	id := 6

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

	// UPDATE user SET ..., deleted_at = ? WHERE id = ?
	affected, deleteErr := engine.ID(id).Delete(user)
	if deleteErr != nil {
		fmt.Println(fmt.Sprintf("%+v", deleteErr))
	}
	fmt.Println(fmt.Sprintf("%+v", affected))

	// SELECT * FROM user WHERE id = ?
	b2, getErr2 := engine.ID(id).Get(user)
	if getErr2 != nil {
		fmt.Println(fmt.Sprintf("%+v", getErr2))
	}
	fmt.Println(fmt.Sprintf("%+v", b2))
	fmt.Println(fmt.Sprintf("%+v", user))
}
