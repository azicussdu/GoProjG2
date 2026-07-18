### ЭНДПОЙНТЫ ДЛЯ ENROLLMENTS

Все эндпойнты требуют заголовок `Authorization: Bearer <access_token>`.
Enroll/Leave доступны только пользователям с ролью `student`.

**POST http://localhost:8080/api/courses/5/enroll** — Записывает текущего студента на курс с ID 5.

**DELETE http://localhost:8080/api/courses/5/enroll** — Отписывает текущего студента от курса с ID 5.

**GET http://localhost:8080/api/enrollments/me** — Возвращает список курсов, на которые записан текущий студент.

---
