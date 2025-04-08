CREATE OR REPLACE TABLE `store_orders` (
    `id` int NOT NULL AUTO_INCREMENT,
    `product_id` int NOT NULL,
    `is_lifetime` boolean NOT NULL,
    `license_id` int NULL,
    `coupon_id` int NULL,
    `final_price` float NULL,
    `status` enum('pending', 'completed', 'cancelled') NOT NULL DEFAULT 'pending',
    `completed_at` timestamp NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `store_products` (`id`),
    FOREIGN KEY (`license_id`) REFERENCES `store_licenses` (`id`),
    FOREIGN KEY (`coupon_id`) REFERENCES `store_coupons` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8;
