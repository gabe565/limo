-- +goose Up
-- +goose StatementBegin
CREATE TABLE files(
                      id int not null,
                      name varchar not null,
                      expires text,
                      created_at timestamp not null default current_timestamp,
                      updated_at timestamp not null default current_timestamp,
                      deleted_at timestamp,
                      PRIMARY KEY (id)
);
CREATE INDEX idx_files_deleted_at ON files(deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_files_deleted_at;
DROP TABLE files;
-- +goose StatementEnd
