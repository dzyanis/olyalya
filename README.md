# O(lya-lya)

OLL is simple implementation of Redis-like in-memory cache database with HTTP/JSON interface.
It's not the best database, but good example of.

It includes three part:
- [Server](server/API.md)
- [Client Library](pkg/client/)
- [Commandline Client](cli/COMMANDS.md)


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
- migrate from [Glide](https://github.com/Masterminds/glide) to [Dep](https://github.com/golang/dep)
- builds and versions
### Doing
### Backlog
- use logger
- add config
- persistence to disk/db
- auth
- perfomance tests
- [raft](https://raft.github.io/)
- scaling(on server-side or on client-side, up to you)
- pass [Go Report Card](https://goreportcard.com/report/github.com/dzyanis/olyalya)
