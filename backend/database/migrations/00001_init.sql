-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id uuid CONSTRAINT PK_user PRIMARY KEY NOT NULL,
    email text CONSTRAINT UQ_email UNIQUE,
    username text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;
-- +goose StatementBegin
-- +goose StatementEnd
