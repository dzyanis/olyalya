echo 'CREATE NEW INSTANCE'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/in/create' -d '{"name": "dz"}'

echo 'GET INFO ABOUT INSTANCE'
curl -X GET 'http://localhost:8080/inst/dz'

echo 'CREATE NEW INSTANCE'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/list' -d '{"name": "dz"}'

echo 'SET STRING'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/set' -d '{"name": "one", "value": "1"}'
curl -X GET 'http://localhost:8080/inst/dz/get/one'

echo 'SET ARRAY'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/set/arr' -d '{"name": "numbers", "value": ["0", "1", "2"]}'
curl -X GET 'http://localhost:8080/inst/dz/get/numbers'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/arr/index/get' -d '{"name": "numbers", "index": 2}'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/arr/index/add' -d '{"name": "numbers", "value": "3"}'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/arr/index/get' -d '{"name": "numbers", "index": 3}'
curl -X DELETE -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/arr/index/del' -d '{"name": "numbers", "index": 0}'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/arr/index/get' -d '{"name": "numbers", "index": 0}'

echo 'SET HASH'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/set/hash' -d '{"name": "author", "value": {"name": "Dzyanis", "year": "1987"}}'
curl -X GET 'http://localhost:8080/inst/dz/get/author'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/hash/key/set' -d '{"name": "author", "key": "profession", "value": "Back-end Developer"}'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/hash/key/get' -d '{"name": "author", "key": "profession"}'
curl -X DELETE -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/hash/key/del' -d '{"name": "author", "key": "year"}'
curl -X GET 'http://localhost:8080/inst/dz/get/author'

echo 'DELETE STRING'
curl -X DELETE 'http://localhost:8080/inst/dz/delete/one'
curl -X GET 'http://localhost:8080/inst/dz/get/one'

echo 'DELETE ARRAY'
curl -X DELETE 'http://localhost:8080/inst/dz/delete/numbers'
curl -X GET 'http://localhost:8080/inst/dz/get/numbers'

echo 'DELETE HASH'
curl -X DELETE 'http://localhost:8080/inst/dz/delete/author'
curl -X GET 'http://localhost:8080/inst/dz/get/author'

echo 'TTL'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/set' -d '{"name": "one", "value": "1", "ttl": 2}'
curl -X GET 'http://localhost:8080/inst/dz/get/one'
echo '... pause 2 sec' && sleep 2
curl -X GET 'http://localhost:8080/inst/dz/get/one'

echo 'TTL SET'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/set' -d '{"name": "one", "value": "1"}'
curl -X GET 'http://localhost:8080/inst/dz/get/one'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/ttl/set' -d '{"name": "one", "ttl": 2}'
curl -X GET 'http://localhost:8080/inst/dz/get/one'
echo '... pause 2 sec' && sleep 2
curl -X GET 'http://localhost:8080/inst/dz/get/one'

echo 'TTL DEL'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/set' -d '{"name": "one", "value": "1", "ttl": 3}'
curl -X GET 'http://localhost:8080/inst/dz/get/one'
curl -X DELETE -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/ttl/del' -d '{"name": "one"}'
curl -X GET 'http://localhost:8080/inst/dz/get/one'
echo '... pause 2 sec' && sleep 3
curl -X GET 'http://localhost:8080/inst/dz/get/one'