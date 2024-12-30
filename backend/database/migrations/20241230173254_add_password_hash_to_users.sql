-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN password_hash text NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN password_hash;
-- +goose StatementEnd
