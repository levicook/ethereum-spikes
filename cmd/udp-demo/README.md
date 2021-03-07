
# udp-demo

These are mainly "notes to self"

## Echo

1. Start a "server" (long running node) with `udp-demo ping-server`
1. Send messages to the server with `nc 127.0.0.1 10000`

echo-server I/O:
```bash
$ udp-demo echo-server
Waiting for shutdown signals...
oy
you there?
okay, cool.
quit
bye!
Received signal, shutting down...
```

nc I/O:
```bash
$ nc 127.0.0.1 10000
oy
you there?
okay, cool.
quit
^C
$
```
