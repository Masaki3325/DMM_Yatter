CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE,
  `password_hash` varchar(255) NOT NULL,
  `display_name` varchar(255),
  `avatar` text,
  `header` text,
  `note` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `status` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account_id` bigint(20) NOT NULL,
  `content` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`account_id`) REFERENCES `account` (`id`)
)

-- CREATE TABLE `attachment` (
--   `id` bigint(20) NOT NULL AUTO_INCREMENT,
--   `type` enum('image', 'video', 'audio', 'file')
--   `url` varchar(2048)
--   `description` text
--   PRIMARY KEY (`id`)
-- )

-- CREATE TABLE `status_attachment` (
--   `status_id` bigint(20) NOT NULL,
--   `attachment_id` bigint(20) NOT NULL,
--   PRIMARY KEY (`id`),
--   FOREIGN KEY (`status_id`) REFERENCES `status` (`id`),
--   FOREIGN KEY (`attachment_id`) REFERENCES `attachment` (`id`)
-- );