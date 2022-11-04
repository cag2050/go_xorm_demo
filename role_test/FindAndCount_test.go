package role_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	"testing"
)

func TestFindAndCount(t *testing.T) {
	var id int64 = 14
	name := ""
	var pageSize int64 = 10
	var pageIndex int64 = 1

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	// 1. 获得角色列表
	roles := make([]model.Role, 0)

	session := engine.Where("1=1")
	if id > 0 {
		session = session.ID(id)
	}
	if len(name) > 0 {
		session = session.Where("name like ?", "%"+name+"%")
	}
	// 返回总记录条数的同时，进行分页
	count, err1 := session.Limit(int(pageSize), int((pageIndex-1)*pageSize)).FindAndCount(&roles)
	if err1 != nil {
		fmt.Println(fmt.Sprintf("%+v", err1))
	}
	fmt.Println(fmt.Sprintf("%+v", count))
	//fmt.Println(fmt.Sprintf("%+v", roles))

	// 返回的数据
	type OneRole struct {
		Role  model.Role
		Users []model.User
	}

	allRole := []OneRole{}

	// 2. 根据每个角色id查询user_id
	for _, v := range roles {
		var oneRole OneRole
		oneRole.Role = v

		fmt.Println(fmt.Sprintf("%+v", v))
		usersInRole := make([]model.UserRole, 0)
		session := engine.Where("1=1")
		//if id > 0 {
		session.Where("role_id = ?", v.Id)
		//}
		userCount, err := session.FindAndCount(&usersInRole)
		if err != nil {
			fmt.Println(fmt.Sprintf("%+v", err))
			return
		}
		fmt.Println(fmt.Sprintf("%+v", userCount))

		// 3. 根据user_id查询user
		for _, v := range usersInRole {
			fmt.Println(fmt.Sprintf("%+v", v))
			user := new(model.User)
			has, err := engine.ID(v.UserId).Get(user)
			if err != nil {
				fmt.Println(fmt.Sprintf("%+v", err))
				return
			}
			if has {
				oneRole.Users = append(oneRole.Users, *user)
			}
		}
		allRole = append(allRole, oneRole)
	}
	fmt.Println(fmt.Sprintf("%+v", allRole))
	fmt.Println(fmt.Sprintf("%+v", len(allRole)))
}
