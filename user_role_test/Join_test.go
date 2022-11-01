package user_role_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestJoin(t *testing.T) {
	name := "user"

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	type UserRoleJoin struct {
		model.User     `xorm:"extends"`
		model.UserRole `xorm:"extends"`
		model.Role     `xorm:"extends"`
	}

	userRoleJoin := make([]UserRoleJoin, 0)
	// 2个表join
	//if err := engine.Table("exporter_user").Join("left", "exporter_user_role", "exporter_user.id = exporter_user_role.user_id").Find(&userRoleJoin); err != nil {
	// 3个表join
	itemCount, findErr := engine.Table("exporter_user").Join("left", "exporter_user_role", "exporter_user.id = exporter_user_role.user_id").Join("left", "exporter_role", "exporter_user_role.role_id = exporter_role.id").Where("exporter_user.name like ?", "%"+name+"%").FindAndCount(&userRoleJoin)
	if findErr != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}
	fmt.Println(fmt.Sprintf("%+v", itemCount))
	fmt.Println(fmt.Sprintf("%+v", userRoleJoin))

	// 访问UserRoleJoin结构体中的具体对象
	for _, v := range userRoleJoin {
		fmt.Println(fmt.Sprintf("%+v", v.User))
		fmt.Println(fmt.Sprintf("%+v", v.UserRole))
		fmt.Println(fmt.Sprintf("%+v", v.Role))
	}
}
