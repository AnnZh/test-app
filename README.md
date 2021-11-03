## test-app

### Запросы (пример)
* Добавление данных:
POST http://localhost:8000/api/get-data
Body:
```json
{
    "date": "20.12.2019 14:41:25",
    "number": "1234 PP-7",
    "speed": "41,5"
}
```

* Авто превысившие указанную скорость за указанную дату:
GET http://localhost:8000/api/queries/over-speed
Body:
```json
{
    "date": "20.12.2019",
    "speed": "21,3"
}
```

* Максимальная и минимальная зафиксированная скорость за указанную дату:
GET http://localhost:8000/api/queries/min-max
Body:
```json
{
    "date": "25.12.2019"
}
```

