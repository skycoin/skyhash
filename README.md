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
