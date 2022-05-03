-- +goose Up
-- +goose StatementBegin
CREATE TABLE files
(
    id         integer  not null primary key,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp,
    deleted_at datetime,
    name       text     not null,
    expires_at datetime
);
CREATE INDEX idx_files_deleted_at ON files (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_files_deleted_at;
DROP TABLE files;
-- +goose StatementEnd
