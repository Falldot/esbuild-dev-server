# esbuild-dev-server

## Only win64

> This plugin allows you to start a local server with hot reloading with [Esbuild](https://esbuild.github.io/)

## Installation
### npm
```
npm i esbuild-dev-server -D
```
### yarn
```
yarn add esbuild-dev-server -D
```
### go
```
go get github.com/Falldot/esbuild-dev-server
```
## Configuration

- `options.Port`, `string`: local server start port.
- `options.Index`, `string`: path to index html file.
- `options.StaticDir`, `string`: path to static files (js, css, img ...).
- `options.WatchDir`, `string`: path to working directory.
- `options.OnLoad`, `() => void`: local server restart event.

## How to use?
### Node.js
```js
const {build, formatMessages} = require("esbuild");
const {esBuildDevServer, startServer, sendError, sendReload} = require("esbuild-dev-server");

;(async () => {
	const builder = await build({
		entryPoints: ["src/index.js"],
		bundle: true,
		minify: false,
		sourcemap: true,
		target: ['chrome58', 'firefox57', 'safari11', 'edge16'],
		outdir: "dist/js",
		incremental: true,
		plugins: [
			esBuildDevServer({
				Port: "8080",
				Index: "dist/index.html",
				StaticDir: "dist",
				WatchDir: "src",
				OnLoad: async () => {
					try {
						await builder.rebuild();
						await sendReload();
					} catch(result) {
						let str = await formatMessages(result.errors, {kind: 'error', color: true});
						await sendError(str.join(""));
					}
				}
			})
		],
	});
	await startServer();
})();
```
### Golang
```go
package main

import (
	"log"

	devserver "github.com/Falldot/esbuild-dev-server"
	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	var result api.BuildResult

	result = api.Build(api.BuildOptions{
		EntryPoints:       []string{"src/index.js"},
		Bundle:            true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Outdir:      "dist/js",
		Write:       true,
		Engines: []api.Engine{
			{api.EngineChrome, "58"},
			{api.EngineFirefox, "57"},
			{api.EngineSafari, "11"},
			{api.EngineEdge, "16"},
		},
		Incremental: true,
		Plugins: []api.Plugin{
			devserver.Plugin(devserver.Options{
				Port:      ":8080",
				Index:     "dist/index.html",
				StaticDir: "dist",
				WatchDir:  "src",
				OnReload: func() {
					result.Rebuild()
				},
			}),
		},
	})
	if len(result.Errors) > 0 {
		log.Fatalln(result.Errors)
	}

	if err := devserver.Start(); err != nil {
        log.Fatalln(result.Errors)
    }
}
```