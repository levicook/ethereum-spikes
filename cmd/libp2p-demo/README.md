
# libp2p-demo

These are mainly "notes to self"

## Ping

1. Start a "server" (long running node) with `libp2p-demo ping-server`
1. Ping the "server" with `libp2p-demo ping-client` using the addr printed by the server

Output from ping-server:
```bash
$ libp2p-demo ping-server
host addr: /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe
Waiting for shutdown signals...
```

Output from ping-client:
```bash
$ libp2p-demo ping-client -c 10 -p /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe
pinging /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 406.77µs
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 96.265µs
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 91.844µs
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 105.481µs
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 89.553µs
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 101.605µs
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 89.189µs
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 74.429µs
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 117.549µs
pinged  /ip4/127.0.0.1/tcp/57642/p2p/QmRbEZY1pnMGREjEnx3fyFff4vxKiyGp7wEj8qWNBvaKEe rtt: 98.74µs
```
