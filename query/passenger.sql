-- name: CreateDriver :one
INSERT INTO driver (
  phone,
  hashed_password,
  name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetDriver :one
SELECT * FROM driver
WHERE id = $1 LIMIT 1;

-- name: GetDriverForUpdate :one
SELECT * FROM driver
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetDriverByPhone :one
SELECT * FROM driver
WHERE phone = $1 LIMIT 1;

-- name: ListDrivers :many
SELECT * FROM driver
ORDER BY id
LIMIT $1
OFFSET $2; -- pagination: offset: skip many rows

-- name: UpdateDriver :one
UPDATE driver
SET phone = $2,
  name = $3,
  date_of_birth = $4
WHERE id = $1
RETURNING *;

-- name: UpdatePassword :one
UPDATE driver
SET hashed_password = $2
WHERE id = $1
RETURNING *;

-- name: Verify :one
UPDATE driver
SET verified = true
WHERE phone = $1
RETURNING *;

-- name: DeleteDriver :exec
DELETE FROM driver
WHERE phone = $1;