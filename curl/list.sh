curl --location --request GET 'http://localhost:8080/v1/list' \
    --header 'Content-Type: application/json' \
    --data '{
        "filter": {
                "title": "t1",
                "description": "d2"
        },
        "sort": {
            "field": "date",
            "order": "asc"
        }
    }'