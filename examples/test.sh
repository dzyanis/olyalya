echo 'CREATE NEW INSTANCE'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/create' -d '{"name": "dz"}'

echo 'GET INFO ABOUT INSTANCE'
curl -X GET 'http://localhost:3000/in/dz'

echo 'List of instances'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:3000/list'

echo 'SET STRING'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/set' -d '{"name": "one", "value": "1"}'
curl -X GET 'http://localhost:3000/in/dz/get/one'

echo 'SET ARRAY'
curl -X GET 'http://localhost:3000/in/dz/get/numbers'
#curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/arr/set' -d '{"name": "numbers", "value": ["zero", "one", "two"]}'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/arr/set' -d '{"name": "numbers", "value": ["zero", "one", "two"], "ttl": 0}'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/arr/el/get' -d '{"name": "numbers", "index": 2}'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/arr/el/add' -d '{"name": "numbers", "value": "three"}'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/arr/el/get' -d '{"name": "numbers", "index": 3}'
curl -X DELETE -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/arr/el/del' -d '{"name": "numbers", "index": 0}'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/arr/el/get' -d '{"name": "numbers", "index": 0}'

echo 'SET HASH'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/hash/set' -d '{"name": "author", "value": {"name": "Dzyanis", "year": "1987"}}'
curl -X GET 'http://localhost:3000/in/dz/get/author'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/hash/el/set' -d '{"name": "author", "key": "profession", "value": "Back-end Developer"}'
curl -X GET -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/hash/el/get' -d '{"name": "author", "key": "profession"}'
curl -X DELETE -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/hash/el/del' -d '{"name": "author", "key": "year"}'
curl -X GET 'http://localhost:3000/in/dz/get/author'

echo 'DELETE STRING'
curl -X DELETE 'http://localhost:3000/in/dz/delete/one'
curl -X GET 'http://localhost:3000/in/dz/get/one'

echo 'DELETE ARRAY'
curl -X DELETE 'http://localhost:3000/in/dz/delete/numbers'
curl -X GET 'http://localhost:3000/in/dz/get/numbers'

echo 'DELETE HASH'
curl -X DELETE 'http://localhost:3000/in/dz/delete/author'
curl -X GET 'http://localhost:3000/in/dz/get/author'

echo 'TTL'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/set' -d '{"name": "one", "value": "1", "ttl": 2}'
curl -X GET 'http://localhost:3000/in/dz/get/one'
echo '... pause 2 sec' && sleep 2
curl -X GET 'http://localhost:3000/in/dz/get/one'

echo 'TTL SET'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/set' -d '{"name": "one", "value": "1"}'
curl -X GET 'http://localhost:3000/in/dz/get/one'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/ttl/set' -d '{"name": "one", "ttl": 2}'
curl -X GET 'http://localhost:3000/in/dz/get/one'
echo '... pause 2 sec' && sleep 2
curl -X GET 'http://localhost:3000/in/dz/get/one'

echo 'TTL DEL'
curl -X POST -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/set' -d '{"name": "one", "value": "1", "ttl": 3}'
curl -X GET 'http://localhost:3000/in/dz/get/one'
curl -X DELETE -H 'Content-Type: application/json' 'http://localhost:3000/in/dz/ttl/del' -d '{"name": "one"}'
curl -X GET 'http://localhost:3000/in/dz/get/one'
echo '... pause 2 sec' && sleep 3
curl -X GET 'http://localhost:3000/in/dz/get/one'