-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE webhooks (
  id INTEGER PRIMARY KEY,
  raw_body BLOB,
  processed_at INTEGER NOT NULL DEFAULT current_timestamp,
  vendor TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE webhooks;
-- +goose StatementEnd
