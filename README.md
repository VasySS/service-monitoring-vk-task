## Запуск

Переименовать .env.example в .env

Из корня репозитория необходимо вызвать следующие команды для запуска контейнеров:

```
make build
make up
```

После чего сервисы будут доступны по следующим адресам:

```
frontend: http://localhost:8081
backend: http://localhost:8082
```

## Функциональность

Контейнеры:

- pinger (golang)
- backend (golang)
- frontend (svelte)
- postgres
- postgres-migrations (goose)

Сначала запускается база данных postgres, к ней применяются миграции с помощью контейнера postgres-migrations.

Затем запускается backend, который будет сохранять результаты пингов в postgres и отдавать их по пути `/v1/statuses`.

Запускается pinger, получает IP-адреса всех локальных контейнеров каждые 10 секунд и запускает 5 воркеров, которые пингуют полученные контейнеры. Когда все контейнеры будут пропингованы, результат отправится в backend.

Также запускается frontend, который получает данные из backend и показывает их в виде таблицы:

![frontend example](/img/frontend_example.png)
