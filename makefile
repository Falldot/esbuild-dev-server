jstest:
	cd example && yarn add -D ../npm/esbuild-dev-server && yarn add -D ../npm/esbuild-dev-server-win32-x64 && yarn js-start

gotest:
	cd example && yarn go-start

dts:
	cd lib && yarn build

build:
	go run cmd/build/build.go

build-all:
	make build \
		build-win-32 \
		build-win-64 \
		build-win-arm-64 \
		build-linux-32 \
		build-linux-64 \
		build-linux-arm \
		build-linux-arm-64 \
		build-darwin-64 \
		build-darwin-arm-64

build-win-32:
	GOOS=windows GOARCH=386 go build -ldflags "-s -w" cmd/devserver/devserver.go
	mv devserver.exe npm/esbuild-dev-server-win32-x32

build-win-64:
	GOOS=windows GOARCH=386 go build -ldflags "-s -w" cmd/devserver/devserver.go
	mv devserver.exe npm/esbuild-dev-server-win32-x64

build-win-arm-64:
	GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" cmd/devserver/devserver.go
	mv devserver.exe npm/esbuild-dev-server-win32-arm64

build-linux-32:
	GOOS=linux GOARCH=386 go build -ldflags "-s -w" cmd/devserver/devserver.go
	mv devserver npm/esbuild-dev-server-linux-x32

build-linux-64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" cmd/devserver/devserver.go
	mv devserver npm/esbuild-dev-server-linux-x64

build-linux-arm:
	GOOS=linux GOARCH=arm go build -ldflags "-s -w" cmd/devserver/devserver.go
	mv devserver npm/esbuild-dev-server-linux-arm

build-linux-arm-64:
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" cmd/devserver/devserver.go
	mv devserver npm/esbuild-dev-server-linux-arm64

build-darwin-64:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" cmd/devserver/devserver.go
	mv devserver npm/esbuild-dev-server-darwin-x64

build-darwin-arm-64:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" cmd/devserver/devserver.go
	mv devserver npm/esbuild-dev-server-darwin-arm64

publish:
	cd npm/esbuild-dev-server && npm publish
	cd npm/esbuild-dev-server-win32-x32 && npm publish
	cd npm/esbuild-dev-server-win32-x64 && npm publish
	cd npm/esbuild-dev-server-win32-arm64 && npm publish
	cd npm/esbuild-dev-server-linux-x32 && npm publish
	cd npm/esbuild-dev-server-linux-x64 && npm publish
	cd npm/esbuild-dev-server-linux-arm && npm publish
	cd npm/esbuild-dev-server-linux-arm64 && npm publish
	cd npm/esbuild-dev-server-darwin-x64 && npm publish
	cd npm/esbuild-dev-server-darwin-arm64 && npm publish