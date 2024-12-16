-- name: CreateWorkplace :one
INSERT INTO Workplace (ID, Institution, Address, LocalBudgetPercentage)
VALUES ($1, $2, $3, $4)
RETURNING *; 
