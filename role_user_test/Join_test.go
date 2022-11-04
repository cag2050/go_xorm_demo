package role_user_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

// 返回角色表的数据，同时带角色信息（用户和角色是一对一关系）
func TestJoin(t *testing.T) {

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	type RoleCount struct {
		model.Role `xorm:"extends"`
		NumUsers   int64
	}
	var roleCount = make([]*RoleCount, 0)
	if err := engine.SQL("select exporter_role.*, (select count(id) from exporter_user_role where exporter_user_role.role_id = exporter_role.id) as num_users from exporter_role, exporter_user_role where exporter_role.id = exporter_user_role.role_id").Find(&roleCount); err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	for k, v := range roleCount {
		fmt.Println(fmt.Sprintf("%+v", k))
		fmt.Println(fmt.Sprintf("%+v", v))
	}
}
