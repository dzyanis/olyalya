#Client commands
## Common commands
###HELP
Print information about command
```
HELP COMMAND_NAME
```

### ECHO
Output strings
```
ECHO "Hello world"
```

Close client
```
EXIT
```

## Instances commands
###LIST
Show the list of instances
```
LIST
```

###CREATE
Create a new instance
```
CREATE name
```

###SELECT
Select an instance
```
SELECT name
```

###DESTROY
Remove an instance
```
DESTROY name
```

## String commands
###GET
Get a string of a key
```
GET name
```

###DEL
Delete a key
```
DEL name
```

###SET
Set the string value of a key.
TTL(time to live)  - set the expire time, in seconds.
```
SET name "value" ttl
```

## Time to Live commands
###TTL/SET
Set the expire time of a key, in seconds.
```
TTL/SET name seconds
```
###TTL/DEL
Remove the expire time of a key, in seconds.
```
TTL/DEL name
```

## Array commands
###ARR/SET
Set the array of a string.

TTL(time to live)  - set the expire time, in seconds.
```
ARR/SET name ["a", "b", "c"] ttl
```
###ARR/GET
Get the array of a string.
```
ARR/GET name
```
###ARR/EL/ADD
Add new string to an array
```
ARR/EL/ADD name value
```

###ARR/EL/GET
Get an element of an array
```
ARR/EL/GET name index
```

###ARR/EL/SET
Set an element of an array
```
ARR/EL/SET name index value
```

###ARR/EL/DEL
Remove an element of an array
```
ARR/EL/DEL name index
```

## Hash commands
###HASH/SET
Set a hash of strings.
```
HASH/SET name {"key": "value"} ttl
```

###HASH/GET
Get a hash of strings
```
HASH/GET name
```

###HASH/EL/GET
Get an element of a hash.
```
HASH/EL/GET name key
```

###HASH/EL/SET
Set an element of a hash.
```
HASH/EL/SET name key value
```

###HASH/EL/DEL
Delete an element of a hash.
```
HASH/EL/DEL name key
```