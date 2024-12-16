-- +goose Up

CREATE TABLE OperationTypes (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    BasePoint VARCHAR(100) NOT NULL,
    Stock INTEGER NOT NULL CHECK (Stock >= 0),
    Cost NUMERIC(10, 2) NOT NULL CHECK (Cost >= 0)
);

-- +goose Down

DROP TABLE OperationTypes;