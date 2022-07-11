CREATE TABLE `people`
(
    `id`                 bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`         datetime(3) DEFAULT NULL,
    `updated_at`         datetime(3) DEFAULT NULL,
    `deleted_at`         datetime(3) DEFAULT NULL,
    `name`               varchar(64) NOT NULL COMMENT '姓名',
    `age`                int unsigned NOT NULL COMMENT '年龄',
    `address_list`          varchar(1024) NOT NULL COMMENT '地址',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = 'people 表';
