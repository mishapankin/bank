# [Тестовое задание](https://docs.google.com/document/d/1QE_Z3l8eLe2bIIC9phF-ye2qWZbjM_XjkiFczjTT430/edit)

## Зависимости
В качестве драйвера базы данных используется [pgx](http://github.com/jackc/pgx).
Для обработки http-запросов используется [gin](http://github.com/gin-gonic/gin)

## Запуск
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
```bash
docker compose down --volumes --remove-orphans
```

## Принцип работы
Любые операции, которые клиент производит, сохраняются в базу данных. Когда клиент
хочет получить баланс на счету, баланс пересчитывается так, чтобы учитывать все новые операции.
Каждая операция учитывается единожды, потому что результат операций сохраняется и сохраняется указатель
на последнюю учтенную операцию. Операции учитываются в порядке добавления в базу данных.
Если сервер упадет, то все операции сохранятся.

## Схема базы данных
Схема базы данных копируется внутрь контейнера, поэтому при ее изменени нужно перекомпилировать контейнер с бд.
```bash
docker compose build database --no-cache
```

## Сеть
В релизной версии к базе данных доступ имеет только контейнер с сервером, но снаружи
доступа к бд нет.