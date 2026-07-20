### МИГРАЦИИ

Установка инструмента migrate вместе с драйвером postgres:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Проверка установки:
```bash
migrate -version
```

Создание файла миграции:
```bash
migrate create -ext sql -dir migrations -seq create_users
```

Выполняем миграцию с помощью команды:
```bash
migrate -path migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
```
Если подставить наши конфиги:
```bash
migrate -path migrations -database "postgres://admin:admin123@localhost:5432/lmsdb?sslmode=disable" up
```