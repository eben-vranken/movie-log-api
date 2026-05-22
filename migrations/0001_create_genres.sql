-- +goose Up
CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE genres;