# Task API Service

Простой HTTP API для создания, получения и удаления задач.  
Задачи эмулируют длительную I/O-bound операцию (3-5 минут).

---

## Запуск

1. Убедитесь, что установлен Go (версия 1.18 и выше).

2. Склонируйте репозиторий и перейдите в папку проекта:

```bash
git clone https://github.com/Yourazak/Workmate
```
## Установка зависимостей

go mod tidy

## Запускаем сервер

go run main.go

## API
Создает задачу (порт 8080)

POST /tasks
Ответ:
{
"id": "uuid"
}

## Получить задачу по ID

GET /tasks/id
Ответ:
{
"id": "uuid",
"status": "in_progress | done",
"created_at": "timestamp",
"duration": "time elapsed"
}

## Удалить задачу

DELETE /tasks/id
Ответ:
Статус 200 или 204 при успешном удалении.

## ТЕСТИРОВАНИЕ
Рекомендуется использовать curl или Postman для отправки запросов.

