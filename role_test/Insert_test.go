package role_test

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
	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	role := new(model.Role)
	role.Name = "role3"
	rand.Seed(time.Now().UnixNano())
	role.Desc = "desc" + strconv.Itoa(rand.Intn(1000)+1)

	affected, insertErr := engine.InsertOne(role)
	if insertErr != nil {
		fmt.Println(fmt.Sprintf("%+v", insertErr))
	}
	fmt.Println(fmt.Sprintf("%+v", affected))
	fmt.Println(fmt.Sprintf("%+v", role))
}
