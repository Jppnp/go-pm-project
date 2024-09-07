-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO project_manage.statuses (name) VALUES ('backlog'), ('inprogress'), ('testing'), ('done');
INSERT INTO project_manage.positions (name) VALUES ('developer'), ('lead developer'), ('quality assurance');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM project_manage.statuses WHERE name in ('backlog', 'inprogress', 'testing', 'done');
DELETE FROM project_manage.positions WHERE name in ('developer', 'lead developer', 'quality assurance');
-- +goose StatementEnd