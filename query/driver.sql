-- name: CreateDriver :one
INSERT INTO driver (
  phone,
  name
) VALUES (
  $1, $2
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

-- name: GetDriverNearby :many
SELECT phone, latitude, longitude, SQRT(
    POW(69.1 * (latitude - sqlc.arg(latitude)::float8), 2) +
    POW(69.1 * (sqlc.arg(longitude)::float8 - longitude) * COS(latitude / 57.3), 2))::float8 AS distance
FROM driver HAVING distance < 25
AND status = 'finding'
ORDER BY distance LIMIT $1;

-- name: ListDrivers :many
SELECT * FROM driver
ORDER BY id
LIMIT $1
OFFSET $2; -- pagination: offset: skip many rows

-- name: UpdateDriver :one
UPDATE driver
SET name = $2,
  date_of_birth = $3
WHERE phone = $1
RETURNING *;

-- name: UpdateLocation :one
UPDATE driver
SET latitude = $2,
  longitude = $3
WHERE phone = $1
RETURNING *;

-- name: UpdateStatus :one
UPDATE driver
SET status = $2
WHERE phone = $1
RETURNING *;

-- name: Verify :one
UPDATE driver
SET verified = true
WHERE phone = $1
RETURNING *;

-- name: DeleteDriver :exec
DELETE FROM driver
WHERE phone = $1;