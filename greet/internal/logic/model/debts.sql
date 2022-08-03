CREATE TABLE `debt` (
                          `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                          `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
                          `type` varchar(255) NOT NUll DEFAULT '' COMMENT '负债账户类型',
                          `debt_amount` int(11) NOT NULL DEFAULT 0 COMMENT '负债金额',
                          `debt_total` int(11) NOT NULL DEFAULT 0 COMMENT '负债总金额',
                          `lines` int(11) NOT NULL DEFAULT 5000 COMMENT '额度',
                          `account_name` varchar(255) NOT NULL DEFAULT '' COMMENT '账户名',
                          `pay_date`  char NULL DEFAULT CURRENT_TIMESTAMP COMMENT '还款日',
                          `bill_date` char NULL DEFAULT CURRENT_TIMESTAMP COMMENT '账单日',
                          `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='负账账户表';
