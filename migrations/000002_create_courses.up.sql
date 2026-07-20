create table courses(
    id          serial
        primary key,
    title       varchar(255)            not null,
    description text,
    price       integer   default 0     not null,
    level       varchar(50),
    is_active   boolean   default false not null,
    teacher_id  integer,
    created_at  timestamp default now() not null,
    updated_at  timestamp default now() not null,
    deleted_at  timestamp
);