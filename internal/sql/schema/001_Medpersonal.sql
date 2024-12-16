-- +goose Up

CREATE TABLE Medpersonal (
    ID SERIAL PRIMARY KEY,
    LastName VARCHAR(50) NOT NULL,
    Address VARCHAR(100) NOT NULL,
    TaxPercentage NUMERIC(4, 2) NOT NULL CHECK (TaxPercentage >= 0)
);

-- +goose Down

DROP TABLE Medpersonal;
