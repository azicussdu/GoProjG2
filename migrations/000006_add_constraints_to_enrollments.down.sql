alter table enrollments
    drop constraint fk_enrollments_users;

alter table enrollments
    drop constraint fk_enrollments_courses;

alter table enrollments
    drop constraint unique_user_course;