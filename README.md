# O(lya-lya)

OLL is simple implementation of Redis-like in-memory cache database with HTTP/JSON interface.
It's not the best database, but good example of.

It includes three part:
- [Server](server/API.md)
- [Client Library](pkg/client/)
- [Commondline Client](cli/COMMANDS.md)


## How to run
Install and launch database:
```bash
go install github.com/dzyanis/olyalya/server
$GOPATH/bin/server
```

Install and launch commandline client:
```bash
go install github.com/dzyanis/olyalya/cli
$GOPATH/bin/cli
```

## Tasks
### Done
- migrate from [Glide](http://bumptech.github.io/glide/) to [Dep](https://github.com/golang/dep)
### Doing
### Backlog
- add config
- persistence to disk/db
- auth
- perfomance tests
- raft
- scaling(on server-side or on client-side, up to you)
