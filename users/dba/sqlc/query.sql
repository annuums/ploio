/*
  id          BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  auth        int     NOT NULL DEFAULT 0,
  name        varchar(32)    NOT NULL,
  email       varchar(256)   NOT NULL,
  pwd         varchar(256)   NOT NULL,
  departure   varchar(32)    NOT NULL DEFAULT '대기중',
  profile     varchar(512)   NOT NULL DEFAULT '/img/logo.png',
  created_at  DATETIME       NOT NULL DEFAULT NOW(),
  updated_at  DATETIME       NOT NULL DEFAULT NOW()
*/

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT
  id, auth, name, email, departure, profile, created_at, updated_at
FROM users
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (
  auth, name, email, pwd, departure, profile
) VALUES (
  ?, ?, ?,
  SHA2(pwd, 256),
  ?,
  '/img/logo.png'
);

-- name: UpdateUser :execresult
UPDATE users SET
  auth=?,
  pwd=?,
  departure=?,
  profile=?
WHERE
  id = ?;
  
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;