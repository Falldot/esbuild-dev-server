jstest:
	cd example && yarn add -D ../npm/esbuild-dev-server && yarn add -D ../npm/esbuild-dev-server-win32-x64 && yarn start

gotest:
	cd example && yarn gostart

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
	set GOOS=windows
	set GOARCH=386
	go build cmd/devserver/devserver.go
	move devserver.exe npm/esbuild-dev-server-win32-x32

build-win-64:
	set GOOS=windows
	set GOARCH=amd64
	go build cmd/devserver/devserver.go
	move devserver.exe npm/esbuild-dev-server-win32-x64

build-win-arm-64:
	set GOOS=windows
	set GOARCH=arm64
	go build cmd/devserver/devserver.go
	move devserver.exe npm/esbuild-dev-server-win32-arm64

build-linux-32:
	set GOOS=linux
	set GOARCH=386
	go build cmd/devserver/devserver.go
	move devserver.exe npm/esbuild-dev-server-linux-x32

build-linux-64:
	set GOOS=linux
	set GOARCH=amd64
	go build cmd/devserver/devserver.go
	move devserver.exe npm/esbuild-dev-server-linux-x64

build-linux-arm:
	set GOOS=linux
	set GOARCH=arm
	go build cmd/devserver/devserver.go
	move devserver.exe npm/esbuild-dev-server-linux-arm

build-linux-arm-64:
	set GOOS=linux
	set GOARCH=arm64
	go build cmd/devserver/devserver.go
	move devserver.exe npm/esbuild-dev-server-linux-arm64

build-darwin-64:
	set GOOS=darwin
	set GOARCH=amd64
	go build cmd/devserver/devserver.go
	move devserver.exe npm/esbuild-dev-server-darwin-x64

build-darwin-arm-64:
	set GOOS=darwin
	set GOARCH=arm64
	go build cmd/devserver/devserver.go
	move devserver.exe npm/esbuild-dev-server-darwin-arm64

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