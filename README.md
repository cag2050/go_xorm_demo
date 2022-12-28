## XORM：Go 语言 ORM 框架

### XORM 使用注意点：
1. DeletedAt对应的数据库字段，必须允许NULL并且设置默认为NULL
举例，go结构体：
```go
type User struct {
    Id int64
    Name string
    DeletedAt time.Time `xorm:"deleted"`
}
```
对应建表语句：
```sql
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL DEFAULT '' COMMENT '姓名',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
```
2. UNIQUE KEY 类型的mysql字段，重复插入时，报错：Error 1062: Duplicate entry '' for key 'your_key'

资料 | 说明
--- | ---
xorm官网 | https://xorm.io/zh/
结合RBAC模型讲解权限管理系统需求及表结构创建 | https://juejin.cn/post/6844904004984504333
