package role_test

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/common"
	"github.com/cag2050/go_xorm_demo/model"
	"testing"
)

func TestCount(t *testing.T) {

	name := "role"

	engine, err := common.CreateXORMEngine()
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}

	role := new(model.Role)

	//total, countErr := engine.Where("id >?", 0).Count(role)
	//total, countErr := engine.Where("name = ?", name).Count(role)
	total, countErr := engine.Where("name like ?", "%"+name+"%").Count(role)
	//total, countErr := engine.Count(role)
	if countErr != nil {
		fmt.Println(fmt.Sprintf("%+v", countErr))
	}
	fmt.Println(fmt.Sprintf("%+v", total))
}
