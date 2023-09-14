# avito-task

`Роутер - chi`

`СУБД - PostgresSQL`

# Запуск

В файле ./configs/configs.yaml установить нужные значения или оставить их по умолчанию.

```bash
make doker-build
make docker-run

```

# Формат запросов 

Запросы выполнял с помощью Postman
## Создание сегментов 

Пример:

`POST localhost:3000/api/segments`
```json
{
    "segments":[
        "AVITO_VOICE_MESSAGES",
        "AVITO_PERFORMANCE_VAS",
        "AVITO_DISCOUNT_30",
        "AVITO_DISCOUNT_50"
    ]
}
```

## Создание пользователей

Пример:

`POST localhost:3000/api/users`

```json
{
    "users":[
        32,
        64, 
        128, 
        128
    ]
}
```

## Обновление сегментов пользователя 

Пример:

`POST localhost:3000/api/update`

```json
{
    "id": 1001,
    "add":[
        "AVITO_VOICE_MESSAGES",
        "AVITO_PERFORMANCE_VAS"
    ],
    "delete":[]
}
```

## Получение сегментов пользователя

Пример: 

`GET localhost:3000/api/segments?user_id=1001`

## Удаление сегмента 

Пример:

`DELETE localhost:3000/api/segments?name=AVITO_VOICE_MESSAGES`

## Удаление пользователя 

Пример:

`DELETE localhost:3000/api/users?id=1001`