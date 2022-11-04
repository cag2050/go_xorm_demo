package user_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	name := "name888"
	rand.Seed(time.Now().UnixNano())
	account := "account" + strconv.Itoa(rand.Intn(1000)+1)

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	user := new(model.User)
	user.Name = name
	user.Account = account
	user.State = 1

	affected, insertErr := engine.InsertOne(user)
	if insertErr != nil {
		fmt.Println(fmt.Sprintf("%+v", insertErr))
		return
	}
	fmt.Println(fmt.Sprintf("affected: %+v", affected))
	// 数据插入成功，但是 strconv.ParseInt() 转换插入的id时失败，也会返回 err
	if affected == 1 {
		fmt.Println(fmt.Sprintf("user.Id: %+v", user.Id))
	}
	fmt.Println(fmt.Sprintf("%+v", user))
}
