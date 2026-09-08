-- name: GetCardOperationByID :one
SELECT *
FROM card_operations
WHERE id = $1 LIMIT 1;

-- name: ListCardOperations :many
SELECT *
FROM card_operations
WHERE card_id = $1
  AND transaction_date_time >= $2
  AND transaction_date_time <= $3
ORDER BY transaction_date_time LIMIT $4;

-- name: ListCardOperationsByChainID :many
SELECT *
FROM card_operations
WHERE chain_id = $1;

-- name: CreateCardOperation :one
INSERT INTO card_operations (event_id, card_id, chain_id, amount, currency, status, transaction_date_time)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: CreateCardOperationsBatch :many
INSERT INTO card_operations (event_id, card_id, chain_id, amount, currency, status, transaction_date_time)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING (id);

-- name: CreateInbox :one
INSERT INTO inbox (event_id, card_id, chain_id, amount, currency, status, transaction_date_time, inbox_status)
VALUES ($1, $2, $3, $4, $5, $6, $7, 'NEW') RETURNING *;

-- name: MarkInboxDone :exec
UPDATE inbox
SET done_at      = now(),
    inbox_status = 'DONE'
WHERE id = $1;

-- name: MarkInboxError :exec
UPDATE inbox
SET inbox_status = 'ERROR'
WHERE id = $1;

-- name: ListInboxByStatusNew :many
SELECT *
FROM inbox
WHERE inbox_status = 'NEW';

-- name: ListInboxByStatusError :many
SELECT *
FROM inbox
WHERE inbox_status = 'ERROR';

-- name: ListInboxByStatusDone :many
SELECT *
FROM inbox
WHERE inbox_status = 'DONE';
