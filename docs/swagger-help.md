### ЭНДПОЙНТЫ ДЛЯ COURSES

Установить CLI для генерации документации:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Далее проверь:
```bash
swag --version
```

Подключить библиотеки в проект
```bash
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

Добавить комментарии Над main():
```go
// @title My API
// @version 1.0
// @description Учебный API
// @host localhost:8080
// @BasePath /
```

Добавить комментарии Над каждым обработчиком:
```go
// @Summary Получить пользователя
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
```

Там где у нас маршруты делаем следующее:
```go
_ "github.com/azicussdu/GoProjG2/docs" // !!! KEREK
swaggerFiles "github.com/swaggo/files"
ginSwagger "github.com/swaggo/gin-swagger"
```

```go
router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

В корне проекта выполнить(пропиши свой путь):
```bash
swag init -g cmd/api/main.go
```

Далее запускаем проект и в браузере набираем:
http://localhost:8080/swagger/index.html

---