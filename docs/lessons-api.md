### ЭНДПОЙНТЫ ДЛЯ COURSES

**GET http://localhost:8080/api/courses/5/lessons** — Возвращает все уроки по ID курса

**GET http://localhost:8080/api/lessons/5** — Возвращает один урок по ID.

**POST http://localhost:8080/api/lessons** — Создает новый урок.  
*Тело запроса:*
```json
{
  "course_id": 5,
  "title": "Getting Started",
  "content": "Learn how to set up your environment and prepare for the course.",
  "position": 1
}
```

**PUT http://localhost:8080/api/lessons/5** — Обновляет урок по ID.  
*Пример тела запроса:*
```json
{
  "course_id": 5,
  "content": "Learn how to set up your environment and prepare for the course."
}
```

**DELETE http://localhost:8080/api/lessons/5** — Удаляет урок по ID.

---