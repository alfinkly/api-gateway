-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT UNIQUE NOT NULL,
                       name TEXT NOT NULL,
                       surname TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL
);

INSERT INTO users (name, username, password, email)
VALUES ('Admin', 'Admin', '$2a$12$z/p.8R9OEbZtr61ZYj3eLesD9FmJLTKalX0er1YIyubD.Attn1712', 'admin@example.com');

INSERT INTO users (name, username, password, email)
VALUES ('User', 'User', '$2a$12$z/p.8R9OEbZtr61ZYj3eLesD9FmJLTKalX0er1YIyubD.Attn1712', 'user@example.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd