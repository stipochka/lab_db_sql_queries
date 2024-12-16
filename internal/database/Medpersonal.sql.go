// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: Medpersonal.sql

package database

import (
	"context"
)

const createMedPersonal = `-- name: CreateMedPersonal :one
INSERT INTO MedPersonal (ID, LastName, Address, TaxPercentage) 
VALUES ($1, $2, $3, $4) 
RETURNING id, lastname, address, taxpercentage
`

type CreateMedPersonalParams struct {
	ID            int32
	Lastname      string
	Address       string
	Taxpercentage string
}

func (q *Queries) CreateMedPersonal(ctx context.Context, arg CreateMedPersonalParams) (Medpersonal, error) {
	row := q.db.QueryRowContext(ctx, createMedPersonal,
		arg.ID,
		arg.Lastname,
		arg.Address,
		arg.Taxpercentage,
	)
	var i Medpersonal
	err := row.Scan(
		&i.ID,
		&i.Lastname,
		&i.Address,
		&i.Taxpercentage,
	)
	return i, err
}