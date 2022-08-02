-- name: CreateAddress :one
INSERT INTO address (
  driver_id,
  detail,
  house_number,
  street,
  ward,
  district,
  city,
  latitude,
  longitude
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetAddress :one
SELECT * FROM address
WHERE id = $1 LIMIT 1;

-- name: GetAddressByDriverID :one
SELECT * FROM address
WHERE driver_id = $1 LIMIT 1;

-- name: GetAddressForUpdate :one
SELECT * FROM address
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAddresses :many
SELECT * FROM address
ORDER BY id
LIMIT $1
OFFSET $2; 

-- name: UpdateAddress :one
UPDATE address
SET detail = $2,
  house_number = $3,
  street = $4,
  ward = $5,
  district = $6,
  city = $7,
  latitude = $8,
  longitude = $9
WHERE id = $1
RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM address
WHERE id = $1;