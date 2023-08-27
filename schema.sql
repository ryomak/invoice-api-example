CREATE TABLE `company` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '会社ID',
  `rand_id` varchar(255) NOT NULL COMMENT 'randID',
  `name` varchar(255) NOT NULL COMMENT '会社名',
  `phone_number` varchar(255) NOT NULL COMMENT '電話番号',
  `postal_code` varchar(255) NOT NULL COMMENT '郵便番号',
  `address` varchar(255) NOT NULL COMMENT '住所',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `rand_id` (`rand_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `company_client` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '取引先ID',
   `company_id` bigint unsigned NOT NULL COMMENT '会社ID',
   `client_company_id` bigint unsigned NOT NULL COMMENT '取引先会社ID',
   `created_at` datetime NOT NULL,
   `updated_at` datetime NOT NULL,
   PRIMARY KEY (`id`),
   UNIQUE KEY `company_cliet_uniq` (`company_id`,`client_company_id`),
   CONSTRAINT `fk_company_client_company_id`
      FOREIGN KEY (`company_id`)
          REFERENCES `company` (`id`)
          ON DELETE RESTRICT ON UPDATE RESTRICT,
   CONSTRAINT `fk_company_client_client_company_id`
      FOREIGN KEY (`client_company_id`)
          REFERENCES `company` (`id`)
          ON DELETE RESTRICT ON UPDATE RESTRICT
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `user` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザID',
   `company_id` bigint unsigned NOT NULL COMMENT '会社ID',
   `rand_id` varchar(255) NOT NULL COMMENT 'randID',
   `name` varchar(255) NOT NULL COMMENT '名前',
   `mail` varchar(255) NOT NULL COMMENT 'メールアドレス',
   `password_hash` varchar(255) NOT NULL COMMENT 'パスワード#',
   `salt` varchar(255) NOT NULL COMMENT 'パスワード用のソルト',
   `created_at` datetime NOT NULL,
   `updated_at` datetime NOT NULL,
   PRIMARY KEY (`id`),
   UNIQUE KEY `rand_id` (`rand_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `bank` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '銀行口座ID',
   `rand_id` varchar(255) NOT NULL COMMENT 'randID',
   `bank_name` varchar(255) NOT NULL COMMENT '銀行名',
   PRIMARY KEY (`id`),
   UNIQUE KEY `rand_id` (`rand_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `bank_branch` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '支店ID',
   `bank_id` bigint unsigned NOT NULL COMMENT '銀行ID',
   `name` varchar(255) NOT NULL COMMENT '銀行支店名',
   PRIMARY KEY (`id`),
   CONSTRAINT `fk_bank_branch_bank_id`
      FOREIGN KEY (`bank_id`)
          REFERENCES `bank` (`id`)
          ON DELETE RESTRICT ON UPDATE RESTRICT
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `bank_account` (
   `id` bigint unsigned NOT NULL COMMENT 'id',
   `company_id` bigint unsigned NOT NULL COMMENT '会社ID',
   `bank_id` bigint unsigned NOT NULL COMMENT '銀行ID',
   `branch_id` bigint unsigned NOT NULL COMMENT '支店ID',
   `holder_name` varchar(255) NOT NULL COMMENT '名義',
   `created_at` datetime NOT NULL,
   `updated_at` datetime NOT NULL,
   PRIMARY KEY (`id`),
   CONSTRAINT `fk_bank_account_company_id`
       FOREIGN KEY (`company_id`)
           REFERENCES `company` (`id`)
           ON DELETE RESTRICT ON UPDATE RESTRICT,
   CONSTRAINT `fk_bank_account_bank_id`
      FOREIGN KEY (`bank_id`)
          REFERENCES `bank` (`id`)
          ON DELETE RESTRICT ON UPDATE RESTRICT,
   CONSTRAINT `fk_bank_account_branch_id`
      FOREIGN KEY (`branch_id`)
          REFERENCES `bank_branch` (`id`)
          ON DELETE RESTRICT ON UPDATE RESTRICT
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `invoice` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '銀行口座ID',
   `rand_id` varchar(255) NOT NULL COMMENT 'randID',
   `company_client_id` bigint unsigned NOT NULL COMMENT '取引先ID',
   `status` enum('init','processing','completed','error') NOT NULL COMMENT 'init=未処理,processing=処理中,completed=支払い済み,error=エラー',
   `issue_at` datetime NOT NULL COMMENT '発行日',
   `amount` bigint unsigned NOT NULL COMMENT '金額',
   `fee` int unsigned NOT NULL COMMENT '手数料',
   `fee_ratio` int unsigned NOT NULL COMMENT '手数料率',
   `tax` int unsigned NOT NULL COMMENT '消費税',
   `tax_ratio` int unsigned NOT NULL COMMENT '消費税率',
   `due_at` datetime NOT NULL COMMENT '支払い期日',
   `created_at` datetime NOT NULL,
   `updated_at` datetime NOT NULL,
   PRIMARY KEY (`id`),
   UNIQUE KEY `rand_id` (`rand_id`),
   CONSTRAINT `fk_invoice_company_client_id`
      FOREIGN KEY (`company_client_id`)
          REFERENCES `company_client` (`id`)
          ON DELETE RESTRICT ON UPDATE RESTRICT
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
