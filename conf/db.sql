create database if not exists kvm_manager charset='utf8mb4';
CREATE TABLE `t_backup_task` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '备份任务ID',
  `name` varchar(100) NOT NULL COMMENT '虚拟机名称',
  `ip` varchar(50) NOT NULL COMMENT 'IP地址',
  `schedule_type` enum('at','cron') DEFAULT NULL COMMENT '计划任务方式',
  `cron_expression` varchar(100) DEFAULT NULL COMMENT 'CRON表达式',
  `at_time` varchar(19) DEFAULT '' COMMENT 'AT时间,格式YYYY-MM-DD HH:MM:SS',
  `retention_period` int NOT NULL DEFAULT '0' COMMENT '保留周期,单位天,默认0,表示永久保存',
  `status` varchar(20) DEFAULT '0' COMMENT '备份任务状态,0:启用,1:禁用',
  `ctime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `utime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_ip` (`ip`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='备份任务表';

CREATE TABLE `t_backup_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '备份ID',
  `backup_id` int NOT NULL COMMENT '备份任务ID',
  `client_ip` varchar(50) NOT NULL COMMENT '操作来源IP地址',
  `log_table` varchar(50) NOT NULL COMMENT '日志类型',
  `operator` varchar(50) NOT NULL COMMENT '操作人',
  `content` text NOT NULL COMMENT '操作内容',
  `ctime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_backup_id` (`backup_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日志表';