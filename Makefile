export CGO_ENABLED := 0
help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build: ## "Building a valid production binary"
	echo "Building a valid production binary"
	go build -o yfapi

lib: ## "Building a static linux library for production"
	echo "Building a static linux binary for production"
	CGO_ENABLED=1 GOOS=linux go build -v -o libyfapi.so -buildmode=c-shared ./cmd/libyfapi

bu: ## "Looks like we have a binary ready to debug"
	echo "Looks like we have a binary ready to debug"
	go build  -gcflags="all=-N -l" -o yfapi

test: ## "Run tests"
	echo "Run tests"
	go clean -testcache
	grc go test -v -cover ./...
	python3 -c 'import ctypes;lib = ctypes.CDLL("./libyfapi.so");lib.GetHistory.restype = ctypes.c_char_p;print(lib.GetHistory(b"AAPL").decode("utf-8")[:100] + "...")'
	python3 -c 'import ctypes;lib = ctypes.CDLL("./libyfapi.so");lib.GetInfo.restype = ctypes.c_char_p;print(lib.GetInfo(b"AAPL").decode("utf-8")[:100] + "...")'
	python3 -c 'import ctypes;lib = ctypes.CDLL("./libyfapi.so");lib.GetQuote.restype = ctypes.c_char_p;print(lib.GetQuote(b"AAPL").decode("utf-8")[:100] + "...")'
	python3 -c 'import ctypes;lib = ctypes.CDLL("./libyfapi.so");lib.GetOptionChain.restype = ctypes.c_char_p;print(lib.GetOptionChain(b"AAPL").decode("utf-8")[:100] + "...")'
	python3 -c 'import ctypes;lib = ctypes.CDLL("./libyfapi.so");lib.GetOptionChainByExpiration.restype = ctypes.c_char_p;print(lib.GetOptionChainByExpiration(b"AAPL", b"2025-12-22").decode("utf-8")[:100] + "...")'
	python3 -c 'import ctypes;lib = ctypes.CDLL("./libyfapi.so");lib.GetExpirationDates.restype = ctypes.c_char_p;print(lib.GetExpirationDates(b"AAPL").decode("utf-8")[:100] + "...")'
	
clean: ## "Cleans the local binaries"
	echo "Cleans the local binaries"
	rm -f yfapi
	rm -f libyfapi.so

all: lib test