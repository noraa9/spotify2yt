# Spotify2YT

Spotify2YT — это CLI-приложение на Go для переноса плейлистов из Spotify в YouTube Music.

## Goals

- Изучить Go на практике.
- Спроектировать расширяемую архитектуру.
- Построить проект так, как это делают в реальных командах.
- Создать MVP, который можно постепенно развивать.

## Roadmap

### Milestone 1 - Foundation

- [x] Спроектирована архитектура
- [x] Определены доменные сущности
- [x] Определены основные сервисы
- [ ] Реализован Provider
- [ ] Реализованы модели
- [ ] CLI

### Milestone 2 - Spotify

- [ ] OAuth
- [ ] Получение плейлистов
- [ ] Получение треков

### Milestone 3 - YouTube Music

- [ ] Авторизация
- [ ] Создание плейлиста
- [ ] Поиск треков

### Milestone 4 - Transfer

- [ ] Перенос треков

## Project Structure

```
cmd/            CLI commands
internal/       Application code
docs/           Documentation
configs/        Configuration files
scripts/        Helper scripts
```

## Development

Проект разрабатывается итеративно.

Каждый этап заканчивается рабочей версией приложения.

Архитектурные решения документируются в папке `docs/decisions`.