# 403unlocker-go
403unlocker-go is a [Golang](https://go.dev) implementation of [403unlocker](https://github.com/403unlocker) project.


## Download
Check [releases page](https://github.com/thehxdev/403unlocker-go/releases/latest) to download binary packages.


## Build

### Linux / macOS
```bash
CGO_ENABLED=0 go build -ldflags='-d -buildid=' .
```

### Windows
```powershell
$env:CGO_ENABLED=0
go build -ldflags='-d -buildid=' .
```

### With Makefile
```bash
make
```

### Cross-Platform compilation
```bash
make cross-plat
```


## Usage
To print a help message:
```bash
./403unlocker-go -help
```

To start testing:
```bash
./403unlocker-go -c config.json
```
