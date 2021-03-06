drop database dev_themiseum;
create database dev_themiseum;
use dev_themiseum;

CREATE TABLE files (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  file_hash VARCHAR(64) NOT NULL,
  signature VARCHAR(130) NOT NULL,
  transaction_hash VARCHAR(64) NULL,
  name VARCHAR(255) NULL,
  description VARCHAR(255) NULL,
  block_id BIGINT NULL,
  address VARCHAR(40) NULL,
  confirmed BOOL NULL DEFAULT FALSE,
  uuid VARCHAR(36) NULL,
  email VARCHAR(255) NULL,
  file_index INT NOT NULL,
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE (file_hash, signature),
	INDEX(file_hash, signature)
);

CREATE TABLE notifications (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	email VARCHAR(100) NOT NULL,
	file_id INT NOT NULL,
	was_sent BOOL NULL,
	is_active BOOL NULL,
	create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	FOREIGN KEY(file_id) REFERENCES files(id) ON DELETE CASCADE,
	UNIQUE (file_id, email),
	INDEX(file_id, email)
);

CREATE TABLE labels (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(40) NOT NULL,
	uuid VARCHAR(36) NOT NULL,
	is_active BOOL DEFAULT TRUE,
	create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	UNIQUE (uuid, name),
	INDEX(uuid, name)
);

CREATE TABLE files_labels (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	file_id INT NOT NULL,
	label_id INT NULL,
	create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	FOREIGN KEY(file_id) REFERENCES files(id) ON DELETE CASCADE,
	FOREIGN KEY(label_id) REFERENCES labels(id) ON DELETE CASCADE
);

CREATE TABLE storages (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name ENUM('swarm', 's3') NULL,
  file_path VARCHAR(255) NULL,
  access ENUM('private', 'public') DEFAULT 'public',
  encrypted BOOL NULL,
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE files_storages (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  file_id INT NOT NULL,
  storage_id INT NULL,
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY(file_id) REFERENCES files(id) ON DELETE CASCADE,
  FOREIGN KEY(storage_id) REFERENCES storages(id) ON DELETE CASCADE
);

CREATE VIEW some_files AS SELECT id, file_hash FROM files;