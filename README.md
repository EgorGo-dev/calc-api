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

### Git Bash / Linux / macOS

```bash
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"a":5,"b":3,"op":"add"}'
```

### PowerShell

```powershell
(Invoke-WebRequest -Uri http://localhost:8080/calc -Method Post -Body '{"a":5,"b":3,"op":"add"}' -ContentType "application/json" -UseBasicParsing).Content
```

Ответ:

```json
{"result":8}
```

## Операции

| Операция | Описание |
|----------|----------|
| `add` | Сложение |
| `sub` | Вычитание |
| `mul` | Умножение |
| `div` | Деление |