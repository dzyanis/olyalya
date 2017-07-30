# O(lya-lya)

OLL is simple implementation of Redis-like in-memory cache database with HTTP/JSON interface.
It's not the best database, but good example of.

It includes three part:
- [Server](olldb/API.md)
- [Client Library](client/)
- [Commondline Client](oll-cli/COMMANDS.md)


##How to run
Install and launch database:
```bash
go install github.com/dzyanis/olyalya/olldb
$GOPATH/bin/olldb
```

Install and launch commandline client:
```bash
go install github.com/dzyanis/olyalya/oll-cli
$GOPATH/bin/oll-cli
```


Backlog:
- add config
- persistence to disk/db
- scaling(on server-side or on client-side, up to you)
- auth
- perfomance tests
