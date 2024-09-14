.PHONY: build
build:
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64  go build -ldflags="-s -w" -o dist/liblux-aarch64.dylib  -buildmode=c-shared .
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64  go build -ldflags="-s -w" -o dist/liblux.dylib  -buildmode=c-shared .
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/liblux.dll -buildmode=c-shared .
