// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: Workplace.sql

package database

import (
	"context"
)

const createWorkplace = `-- name: CreateWorkplace :one
INSERT INTO Workplace (ID, Institution, Address, LocalBudgetPercentage)
VALUES ($1, $2, $3, $4)
RETURNING id, institution, address, localbudgetpercentage
`

type CreateWorkplaceParams struct {
	ID                    int32
	Institution           string
	Address               string
	Localbudgetpercentage string
}

func (q *Queries) CreateWorkplace(ctx context.Context, arg CreateWorkplaceParams) (Workplace, error) {
	row := q.db.QueryRowContext(ctx, createWorkplace,
		arg.ID,
		arg.Institution,
		arg.Address,
		arg.Localbudgetpercentage,
	)
	var i Workplace
	err := row.Scan(
		&i.ID,
		&i.Institution,
		&i.Address,
		&i.Localbudgetpercentage,
	)
	return i, err
}
