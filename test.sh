echo 'CREATE NEW INSTANCE'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/create' -d '{"name": "dz"}'
echo 'GET INFO ABOUT INSTANCE'
curl -X GET 'http://localhost:8080/db/dz'

echo 'SET STRING'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/dz/set' -d '{
    "name": "one",
    "value": "1"
}'
curl -X GET 'http://localhost:8080/db/dz/get/one'

echo 'SET ARRAY'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/dz/arr/set' -d '{
    "name": "two",
    "value": ["ABC", "DEF", "GHI"]
}'
curl -X GET 'http://localhost:8080/db/dz/get/two'

echo 'SET HASH'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/dz/hash/set' -d '{
    "name": "three",
    "value": {"name": "Dzyanis", "year": "1987"}
}'
curl -X GET 'http://localhost:8080/db/dz/get/three'

echo 'DELETE STRING'
curl -X DELETE 'http://localhost:8080/db/dz/delete/one'
curl -X GET 'http://localhost:8080/db/dz/get/one'

echo 'DELETE ARRAY'
curl -X DELETE 'http://localhost:8080/db/dz/delete/two'
curl -X GET 'http://localhost:8080/db/dz/get/two'

echo 'DELETE HASH'
curl -X DELETE 'http://localhost:8080/db/dz/delete/three'
curl -X GET 'http://localhost:8080/db/dz/get/three'