jstest:
	cd example && yarn add -D ../npm && yarn start

gotest:
	cd example && yarn gostart

dts:
	cd lib && yarn build

build:
	go run cmd/build/build.go
	go build cmd/devserver/devserver.go
	move devserver.exe npm

npm:
	cd npm && npm publish