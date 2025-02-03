-- name: CreateUser :one
INSERT INTO "User" (id, first_name, last_name, phone_number, referal_code, referred_by, language, profile_pic) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "User" WHERE id = $1;

-- name: UpdateUser :one
UPDATE "User" 
SET net_balance = $2, referal_count = $3 
WHERE id = $1 
RETURNING *;

-- Select All Users
SELECT * FROM "User";

-- Select Users by referred_by
SELECT * FROM "User" WHERE referred_by = $1;

-- Select a User by ID
SELECT * FROM "User" WHERE id = $1;
