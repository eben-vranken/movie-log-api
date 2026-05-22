-- +goose Up
CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    genre_id INT NOT NULL,
    director VARCHAR(32) NOT NULL,
    release_date DATE NOT NULL,
    runtime INT NOT NULL,
    rating DECIMAL(3, 1) NOT NULL
);

-- +goose Down
DROP TABLE plants;