CREATE TABLE IF NOT EXISTS `store_licenses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` UUID NOT NULL,
  `name` varchar(255) NOT NULL,
  `valid_till` timestamp NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
