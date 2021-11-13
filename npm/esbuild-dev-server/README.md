# esbuild-dev-server

This plugin allows you to start a local server with hot reloading for [Esbuild](https://esbuild.github.io/)

More community [plugins](https://github.com/esbuild/community-plugins)

## Installation
`npm`
```
npm i esbuild-dev-server -D
```
`yarn`
```
yarn add esbuild-dev-server -D
```
`go`
```
go get github.com/Falldot/esbuild-dev-server
```
## Configuration

- `options.port`, `string`: local server start port.
- `options.index`, `string`: path to index html file.
- `options.staticDir`, `string`: path to static files (js, css, img ...).
- `options.watchDir`, `string`: path to working directory.
- `options.onBeforeRebuild`, `() => void`: event before rebuild.
- `options.onAfterRebuild`, `() => void`: event after rebuild.

## How to use?
### Node.js
`esbuild.config.js`
```js
const {build} = require("esbuild")
const esBuildDevServer = require("esbuild-dev-server")

esBuildDevServer.start(
	build({
		entryPoints: ["src/index.js"],
		outdir: "dist",
		incremental: true,
		// and more options ...
	}),
	{
		port:      "8080", // optional, default: 8080
		watchDir:  "src", // optional, default: "src"
		index:     "dist/index.html", // optional
		staticDir: "dist", // optional
		onBeforeRebuild: {}, // optional
		onAfterRebuild:  {}, // optional
	}
)
```
`package.json`
```json
"scripts": {
    "dev": "node esbuild.config.js",
},
```
`dist/index.html`
```html
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<div id="root"></div>
		<script src="index.js"></script>
	</body>
</html>
```
### Golang
`esbuild.config.go`
```go
package main

import (
	devserver "github.com/Falldot/esbuild-dev-server"
	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	devserver.Start(
		api.Build(api.BuildOptions{
			EntryPoints: []string{"src/index.js"},
			Outdir:      "dist",
			Incremental: true,
			// and more options ...
		}),
		devserver.Options{
			Port:      "8080", // optional, default: 8080
			WatchDir:  "src", // optional, default: "src"
			Index:     "dist/index.html", // optional
			StaticDir: "dist", // optional
			OnBeforeRebuild: func() {}, // optional
			OnAfterRebuild:  func() {}, // optional
		},
	)
}
```
`package.json`
```json
"scripts": {
    "dev": "go esbuild.config.go",
},
```