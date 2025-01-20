CREATE TABLE IF NOT EXISTS `store_orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `license_id` int NOT NULL,
  `cupon_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `store_products` (`id`),
  FOREIGN KEY (`license_id`) REFERENCES `store_licenses` (`id`),
  FOREIGN KEY (`cupon_id`) REFERENCES `store_cupones` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
