-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS project_manage.healthcheck (
  id INT NOT NULL primary key,
  name VARCHAR(255)
);
INSERT INTO project_manage.healthcheck VALUES(1, 'I''m fine, ok :)');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS project_manage.healthcheck;
-- +goose StatementEnd
