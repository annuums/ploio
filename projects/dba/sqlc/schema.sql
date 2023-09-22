CREATE TABLE projects (
  id          BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name        varchar(256)        NOT NULL,
  client      varchar(64)         NOT NULL,
  period      VARCHAR(32)         NOT NULL,
  status      INT                 NOT NULL DEFAULT 0,
  markdown    LONGTEXT        NOT NULL DEFAULT "# EMPTY MARKDOWN",
  manager_id  BIGINT          NOT NULL DEFAULT -1,
  show_calc   INT(1)          NOT NULL DEFAULT 1,
  created_at  DATETIME        NOT NULL DEFAULT NOW(),
  updated_at  DATETIME        NOT NULL DEFAULT NOW(),
  CONSTRAINT `f_manager_id_project`
    FOREIGN KEY (`manager_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);