CREATE TABLE url_mapping (
  id bigint NOT NULL PRIMARY KEY,
  long_url varchar(255) NOT NULL,
  short_hash varchar(30) NOT NULL UNIQUE,
  INDEX (long_url)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
