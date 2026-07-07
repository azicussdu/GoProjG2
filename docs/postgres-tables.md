### ТАБЛИЦЫ

```sql
create table courses (
                         id            serial primary key,
                         title         varchar(255) not null,
                         description   text,
                         price         integer not null default 0,
                         level         varchar(50),
                         is_active     boolean not null default false,
                         teacher_id    integer, -- not null references users(id) on delete restrict,
                         created_at    timestamp not null default now(),
                         updated_at    timestamp not null default now(),
                         deleted_at    timestamp null
);
```

```
create table lessons (
                         id          serial primary key,
                         course_id   integer not null references courses(id) on delete cascade,
                         title       varchar(255) not null,
                         content     text,
                         position    integer not null default 1, -- order of lesson in course
                         created_at  timestamp not null default now(),
                         updated_at  timestamp not null default now(),
                         deleted_at  timestamp null
);
```

```
create table users (
                       id              serial primary key,
                       full_name       varchar(255) not null,
                       email           varchar(255) not null unique,
                       password_hash   text not null,
                       role            varchar(50) not null default 'student', -- student | teacher | admin
                       is_active       boolean not null default true,
                       created_at      timestamp not null default now(),
                       updated_at      timestamp not null default now()
);
```

---