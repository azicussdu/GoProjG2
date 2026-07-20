create table lessons
(
    id         serial
        primary key,
    course_id  integer                 not null
        references courses
            on delete cascade,
    title      varchar(255)            not null,
    content    text,
    position   integer   default 1     not null,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp
);