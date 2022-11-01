package model

import "time"

// User SnakeMapper 支持struct为驼峰式命名，表结构为下划线命名之间的转换，这个是默认的Maper；https://xorm.io/zh/docs/chapter-02/1.mapping/
type User struct {
	// 如果field名称为Id而且类型为int64并且没有定义tag，则会被xorm视为主键，并且拥有自增属性。如果想用Id以外的名字或非int64类型做为主键名，必须在对应的Tag上加上xorm:"pk"来定义主键，加上xorm:"autoincr"作为自增。这里需要注意的是，有些数据库并不允许非主键的自增属性。https://xorm.io/zh/docs/chapter-02/4.columns/
	Id        int64
	Name      string
	Account   string
	Password  string
	Mobile    string
	State     int64
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (User) TableName() string {
	return "exporter_user"
}
