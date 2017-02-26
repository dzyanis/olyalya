# O(lya-lya)

OLL is simple implementation of Redis-like in-memory cache database with HTTP/JSON interface.
It's not the best database, but good example of.

It includes three part:
- [Server](server.go)
- [Library](API.md)
- [Commond line](COMMANDS.md)


Backlog:
- persistence to disk/db
- scaling(on server-side or on client-side, up to you)
- auth
- perfomance tests