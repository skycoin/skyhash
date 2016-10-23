# skyhash

### Normal install

```
$ go get -u github.com/skycoin/skyhash
$ go install github.com/skycoin/skyhash
$ $GOPATH/bin/skyhash
```
### Development on docker

Compile image:

```
sudo docker build -t skyhash .
```

Run:

```
sudo docker run -it --net=host --rm -p 6481:6481 --name skyhash-instance skyhash
```


## Examples

### Create node:

```
POST http://127.0.0.1:6481/nodemanager/nodes/start
```

Curl example:

```
curl \
-X POST \
http://127.0.0.1:6481/nodemanager/nodes/start
```

This endpoint will return the id of the new node just created.

e.g.

```
1
```

### List all nodes:

```
GET http://127.0.0.1:6481/nodemanager/nodes
```

Curl example:

```
curl http://127.0.0.1:6481/nodemanager/nodes
```
This endpoint will return a map of nodes with some info about each of them.

e.g.

```json
{
    "1": {
        "Address": "",
        "Port": 6061
    },
    "2": {
        "Address": "",
        "Port": 6062
    },
    "3": {
        "Address": "",
        "Port": 6063
    },
    "4": {
        "Address": "",
        "Port": 6064
    },
    "5": {
        "Address": "",
        "Port": 6065
    },
    "6": {
        "Address": "",
        "Port": 6066
    },
    "7": {
        "Address": "",
        "Port": 6067
    }
}
```

### List transports for a specific node (by node id):

```
GET http://127.0.0.1:6481/nodemanager/transports?id=1
```

Curl example:

```
curl http://127.0.0.1:6481/nodemanager/transports?id=1
```

This endpoint will return an array of addresses to which this particular node is connected to.

e.g.

```
[
    "127.0.0.1:6067",
    "127.0.0.1:6064",
    "127.0.0.1:6065"
]
```

### Create a new transport (connect a node to another node):

```
POST http://127.0.0.1:6481/nodemanager/transports?id=1

Payload:

{
  "ip":"127.0.0.1",
  "port":"6064"
}
```

Curl example:

```
curl \
-H "Content-Type: application/json" \
-X POST \
-d '{"ip":"127.0.0.1","port":"6064"}' \
http://127.0.0.1:6481/nodemanager/transports?id=1
```

This endpoint won't return a response body.
