#Client commands
## Common
Example:
```
ECHO "Hello world"
```

```
HELP cmd
```

```
EXIT
```

## Instances
```
INST/LIST
```

```
INST/CREATE name
```

```
INST/SELECT name
```


```
INST/DESTROY name
```

## Main command
```
GET name
```

```
DEL name
```

```
SET name "value"
```

```
TTL name seconds
```

```
HAS name
```

## Array
```
ARR/SET name ["a", "b", "c"] ttl
```

```
ARR/INDEX/GET name index
```

```
ARR/INDEX/SET name index value
```

```
ARR/INDEX/DEL name index
```

## Hash
```
HASH/SET name {"key": "value"} ttl
```

```
HASH/KEY/GET name key
```

```
HASH/KEY/SET name key value
```

```
HASH/KEY/DEL name key
```