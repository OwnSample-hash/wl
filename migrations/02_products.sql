CREATE OR REPLACE TABLE `store_products` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` text NOT NULL,
    `price_per_month` decimal(10, 2) NOT NULL,
    `price` decimal(10, 2) NOT NULL,
    `one_month` boolean NOT NULL,
    `life_time` boolean NOT NULL,
    `discount` decimal(10, 2) NOT NULL,
    `is_active` boolean NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8;
