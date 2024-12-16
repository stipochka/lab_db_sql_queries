-- name: CreateMedPersonal :one
INSERT INTO MedPersonal (ID, LastName, Address, TaxPercentage) 
VALUES ($1, $2, $3, $4) 
RETURNING *;