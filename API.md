#API
## Error
Example:
```json
{
  "status": "ERROR",
  "error":  "Variable is not exist"
}
```

## Instances commands
###Create a new instance
####Request
- Path: /create
- Method: POST
- Body:
```json
{"name": "instance_name"}
```
####Respond
```json
{"status":"OK"}
```

###List of an instance
####Request
- Path: /list
- Method: GET
####Respond
```json
{
  "status": "OK",
  "count":  2,
  "names":  [
    "instnce_name_1",
    "instnce_name_2",
    "instnce_name_n"
  ]
}
```

###Remove an instance
####Request
- Path: /destroy
- Method: DELETE
####Respond
```json
{"status":"OK"}
```

## String commands
###Get a string
####Request
- Path: /in/:intrance_name/get/:variable_name
- Method: GET
####Respond
```json
{"value":"1"}
```

###Delete
####Request
- Path: /in/:intrance_name/del/:variable_name
- Method: DELETE
####Respond
```json
{"status":"OK"}
```

###Set a string
####Request
- Path: /in/:intrance_name/set
- Method: POST
- Body:
```json
{
  "name": "variable_name",
  "value": "value"
}
```
####Respond
```json
{"status":"OK"}
```

## Time to Live commands
###Set the expire time (TTL)
####Request
- Path: /in/:intrance_name/ttl/set
- Method: POST
- Body:
```json
{
    "name": "variable_name",
    "ttl": 42
}
```
####Respond
```json
{"status":"OK"}
```

###Remove the expire time (TTL)
####Request
- Path: /in/:intrance_name/ttl/del
- Method: DELETE
- Body:
```json
{"name": "variable_name"}
```
####Respond
```json
{"status":"OK"}
```

## Array commands
###Set an array
####Request
- Path: /in/:instance_name/arr/set
- Method: POST
- Body:
```json
{
  "name":   "array_name",
  "value":  ["zero", "one", "two"],
  "ttl":    0
}
```
####Respond
```json
{"status":"OK"}
```

###Get an element of an array
####Request
- Path: /in/:intrance_name/arr/el/get
- Method: GET
- Body:
```json
{
  "name": "array_name",
  "index": 2
}
```
####Respond
```json
{
  "status": "OK",
  "value":  "string"
}
```

###Add a string to an array
####Request
- Path: /in/:intrance_name/arr/el/add
- Method: POST
- Body:
```json
{
  "name": "array_name",
  "value": "value"
}
```
####Respond
```json
{"status":"OK"}
```

###Set/Update an element of an array
####Request
- Path: /in/:intrance_name/arr/el/set
- Method: POST
- Body:
```json
{
  "name": "variable_name",
  "index": 2,
  "value": "string"
}
```
####Respond
```json
{"status":"OK"}
```

###Remove an element of an array
####Request
- Path: /in/:intrance_name/arr/el/del
- Method: DELETE
- Body:
```json
{
  "name": "variable_name",
  "index": 0
}
```
####Respond
```json
{"status":"OK"}
```

## Hash commands
###Set a hash of strings.
####Request
- Path: /in/:intrance_name/hash/set
- Method: POST
- Body:
```json
{
  "name": "hash_name",
  "value": {
      "key1": "string1",
      "key2": "string2"
  }
}
```
####Respond
```json
{"status":"OK"}
```

###HASH/EL/GET
####Request
- Path: /in/:intrance_name/hash/el/get
- Method: GET
- Body:
```json
{
  "name": "hash_name",
  "key": "key_name"
}
```
####Respond
```json
{
  "status": "OK",
  "value":  "string"
}
```

###HASH/EL/SET
####Request
- Path: /in/:intrance_name/hash/el/set
- Method: POST
- Body:
```json
{
  "name":  "hash_name",
  "key":   "key_name",
  "value": "string"
}
```
####Respond
```json
{"status":"OK"}
```

###Remove an element of a hash
####Request
- Path: /in/:intrance_name/hash/el/del
- Method: DELETE
- Body:
```json
{
  "name":  "hash_name",
  "key":   "key_name"
}
```
####Respond
```json
{"status":"OK"}
```