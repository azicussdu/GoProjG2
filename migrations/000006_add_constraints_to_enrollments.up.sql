alter table enrollments
    add constraint fk_enrollments_users
        foreign key (student_id) references users(id)
            on delete cascade;

alter table enrollments
    add constraint fk_enrollments_courses
        foreign key (course_id) references courses(id)
            on delete cascade;

alter table enrollments
    add constraint unique_user_course
        unique (student_id, course_id);