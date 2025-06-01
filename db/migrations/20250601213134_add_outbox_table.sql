-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS outboxes (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    processed boolean NOT NULL DEFAULT FALSE,
    event_type VARCHAR(255) NOT NULL,
    event_payload JSONB NOT NULL,
    event_metadata 

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
