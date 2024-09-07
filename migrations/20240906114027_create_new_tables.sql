-- written by Jppnp
-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE table project_manage.positions (
    id bigserial primary key,
    name varchar(255) not null,
    created_at timestamptz default now()
);
CREATE table project_manage.users(
    id bigserial primary key,
    name varchar(255) not null,
    position_id int not null,
    created_at timestamptz default now(),
    constraint fk_position
        foreign key (position_id)
            references project_manage.positions(id)
);

CREATE table project_manage.statuses (
    id serial primary key,
    name varchar(255) not null,
    created_at timestamptz default now()
);

create table project_manage.stories (
    id bigserial primary key,
    title varchar(255) not null,
    detail varchar(255) not null,
    status_id int not null,
    created_at timestamptz default now(),
    created_by bigint not null,
    constraint fk_status
        foreign key (status_id)
            references project_manage.statuses(id)
);

CREATE table project_manage.subtasks (
    id bigserial primary key,
    title varchar(255) not null,
    detail varchar(255) not null,
    status_id int not null,
    assignee bigint,
    story_id bigint not null,
    created_at timestamptz default now(),
    created_by bigint not null,
    constraint fk_status
        foreign key (status_id)
            references project_manage.statuses(id),
    constraint fk_assignee
        foreign key (assignee)
            references project_manage.users(id),
    constraint fk_story
        foreign key (story_id)
            references project_manage.stories(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE project_manage.subtasks CASCADE;
DROP TABLE project_manage.stories CASCADE;
DROP TABLE project_manage.statuses CASCADE;
DROP TABLE project_manage.users CASCADE;
DROP TABLE project_manage.positions CASCADE;
-- +goose StatementEnd
