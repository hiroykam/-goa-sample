-- +migrate Up
CREATE TABLE IF NOT EXISTS samples
(
	id int unsigned NOT NULL AUTO_INCREMENT,
	user_id int unsigned NOT NULL,
	name varchar(255) NOT NULL,
	detail text NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT = 'inquiries';

-- +migrate Down

DROP TABLE IF EXISTS samples;
