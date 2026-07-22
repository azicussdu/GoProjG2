create table enrollments (
    id              serial primary key,
    student_id      integer not null,
    course_id       integer not null,
    created_at      timestamp not null default now(),
    updated_at      timestamp not null default now()
);