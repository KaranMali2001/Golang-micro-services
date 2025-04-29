-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY created_at
LIMIT $1 OFFSET $2;

-- name: CreateAccount :exec
INSERT INTO accounts (
  id, name
) VALUES (
  $1, $2
);

-- name: UpdateAccount :exec
UPDATE accounts
SET name = $2,
   
    
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;