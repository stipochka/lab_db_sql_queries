-- +goose Up

CREATE TABLE WorkActivity (
    Contract SERIAL PRIMARY KEY,
    Date VARCHAR(50) NOT NULL,
    MedpersonalID INTEGER NOT NULL REFERENCES Medpersonal(ID) ON DELETE CASCADE,
    WorkplaceID INTEGER NOT NULL REFERENCES Workplace(ID) ON DELETE CASCADE,
    OperationID INTEGER NOT NULL REFERENCES OperationTypes(ID) ON DELETE CASCADE,
    Quantity INTEGER NOT NULL CHECK (Quantity > 0),
    Payment NUMERIC(10, 2) NOT NULL CHECK (Payment >= 0)
);

-- +goose Down

DROP TABLE WorkActivity;
