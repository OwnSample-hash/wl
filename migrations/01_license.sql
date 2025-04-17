CREATE OR REPLACE TABLE `store_licenses` (
    `id` int NOT NULL AUTO_INCREMENT,
    `uuid` UUID NOT NULL,
    `user_id` BIGINT NOT NULL REFERENCES `store_users` (`secondary_id`),
    `valid_till` timestamp NOT NULL,
    `grace_period` int NOT NULL DEFAULT 14,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
  );
