# Architecture

## Overview

Spotify2YT построен по принципу разделения ответственности.

Каждый компонент отвечает только за одну задачу.

## High-level Architecture

```text
               CLI
                │
                ▼
        TransferService
           │         │
           ▼         ▼
SpotifyProvider   YTMusicProvider
           │         │
           ▼         ▼
     Spotify API   YTMusic API
```

## Components

### CLI

Точка входа в приложение.

Отвечает только за взаимодействие с пользователем и запуск команд.

---

### TransferService

Координирует процесс переноса.

Он не знает ничего о:

- HTTP
- OAuth
- JSON
- особенностях Spotify API
- особенностях YouTube Music API

Работает исключительно через интерфейс `Provider`.

---

### Provider

Абстракция музыкального сервиса.

Каждая музыкальная платформа реализует собственный Provider.

Например:

- SpotifyProvider
- YTMusicProvider

Provider отвечает за:

- авторизацию;
- работу с API;
- преобразование данных во внутренние модели приложения.

---

### Domain Models

Общие модели приложения.

Они не зависят от внешних API.

Основные модели:

- Track
- Playlist

---

## Design Principles

Во время разработки придерживаемся следующих принципов:

- Single Responsibility Principle (SRP)
- Separation of Concerns
- Composition over Inheritance
- MVP First
- Clean Git History

## Future Improvements

После MVP могут быть добавлены:

- Cache
- Retry
- Concurrent transfer
- Smart Matcher
- Additional Providers (Apple Music, Deezer)