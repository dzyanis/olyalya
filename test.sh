curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/create' -d '{"name": "dz"}'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:8080/db'

curl -X GET 'http://localhost:8080/db/dz'

curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/dz/set' -d '{
    "one": "1"
}'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/dz/set' -d '{
    "two": "2"
    "three": "3"
}'

curl -X GET 'http://localhost:8080/db/dz/get/one'
curl -X GET 'http://localhost:8080/db/dz/get/two'
curl -X GET 'http://localhost:8080/db/dz/get/three'