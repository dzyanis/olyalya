#Client commands
# Data Bases
```
HELP
```

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
DEL name
```

# Main command
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
TTL name second
```

```
HAS name
```
## Array
```
ARR/SET name [1,2,3] ttl
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
HASH/SET name {"one": "1"} ttl
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