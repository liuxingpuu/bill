CREATE TABLE `account` (
                          `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                          `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
                          `type` varchar(255) NOT NUll DEFAULT '' COMMENT '资产类型',
                          `assets_balance` int(11) NOT NULL DEFAULT 0 COMMENT '资产余额',
                          `account_name` varchar(255) NOT NULL DEFAULT '' COMMENT '账户名',
                          `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户资产表';
