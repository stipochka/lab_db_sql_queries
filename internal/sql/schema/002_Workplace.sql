-- +goose Up

CREATE TABLE Workplace (
    ID SERIAL PRIMARY KEY,
    Institution VARCHAR(100) NOT NULL,
    Address VARCHAR(100) NOT NULL,
    LocalBudgetPercentage NUMERIC(4, 2) NOT NULL CHECK (LocalBudgetPercentage >= 0)
);

-- +goose Down

DROP TABLE Workplace;