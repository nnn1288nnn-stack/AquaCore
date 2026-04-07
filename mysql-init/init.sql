-- 澎湖數位老船長 - 數據庫初始化腳本
-- MySQL/MariaDB Initialization Script

-- 選擇資料庫
USE `aquaculture_db`;

-- ========================================
-- 1. 用戶表 (Users)
-- ========================================
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(100) NOT NULL COMMENT '姓名',
  `phone` VARCHAR(20) COMMENT '電話號碼',
  `email` VARCHAR(100) COMMENT '電子郵件',
  `preferred_language` VARCHAR(10) DEFAULT 'zh-TW' COMMENT '偏好語言 (zh-TW/en)',
  `role` ENUM('admin', 'operator', 'viewer') DEFAULT 'operator' COMMENT '角色',
  `status` ENUM('active', 'inactive') DEFAULT 'active' COMMENT '狀態',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
  INDEX `idx_phone` (`phone`),
  INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用戶表';

-- ========================================
-- 2. 環境數據表 (Environmental_Data)
-- ========================================
DROP TABLE IF EXISTS `environmental_data`;
CREATE TABLE `environmental_data` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `water_temperature` DECIMAL(5,2) COMMENT '水溫 (°C)',
  `salinity` DECIMAL(5,2) COMMENT '鹽度 (ppt)',
  `dissolved_oxygen` DECIMAL(5,2) COMMENT '溶氧量 (mg/L)',
  `ph_level` DECIMAL(4,2) COMMENT 'pH 值',
  `ammonia` DECIMAL(5,2) COMMENT '氨值 (mg/L)',
  `recorded_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '記錄時間',
  `notes` TEXT COMMENT '備註',
  INDEX `idx_recorded_at` (`recorded_at`),
  INDEX `idx_date` (`recorded_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='環境數據表';

-- ========================================
-- 3. 資產/庫存表 (Assets)
-- ========================================
DROP TABLE IF EXISTS `assets`;
CREATE TABLE `assets` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(100) NOT NULL COMMENT '物品名稱',
  `category` VARCHAR(50) COMMENT '分類 (飼料/網具/藥劑/工具)',
  `quantity` INT DEFAULT 0 COMMENT '數量',
  `unit` VARCHAR(20) COMMENT '單位 (kg/個/盒/公斤)',
  `reorder_level` INT DEFAULT 10 COMMENT '再訂購點',
  `unit_cost` DECIMAL(10,2) COMMENT '單位成本',
  `supplier` VARCHAR(100) COMMENT '供應商',
  `last_updated_by` INT COMMENT '最後更新者',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
  FOREIGN KEY (`last_updated_by`) REFERENCES `users`(`id`) ON DELETE SET NULL,
  INDEX `idx_category` (`category`),
  INDEX `idx_quantity` (`quantity`),
  INDEX `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='資產/庫存表';

-- ========================================
-- 4. 任務表 (Tasks)
-- ========================================
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `title` VARCHAR(200) NOT NULL COMMENT '任務名稱',
  `description` TEXT COMMENT '任務描述',
  `assigned_to` INT COMMENT '負責人',
  `status` ENUM('pending', 'in-progress', 'completed', 'cancelled') DEFAULT 'pending' COMMENT '狀態',
  `priority` ENUM('low', 'medium', 'high', 'urgent') DEFAULT 'medium' COMMENT '優先級',
  `due_date` DATE COMMENT '截止日期',
  `completed_at` TIMESTAMP NULL COMMENT '完成時間',
  `created_by` INT COMMENT '建立者',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
  FOREIGN KEY (`assigned_to`) REFERENCES `users`(`id`) ON DELETE SET NULL,
  FOREIGN KEY (`created_by`) REFERENCES `users`(`id`) ON DELETE SET NULL,
  INDEX `idx_status` (`status`),
  INDEX `idx_assigned_to` (`assigned_to`),
  INDEX `idx_due_date` (`due_date`),
  INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='任務表';

-- ========================================
-- 5. 操作日誌表 (Operation_Logs)
-- ========================================
DROP TABLE IF EXISTS `operation_logs`;
CREATE TABLE `operation_logs` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT COMMENT '操作用戶',
  `operation_type` VARCHAR(50) COMMENT '操作類型 (create/update/delete)',
  `resource_type` VARCHAR(50) COMMENT '資源類型 (users/assets/tasks)',
  `resource_id` INT COMMENT '資源 ID',
  `details` JSON COMMENT '操作詳情',
  `ip_address` VARCHAR(45) COMMENT 'IP 地址',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '操作時間',
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE SET NULL,
  INDEX `idx_user_id` (`user_id`),
  INDEX `idx_created_at` (`created_at`),
  INDEX `idx_resource` (`resource_type`, `resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日誌表';

-- ========================================
-- 6. 示例數據插入
-- ========================================

-- 插入示例用戶
INSERT INTO `users` (`name`, `phone`, `email`, `role`) VALUES
('王老漁夫', '0912345678', 'wang@aquaculture.tw', 'admin'),
('李小姐', '0987654321', 'lee@aquaculture.tw', 'operator'),
('陳先生', '0933333333', 'chen@aquaculture.tw', 'operator');

-- 插入示例環境數據
INSERT INTO `environmental_data` (`water_temperature`, `salinity`, `dissolved_oxygen`, `ph_level`, `ammonia`) VALUES
(24.5, 30.2, 7.5, 7.8, 0.02),
(24.3, 30.1, 7.4, 7.7, 0.03),
(24.8, 30.3, 7.6, 7.9, 0.01);

-- 插入示例庫存
INSERT INTO `assets` (`name`, `category`, `quantity`, `unit`, `reorder_level`, `unit_cost`) VALUES
('高級飼料 A', '飼料', 150, '公斤', 50, 45.00),
('高級飼料 B', '飼料', 80, '公斤', 50, 52.00),
('大網具', '網具', 5, '個', 2, 2500.00),
('小網具', '網具', 15, '個', 5, 800.00),
('消毒藥劑', '藥劑', 20, '瓶', 10, 250.00);

-- 插入示例任務
INSERT INTO `tasks` (`title`, `description`, `assigned_to`, `status`, `priority`, `due_date`, `created_by`) VALUES
('檢查水質', '每日水質參數檢測', 2, 'completed', 'high', CURDATE(), 1),
('補充飼料', '補充高級飼料 A', 3, 'pending', 'medium', DATE_ADD(CURDATE(), INTERVAL 1 DAY), 1),
('網具維護', '檢查並修復受損網具', 2, 'in-progress', 'medium', DATE_ADD(CURDATE(), INTERVAL 3 DAY), 1);

-- ========================================
-- 視圖定義
-- ========================================

-- 庫存預警視圖
DROP VIEW IF EXISTS `v_low_stock_alert`;
CREATE VIEW `v_low_stock_alert` AS
SELECT 
  `id`,
  `name`,
  `category`,
  `quantity`,
  `reorder_level`,
  `unit`,
  CASE 
    WHEN `quantity` <= `reorder_level` THEN '立即補充'
    WHEN `quantity` <= `reorder_level` * 1.5 THEN '即將不足'
    ELSE '充足'
  END AS stock_status
FROM `assets`
WHERE `quantity` <= `reorder_level` * 1.5
ORDER BY `quantity` ASC;

-- 待辦任務視圖
DROP VIEW IF EXISTS `v_pending_tasks`;
CREATE VIEW `v_pending_tasks` AS
SELECT 
  t.`id`,
  t.`title`,
  u.`name` AS assigned_to_name,
  t.`priority`,
  t.`due_date`,
  DATEDIFF(t.`due_date`, CURDATE()) AS days_until_due
FROM `tasks` t
LEFT JOIN `users` u ON t.`assigned_to` = u.`id`
WHERE t.`status` IN ('pending', 'in-progress')
ORDER BY t.`priority` DESC, t.`due_date` ASC;

-- ========================================
-- 所有表創建完成
-- ========================================
