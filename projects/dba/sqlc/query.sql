/*
  id      BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name    varchar(256)        NOT NULL,
  client  varchar(64)         NOT NULL,
  period  VARCHAR(32)         NOT NULL,
  manager_id BIGINT           NOT NULL DEFAULT -1,
  show_calc   INT(1)          NOT NULL DEFAULT 1,
  created_at  DATETIME        NOT NULL DEFAULT NOW(),
  updated_at  DATETIME        NOT NULL DEFAULT NOW(),
  CONSTRAINT `f_manager_id_project`
    FOREIGN KEY (`manager_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
*/
-- name: GetProject :one
SELECT * FROM projects
WHERE id = ? LIMIT 1;

-- name: ListCalcProjects :many
SELECT * FROM projects
WHERE show_calc = 1
ORDER BY client, period, name;

-- name: ListProjects :many
SELECT * FROM projects
ORDER BY client, period, name;

-- name: ListProjectsByManager :many
SELECT * FROM projects
WHERE manager_id = ?
ORDER BY client, period, name;

-- name: ListProjectCalcsByManager :many
SELECT * FROM projects
WHERE manager_id = ? AND show_calc = 1
ORDER BY client, period, name;

-- name: CreateProject :execresult
INSERT INTO projects (
  name, client, period, manager_id, markdown
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: UpdateProject :execresult
UPDATE projects SET
  name = ?,
  client = ?,
  period = ?,
  manager_id = ?,
  markdown = ?
WHERE id = ?;

-- name: UpdateProjectStatus :execresult
UPDATE projects SET
  status = ?
WHERE id = ?;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = ?;