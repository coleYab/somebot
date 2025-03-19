-- name: GetAllFaqs :many
SELECT * FROM "FAQ";

-- name: CreateFAQ :one
INSERT INTO "FAQ" (question, answer) 
VALUES ($1, $2)
RETURNING *;

-- name: UpdateFAQ :exec
UPDATE "FAQ"
SET
    question = $2,
    answer = $2
WHERE id = $1
RETURNING *;