-- wrtten by jppnp
-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SCHEMA IF NOT EXISTS project_manage;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP SCHEMA IF EXISTS project_manage;
-- +goose StatementEnd
