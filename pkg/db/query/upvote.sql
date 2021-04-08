-- name: CreateUpvote :one
INSERT INTO upvotes (
    "userId", 
    "postId", 
    value
) VALUES (
    $1, $2, $3
) RETURNING "userId", "postId", value, created_at, updated_at;

-- name: GetUpvote :one
SELECT "userId", "postId", value, created_at, updated_at FROM upvotes WHERE "userId" = $1 AND "postId" = $2 LIMIT 1;

-- name: UpdateUpvote :exec
UPDATE upvotes SET value = $3 WHERE "userId" = $1 AND "postId" = $2;
-- UPDATE upvotes SET value = 12 WHERE "userId" = 1 AND "postId" = 1 RETURNING *;

-- name: DeleteUpvoteByPost :exec
DELETE FROM upvotes WHERE "postId" = $1;