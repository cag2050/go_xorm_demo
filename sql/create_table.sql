CREATE TABLE `exporter_user` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(256) NOT NULL DEFAULT '' COMMENT '姓名',
    `account` varchar(256) NOT NULL DEFAULT '' COMMENT '账号',
    `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
    `mobile` varchar(128) NOT NULL DEFAULT '' COMMENT '手机号',
    `state` int(8) NOT NULL DEFAULT '1' COMMENT '是否启用，0-否;1-是',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `exporter_user_role` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
    `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COMMENT='用户和角色的关系表';

CREATE TABLE `exporter_role` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(256) NOT NULL DEFAULT '' COMMENT '角色名',
    `desc` varchar(1024) NOT NULL DEFAULT '' COMMENT '角色描述',
    `creator` varchar(256) NOT NULL DEFAULT '' COMMENT '创建人',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

CREATE TABLE `exporter_menu` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `menu_pid` int(11) NOT NULL COMMENT '父栏目id',
    `label` varchar(256) NOT NULL DEFAULT '' COMMENT '栏目名字',
    `path` varchar(256) NOT NULL DEFAULT '' COMMENT '栏目路径',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COMMENT='栏目权限表';

CREATE TABLE `exporter_role_menu` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
    `menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '栏目id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `role_menu` (`role_id`,`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=187 DEFAULT CHARSET=utf8mb4 COMMENT='角色和栏目的关系表';