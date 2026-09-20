# 🧮 Calc API

> REST API калькулятор на Go. Принимает JSON, считает в горутине, логирует запросы через middleware.

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![HTTP](https://img.shields.io/badge/HTTP-net%2Fhttp-005571?style=for-the-badge)](https://pkg.go.dev/net/http)
[![JSON](https://img.shields.io/badge/JSON-encoding%2Fjson-000000?style=for-the-badge&logo=json&logoColor=white)](https://pkg.go.dev/encoding/json)

---

## ✨ Возможности

- 🌐 **HTTP-сервер** на `net/http`
- 📦 **Приём и отдача JSON**
- 🪵 **Middleware** для логирования запросов
- ⚡ **Вычисления в горутине** через каналы и `select`
- 🛡️ **Обработка ошибок** (деление на ноль, неизвестная операция, плохой JSON)

## 🛠 Стек

| Технология | Что используется |
|------------|------------------|
| 🐹 Язык | **Go 1.22+** |
| 📚 Стандартная библиотека | `net/http`, `encoding/json`, `errors`, `log`, `time` |
| ⚙️ Конкурентность | горутины, каналы, `select` |
| 🧩 Архитектура | middleware |

## 🚀 Запуск

```bash
go run .
```

Сервер запустится на **http://localhost:8080**

## 📡 Пример запроса

### 🐧 Git Bash / Linux / macOS

```bash
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"a":5,"b":3,"op":"add"}'
```

### 🪟 PowerShell

```powershell
(Invoke-WebRequest -Uri http://localhost:8080/calc -Method Post -Body '{"a":5,"b":3,"op":"add"}' -ContentType "application/json" -UseBasicParsing).Content
```

### ✅ Ответ

```json
{"result":8}
```

## ➕ Операции

| Операция | Описание | Пример |
|----------|----------|--------|
| `add` | Сложение | `5 + 3 = 8` |
| `sub` | Вычитание | `5 - 3 = 2` |
| `mul` | Умножение | `5 * 3 = 15` |
| `div` | Деление | `6 / 3 = 2` |

## 📂 Исходный код

👉 [**main.go**](./main.go) — весь код проекта в одном файле: обработчики, middleware, горутины и логика калькулятора.

## 📝 Лицензия

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)