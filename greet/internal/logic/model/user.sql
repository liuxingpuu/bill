CREATE TABLE `user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `mobile` char(11)  NULL DEFAULT '' COMMENT '用户电话',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `email` varchar(255) NULL DEFAULT '' COMMENT '用户邮箱',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `nickname` varchar(16) NOT NULL COMMENT '用户昵称',
    `avatar` varchar(256) NOT NULL COMMENT '用户头像',
    `last_active_time` timestamp NULL COMMENT '上一次活跃时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `mobile_unique` (`mobile`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';