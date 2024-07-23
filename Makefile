BIN := 403unlocker-go
BUILD_DIR := ./build

$(BIN): $(wildcard *.go) $(wildcard */*.go)
	CGO_ENABLED=0 go build -ldflags='-d -buildid=' .

cross-plat:
	@rm -rf $(BUILD_DIR)
	python3 build.py

clean:
	rm -rf $(BIN) result-*.json $(BUILD_DIR)
	go clean
