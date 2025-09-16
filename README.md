# Wails ToDo App

Простое ToDo-приложение, написанное на **Go + Wails (Go + React)**.  
Поддерживает добавление, удаление, завершение задач, фильтрацию и переключение тёмной/светлой темы.

## Функционал

- Добавление задач
- Пометка задачи как выполненной
- Удаление задачи с подтверждением
- Фильтрация задач:
    - Все
    - Активные
    - Завершённые
- Переключение между светлой и тёмной темой

## Структура проекта

```bash
wails-app
├── frontend/           # Интерфейс (HTML/CSS/JS)
├── main.go             # Бэкенд на Go
├── go.mod
├── go.sum
└── wails.json          # Конфигурация Wails
```

## Запуск
- Настроить строку подключения к pq в файле wails_app/database/database.go
-  Создать таблицу из файла 1.initial_tables.sql
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
git clone https://github.com/beginore/test_to_do_list.git
cd wails_app
wails dev
```
