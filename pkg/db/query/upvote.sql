-- name: CreateUpvote :one
INSERT INTO upvotes (
    "userId", 
    "postId", 
    value
) VALUES (
    $1, $2, $3
) RETURNING "userId", "postId", value, created_at, updated_at;