-- +goose Up

-- +goose StatementBegin

CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        login VARCHAR(255) NOT NULL UNIQUE,
        age INT NOT NULL
    );

-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin

DROP TABLE IF EXISTS users;

-- +goose StatementEnd