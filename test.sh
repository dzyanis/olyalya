curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/create' -d '{"name": "dz"}'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:8080/db'

curl -X POST 'http://localhost:8080/db/dz/set' -d '{
    "one": 1
}'
curl -X POST 'http://localhost:8080/db/dz/set' -d '{
    "two": 2
    "three": 3
}'