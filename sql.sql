

-- 用户表
CREATE TABLE `user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `name` varchar(45) NOT NULL COMMENT '用户账号',
    `nick_name` varchar(45) NOT NULL COMMENT '用户昵称',
    `password` varchar(45) NOT NULL COMMENT '用户密码',
    `status`    tinyint(4) NOT NULL   COMMENT '用户状态：1/正常,-1/禁用',
    `create_at` datetime DEFAULT NULL COMMENT '创建时间',
    `update_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;