-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `phone` VARCHAR(11) NOT NULL COMMENT '手机号（11位数字）',
    `password_hash` VARCHAR(255) NOT NULL COMMENT '密码哈希（bcrypt）',
    `nickname` VARCHAR(50) DEFAULT NULL COMMENT '昵称（可选）',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '账户状态：0=正常，1=锁定，2=禁用',
    `failed_attempts` INT NOT NULL DEFAULT 0 COMMENT '连续登录失败次数',
    `locked_until` DATETIME DEFAULT NULL COMMENT '锁定截止时间',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_phone` (`phone`),
    KEY `idx_status` (`status`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

