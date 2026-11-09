# 🐹 hello-go — Docker image → GHCR

Учебный проект CI/CD для Go. GitHub Actions проверяет форматирование, выполняет статический анализ и тесты, собирает приложение и публикует Docker-образ в GitHub Container Registry после успешного CI.

## Pipeline

1. `gofmt` — проверка форматирования.
2. `go vet ./...` — статический анализ.
3. `go test ./... -v` — модульные тесты.
4. `go build` — сборка бинарного файла.
5. Docker Buildx собирает multi-stage образ и публикует его в GHCR.

Pull request запускает только CI-проверки. Публикация выполняется после успешного CI при push в `main`.

## Локальный запуск

Требуется Go версии, указанной в `go.mod`, или новее:

```bash
gofmt -w .
go vet ./...
go test ./... -v
go run .
```

## Docker

После первого успешного запуска workflow:

```bash
docker pull ghcr.io/anitouchq/hello-go:latest
docker run --rm ghcr.io/anitouchq/hello-go:latest
```

Пакет можно сделать публичным в GitHub → Packages → Package settings.


## GHCR access troubleshooting

Если публикация завершается ошибкой `permission_denied: write_package`, откройте GitHub Profile → Packages → нужный пакет → **Package settings → Manage Actions access → Add repository**. Добавьте `hello-go` и установите роль **Write**. Затем повторно запустите неуспешный workflow во вкладке Actions.

## Permissions

Для публикации используется встроенный `GITHUB_TOKEN` с разрешением `packages: write`. Внешние секреты не нужны.

## Структура

```text
.
├── .github/workflows/ci.yml
├── greeting/
│   ├── greeting.go
│   └── greeting_test.go
├── Dockerfile
├── go.mod
├── main.go
└── README.md
```

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)
