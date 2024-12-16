-- name: CreateWorkActivity :one
INSERT INTO WorkActivity (Contract, Date, MedpersonalID, WorkplaceID, OperationID, Quantity, Payment)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

