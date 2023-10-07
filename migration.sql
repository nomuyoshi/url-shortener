CREATE TABLE url_mapping (
  id bigint NOT NULL PRIMARY KEY,
  long_url varchar(255) NOT NULL UNIQUE,
  short_url varchar(30) NOT NULL UNIQUE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
