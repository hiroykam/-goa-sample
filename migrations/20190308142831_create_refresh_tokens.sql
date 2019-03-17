-- +migrate Up

CREATE TABLE IF NOT EXISTS refresh_tokens
(
	jti VARCHAR(36) NOT NULL,
	user_id int unsigned UNIQUE NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
	deleted_at datetime,
	PRIMARY KEY (jti)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down

DROP TABLE IF EXISTS refresh_tokens;
