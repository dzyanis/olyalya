#Client commands
## Common
###ECHO
Example:
```ECHO "Hello world"```

## Data Bases
```HELP cmd```

```DB/LIST```

```DB/CREATE name```

```DB/SELECT name```

```DB/DEL name```

## Main command
```GET name```

```DEL name```

```SET name "value"```

```TTL name seconds```

```HAS name```

## Array
```ARR/SET name ["a", "b", "c"] ttl```

```ARR/INDEX/GET name index```

```ARR/INDEX/SET name index value```

```ARR/INDEX/DEL name index```

## Hash
```HASH/SET name {"key": "value"} ttl```

```HASH/KEY/GET name key```

```HASH/KEY/SET name key value```

```HASH/KEY/DEL name key```