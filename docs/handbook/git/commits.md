# Git Commits

## Conventional Commits

В проекте используются Conventional Commits.

Формат:

```text
<type>(<scope>): <description>
```

Например:

```text
feat(domain): add Track model
docs: update project roadmap
refactor(provider): split provider interfaces
fix(spotify): handle expired access token
```

---

## Основные типы

### feat

Добавление новой функциональности.

Примеры:

```text
feat(domain): add Playlist model
feat(provider): implement Spotify provider
```

---

### fix

Исправление ошибок.

Примеры:

```text
fix(spotify): handle empty playlists
fix(cli): validate playlist selection
```

---

### refactor

Изменение структуры кода без изменения поведения.

Примеры:

```text
refactor(provider): split interfaces
refactor(transfer): simplify matching flow
```

---

### docs

Изменения только в документации.

Примеры:

```text
docs: add architecture overview
docs: update handbook
```

---

### test

Добавление или изменение тестов.

Примеры:

```text
test(track): add unit tests
```

---

### chore

Технические изменения, не влияющие на функциональность.

Примеры:

```text
chore: update dependencies
chore: add gitignore
```

---

# "Как выбрать тип коммита".

| Если ты...                                   | Используй  |
| -------------------------------------------- | ---------- |
| Добавил новую возможность                    | `feat`     |
| Исправил баг                                 | `fix`      |
| Только переписал код без изменения поведения | `refactor` |
| Менял только документацию                    | `docs`     |
| Добавил тесты                                | `test`     |
| Обновил зависимости, `.gitignore`, CI        | `chore`    |


---

## Scope

Scope показывает, какая часть проекта изменилась.

Примеры:

- domain
- provider
- transfer
- spotify
- ytmusic
- cli
- docs

---

## Перед коммитом

Перед каждым коммитом выполняем:

```bash
go fmt ./...
go build ./...
git status
git diff --staged
```

Только после проверки создаем коммит.

---

## Правила проекта

Один коммит — одна законченная задача.

Не коммитим:

- незаконченный код;
- временные решения;
- закомментированный код;
- бинарные файлы.

Каждый коммит должен собираться и иметь понятное описание.