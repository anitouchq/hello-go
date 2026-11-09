# hello-go

[![CI](https://github.com/anitouchq/hello-go/actions/workflows/ci.yml/badge.svg)](https://github.com/anitouchq/hello-go/actions)

Учебный проект на Go: pipeline проверяет форматирование, выполняет статический анализ и тесты, собирает приложение и публикует Docker-образ в GitHub Container Registry (GHCR).

## Что делает pipeline

1. `gofmt` — проверяет форматирование исходников.
2. `go vet ./...` — выполняет статический анализ.
3. `go test ./... -v` — запускает модульные тесты.
4. `go build` — проверяет сборку приложения.
5. Docker Buildx собирает multi-stage образ и публикует его в GHCR.

Pull request запускает только CI-проверки. Публикация выполняется при push в `main` после успешного job `test`.

## Локальный запуск

Требуется Go версии, указанной в `go.mod`, или новее:

```bash
gofmt -w .
go vet ./...
go test ./... -v
go run .
```

## Docker image

```bash
docker pull ghcr.io/anitouchq/hello-go:latest
docker run --rm ghcr.io/anitouchq/hello-go:latest
```

[Открыть пакет в GHCR](https://github.com/users/anitouchq/packages/container/hello-go)

## Структура проекта

| Файл или каталог | Назначение |
|---|---|
| `.github/workflows/ci.yml` | CI-проверки и публикация Docker-образа |
| `greeting/greeting.go` | Функции приветствия и вспомогательная логика |
| `greeting/greeting_test.go` | Тесты пакета `greeting` |
| `main.go` | Точка входа приложения |
| `Dockerfile` | Multi-stage сборка контейнера |
| `go.mod` | Имя Go-модуля и версия Go |
| `.dockerignore` | Файлы, исключаемые из Docker build context |
| `.gitignore` | Локальные файлы, исключаемые из Git |

## Результаты

- [Последние запуски GitHub Actions](https://github.com/anitouchq/hello-go/actions)
- [Docker package в GHCR](https://github.com/users/anitouchq/packages/container/hello-go)

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)

## Скриншоты

Добавьте изображения в каталог `screenshots/`:

- `screenshots/application.png` — пример запуска приложения;
- `screenshots/github-actions.png` — успешный CI и Docker publish;
- `screenshots/ghcr-package.png` — опубликованный Docker package в GHCR.

![Пример работы приложения](screenshots/application.png)
![GitHub Actions](screenshots/github-actions.png)
![Docker package в GHCR](screenshots/ghcr-package.png)

## Как опубликовать новую версию Docker image

В этом проекте публикация выполняется не через GitHub Releases, а автоматически при push в `main`. После успешных тестов GitHub Actions собирает и отправляет образ в GHCR.

```bash
git add .
git commit -m "feat: describe your change"
git push origin main
```

Образ: `ghcr.io/anitouchq/hello-go:latest`.

- [GitHub Actions](https://github.com/anitouchq/hello-go/actions)
- [Docker package в GHCR](https://github.com/users/anitouchq/packages/container/hello-go)
