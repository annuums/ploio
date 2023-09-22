/*
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
*/
-- name: GetJob :one
SELECT * FROM jobs
WHERE id = ? LIMIT 1;

-- name: ListCalcJobs :many
SELECT * FROM jobs
WHERE show_calc = 1
ORDER BY project_id, period, name;

-- name: ListJobs :many
SELECT * FROM jobs
ORDER BY project_id, period, name;

-- name: ListJobsByManager :many
SELECT * FROM jobs
WHERE manager_id = ?
ORDER BY project_id, period, name;

-- name: ListJobsCalcsByManager :many
SELECT * FROM jobs
WHERE manager_id = ? AND show_calc = 1
ORDER BY project_id, period, name;

-- name: CreateJob :execresult
INSERT INTO jobs (
  name, period, manager_id, project_id, markdown
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: UpdateJob :execresult
UPDATE jobs SET
  name = ?,
  period = ?,
  manager_id = ?,
  project_id = ?,
  status = ?,
  show_calc = ?,
  markdown = ?
WHERE id = ?;

-- name: UpdateJobStatus :execresult
UPDATE jobs SET
  status = ?
WHERE id = ?;

-- name: DeleteJob :exec
DELETE FROM jobs
WHERE id = ?;