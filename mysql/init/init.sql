-- dbコンテナが作られたタイミングで実行される
CREATE DATABASE IF NOT EXISTS db;
CREATE TABLE IF NOT EXISTS db.todos(
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `text` varchar(255) NOT NULL,
  `done` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO db.todos (text) VALUES ('sample');