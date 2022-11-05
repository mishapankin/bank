# [Тестовое задание](https://docs.google.com/document/d/1QE_Z3l8eLe2bIIC9phF-ye2qWZbjM_XjkiFczjTT430/edit)

В качестве драйвера базы данных используется [pgx](http://github.com/jackc/pgx).
Для обработки http-запросов используется [gin](http://github.com/gin-gonic/gin)

Во время разработки бд запускается в контейнере, а http-сервер на хосте. (Потому что нужно часто пересобирать приложение)
Для релиза и бд, и http-сервер запускаются в контейнере.

Запуск обоих серверов в контейнерах (релиз)
```
docker compose -f docker-compose-prod.yml build server
docker compose -f docker-compose-prod.yml up -d
```

Запуск сервера на локальной машине (для разработки)
```bash
make run
```

Запуск контейнера с базой данных (для разработки)
```bash
docker compose up -d
```

Остановка контейнера
```
docker compose down --volumes --remove-orphans
```