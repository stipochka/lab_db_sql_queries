-- name: CreateOperationTypes :one
INSERT INTO OperationTypes (ID, Name, Basepoint, Stock, Cost)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
