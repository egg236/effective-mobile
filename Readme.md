# REST-сервис для управления подписками

REST-сервис для CRUDL-операций с записями о подписках пользователей. Использует `go1.22`.

---

## API

API описан в формате Swagger [здесь](./docs/swagger.yaml)

---

## Используемые зависимости

* **Роутер**: `github.com/gorilla/mux`
* **Драйвер для БД**: `github.com/jackc/pgx/v4/stdlib`
* **Логирование**: `log/slog`
* **Миграции**: `github.com/golang-migrate/migrate` через docker-compose сервис
* **UUID**: `github.com/google/uuid`

---

## Конфигурация

### Запуск
Проект можно запустить через docker-compose. Для этого необходим запущенный docker-daemon. Далее нужно запустить команду:
```bash
docker compose up
```
Она автоматически поднимет instance postgreSQL (данные не вынесены в volume и не сохраняются после перезапуска), применит миграции и запустит приложение.

### Переменные окружения

Все переменные окружения задаются в файле `.env` в корне проекта.

**Пример:**
```
# Настройки базы данных (для POSTGRES сервиса)
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=effective_mobile_db

# Настройки базы данных (для приложения)
DB_HOST: db
DB_PORT: 5432
DB_USER: record_admin
DB_PASS: pass123
DB_NAME: production

# Порт сервера сервера
PORT=8090
```