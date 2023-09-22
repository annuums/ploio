CREATE TABLE jobs (
  id          BIGINT            NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name        VARCHAR(256)      NOT NULL,
  period      VARCHAR(32)       NOT NULL,
  status      INT               NOT NULL DEFAULT 0,
  manager_id  BIGINT            NOT NULL DEFAULT -1,
  project_id  BIGINT            NOT NULL DEFAULT -1,
  markdown    LONGTEXT          NOT NULL DEFAULT "# EMPTY MARKDOWN",
  show_calc   INT(1)            NOT NULL DEFAULT 0,
  created_at  DATETIME          NOT NULL DEFAULT NOW(),
  updated_at  DATETIME          NOT NULL DEFAULT NOW(),
  CONSTRAINT `f_manager_id_job`
    FOREIGN KEY (`manager_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `f_project_id_job`
    FOREIGN KEY (`project_id`)
    REFERENCES `projects` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);