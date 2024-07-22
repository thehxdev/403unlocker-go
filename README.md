# 403unlocker-go
403unlocker-go is a [Golang](https://go.dev) implementation of [403unlocker](https://github.com/403unlocker) project.


## Build
To build from source-code, you need go compiler and GNU make. Then:
```bash
make
```
The command above will produce `403unlocker-go` executable.

### Manual build
You can use `go` compiler directly to build project:
```bash
CGO_ENABLED=0 go build -ldflags='-d -buildid=' .
```


## Usage
To see a help message:
```bash
./403unlocker-go -help
```

To start testing:
```bash
./403unlocker-go -c config.json
```
