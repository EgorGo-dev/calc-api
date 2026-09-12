# Calc API

REST API калькулятор на Go. Принимает JSON, считает в горутине, логирует запросы через middleware.

## Возможности

- HTTP-сервер на `net/http`
- Приём и отдача JSON
- Middleware для логирования
- Вычисления в горутине через каналы и `select`
- Обработка ошибок

## Стек

- Go 1.22+
- `net/http`, `encoding/json`
- Горутины, каналы, `select`, middleware

## Запуск

```bash
go run .
```

Сервер запустится на `http://localhost:8080`.

## Пример запроса

```bash
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"a":5,"b":3,"op":"add"}'
```

Ответ:

```json
{"result":8}
```

> 💡 В PowerShell используйте экранирование: `-d "{\"a\":5,\"b\":3,\"op\":\"add\"}"`

## Операции

| Операция | Описание |
|----------|----------|
| `add` | Сложение |
| `sub` | Вычитание |
| `mul` | Умножение |
| `div` | Деление |