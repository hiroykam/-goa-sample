-- +migrate Up

CREATE TABLE IF NOT EXISTS users
(
	id int unsigned NOT NULL AUTO_INCREMENT,
	email varchar(255) UNIQUE NOT NULL,
  hashed_password varchar(255) NOT NULL,
	revoked_at datetime,
	locked_at  datetime,
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
	deleted_at datetime,
	PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

INSERT INTO users (email, hashed_password) VALUES("sample@goa-sample.test.com", "$2a$10$jOvPvkt8BirlbRfrkGFbae/nSe7U90suYvqB.bvIbg.cu4D3UPx2G");

-- +migrate Down

DROP TABLE IF EXISTS users;
