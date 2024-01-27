curl --location 'http://localhost:8080/v1/create' \
    --header 'Content-Type: application/json' \
    --data '{
        "title": "t1",
        "description": "t1",
        "image": "ZG8gdGVzdCBodWdlbWFu",
        "date": "2024-01-08T00:00:00+07:00",
        "status": "IN_PROGRESS"
    }'