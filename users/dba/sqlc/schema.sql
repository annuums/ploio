CREATE TABLE users (
  id          BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  auth        int     NOT NULL DEFAULT 0,
  name        varchar(32)    NOT NULL,
  email       varchar(256)   NOT NULL,
  pwd         varchar(256)   NOT NULL,
  departure   varchar(32)    NOT NULL DEFAULT '대기중',
  profile     varchar(512)   NOT NULL DEFAULT '/img/logo.png',
  created_at  DATETIME       NOT NULL DEFAULT NOW(),
  updated_at  DATETIME       NOT NULL DEFAULT NOW()
);