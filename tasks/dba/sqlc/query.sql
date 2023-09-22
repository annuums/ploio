/*
  id          BIGINT            NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name        VARCHAR(256)      NOT NULL,
  period      VARCHAR(32)       NOT NULL,
  manager_id  BIGINT            NOT NULL DEFAULT -1,
  job_id      BIGINT            NOT NULL DEFAULT -1,
  markdown    LONGTEXT          NOT NULL DEFAULT "# EMPTY MARKDOWN",
  created_at  DATETIME          NOT NULL DEFAULT NOW(),
  updated_at  DATETIME          NOT NULL DEFAULT NOW(),
  CONSTRAINT `f_manager_id_task`
    FOREIGN KEY (`manager_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `f_job_id_task`
    FOREIGN KEY (`job_id`)
    REFERENCES `jobs` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
*/
-- name: GetTask :one
SELECT * FROM tasks
WHERE id = ? LIMIT 1;

-- name: ListCalcTasks :many
SELECT * FROM tasks
WHERE show_calc = 1
ORDER BY job_id, name;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY job_id, period, name;

-- name: ListTasksByJob :many
SELECT * FROM tasks
WHERE job_id = ?
ORDER BY period, name;

-- name: ListCalcTasksByJob :many
SELECT * FROM tasks
WHERE job_id = ? AND show_calc = 1
ORDER BY name;

-- name: CreateTask :execresult
INSERT INTO tasks (
  name, period, manager_id, job_id, markdown
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: UpdateTask :execresult
UPDATE tasks SET
  name = ?,
  period = ?,
  manager_id = ?,
  job_id = ?,
  markdown = ?
WHERE id = ?;

-- name: UpdateTaskStatus :execresult
UPDATE tasks SET
  status = ?
WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;