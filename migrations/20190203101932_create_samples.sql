-- +migrate Up

CREATE TABLE IF NOT EXISTS samples
(
	id int unsigned NOT NULL AUTO_INCREMENT,
	user_id int unsigned NOT NULL,
	name varchar(255) NOT NULL,
	detail text NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
	deleted_at datetime,
	PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down

DROP TABLE IF EXISTS samples;
