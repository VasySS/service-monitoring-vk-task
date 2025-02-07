-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS container_status (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    container_id TEXT NOT NULL,
    ip TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS container_status;
-- +goose StatementEnd
