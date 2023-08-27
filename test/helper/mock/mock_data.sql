INSERT INTO `bank` (`id`, `rand_id`, `name`)
VALUES
    (1, 'test_bank', '三菱UFJ銀行')
ON DUPLICATE KEY UPDATE
     rand_id = VALUES(`rand_id`),
     name = VALUES(`name`);

INSERT INTO `bank_branch` (`id`, `bank_id`, `name`)
VALUES
    (1, 1, 'アンドロメダ支店')
ON DUPLICATE KEY UPDATE
     bank_id = VALUES(`bank_id`),
     name = VALUES(`name`);

INSERT INTO `bank_account` (`id`, `branch_id`, `holder_name`, `number`, `created_at`, `updated_at`)
VALUES
    (1, 1, 'カブシキガイシャタナカ', '111111111', '2023-01-01 00:00:00', '2023-01-01 00:00:00')
ON DUPLICATE KEY UPDATE
     branch_id = VALUES(`branch_id`),
     holder_name = VALUES(`holder_name`),
     number = VALUES(`number`),
     created_at = VALUES(`created_at`),
     updated_at = VALUES(`updated_at`);

INSERT INTO `company` (`id`, `rand_id`, `name`, `representative_name`, `phone_number`, `postal_code`, `address`, `created_at`, `updated_at`)
VALUES
    (1, 'test_company', '株式会社太郎', '田中太郎', '012012340456', '1111111', '東京都', '2023-01-01 00:00:00', '2023-01-01 00:00:00')
ON DUPLICATE KEY UPDATE
     rand_id = VALUES(`rand_id`),
     name = VALUES(`name`),
     representative_name = VALUES(`representative_name`),
     phone_number = VALUES(`phone_number`),
     postal_code = VALUES(`postal_code`),
     address = VALUES(`address`),
     created_at = VALUES(`created_at`),
     updated_at = VALUES(`updated_at`);

INSERT INTO `user` (`id`, `company_id`, `rand_id`, `name`, `mail`, `password_hash`, `password_salt`, `created_at`, `updated_at`)
VALUES
    (1, 1, 'test_user', '田中太郎', 'taro@taro.com', 'xxxxxxxx', 'xxxxxxx', '2023-01-01 00:00:00', '2023-01-01 00:00:00')
ON DUPLICATE KEY UPDATE
     company_id = VALUES(`company_id`),
     rand_id = VALUES(`rand_id`),
     name = VALUES(`name`),
     mail = VALUES(`mail`),
     password_hash = VALUES(`password_hash`),
     password_salt = VALUES(`password_salt`),
     created_at = VALUES(`created_at`),
     updated_at = VALUES(`updated_at`);

INSERT INTO `company_client` (`id`, `rand_id`, `company_id`, `name`, `bank_account_id`, `representative_name`, `phone_number`, `postal_code`, `address`, `created_at`, `updated_at`)
VALUES
    (1, 'test_client', 1, '取引先A', 1, 'テスト', '123001010101', '1111111', '三鷹市', '2023-01-01 00:00:00', '2023-01-01 00:00:00')
ON DUPLICATE KEY UPDATE
     rand_id = VALUES(`rand_id`),
     company_id = VALUES(`company_id`),
     name = VALUES(`name`),
     bank_account_id = VALUES(`bank_account_id`),
     representative_name = VALUES(`representative_name`),
     phone_number = VALUES(`phone_number`),
     postal_code = VALUES(`postal_code`),
     address = VALUES(`address`),
     created_at = VALUES(`created_at`),
     updated_at = VALUES(`updated_at`);

INSERT INTO `invoice` (`id`, `rand_id`, `company_id`, `company_client_id`, `status`, `issue_at`, `payment_amount`, `billing_amount`, `fee`, `fee_ratio`, `tax`, `tax_ratio`, `due_at`, `created_at`, `updated_at`)
VALUES
	(1, 'inv-cjlk6r5315okaj305qd0', 1, 1, 'paid', '2022-10-01 15:00:00', 10000, 10440, 400, 40, 40, 100, '2022-10-01 15:00:00', '2022-10-01 15:00:00', '2022-10-01 15:00:00')
ON DUPLICATE KEY UPDATE
     rand_id = VALUES(`rand_id`),
     company_id = VALUES(`company_id`),
     company_client_id = VALUES(`company_client_id`),
     status = VALUES(`status`),
     issue_at = VALUES(`issue_at`),
     payment_amount = VALUES(`payment_amount`),
     billing_amount = VALUES(`billing_amount`),
     fee = VALUES(`fee`),
     fee_ratio = VALUES(`fee_ratio`),
     tax = VALUES(`tax`),
     tax_ratio = VALUES(`tax_ratio`),
     due_at = VALUES(`due_at`),
     created_at = VALUES(`created_at`),
     updated_at = VALUES(`updated_at`);