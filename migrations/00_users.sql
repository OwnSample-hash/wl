CREATE OR REPLACE TABLE `store_users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `secondary_id` long NOT NULL,
    `email` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
  );
