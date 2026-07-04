### ЭНДПОЙНТЫ ДЛЯ COURSES

**GET http://localhost:8080/api/courses/5/lessons** — Возвращает все уроки по ID курса

**GET http://localhost:8080/api/lessons/5** — Возвращает один урок по ID.

**POST http://localhost:8080/api/lessons** — Создает новый урок.  
*Тело запроса:*
```json
{
  "title": "Introduction to Go Programming",
  "description": "Learn Go fundamentals including goroutines, channels, and error handling.",
  "price": 5999,
  "level": "advanced",
  "is_active": true,
  "teacher_id": 1
}
```

**PUT http://localhost:8080/api/courses/5** — Обновляет курс по ID.  
*Пример тела запроса:*
```json
{
  "title": "Introduction to Go Programming",
  "price": 5999,
  "level": "advanced"
}
```

**DELETE http://localhost:8080/api/courses/5** — Удаляет курс по ID.

---