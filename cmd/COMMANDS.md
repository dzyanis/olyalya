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
HELP cmd
```

```
EXIT
```

## Instances
```
LIST
```

```
CREATE name
```

```
SELECT name
```


```
DESTROY name
```

## Main command
```
GET name
```

```
DEL name
```

```
SET name "value" ttl
```

## Time to Live
```
TTL/SET name seconds
```

```
TTL/DEL name
```

## Array
```
ARR/GET name
```

```
ARR/SET name ["a", "b", "c"] ttl
```

```
ARR/EL/ADD name value
```

```
ARR/EL/GET name index
```

```
ARR/EL/SET name index value
```

```
ARR/EL/DEL name index
```

## Hash
```
HASH/GET name
```

```
HASH/SET name {"key": "value"} ttl
```

```
HASH/EL/GET name key
```

```
HASH/EL/SET name key value
```

```
HASH/EL/DEL name key
```