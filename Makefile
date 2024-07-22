BIN := 403unlocker-go

$(BIN): $(wildcard *.go) $(wildcard */*.go)
	CGO_ENABLED=0 go build -ldflags='-d -buildid=' .

clean:
	rm -rf $(BIN) result-*.json
	go clean
