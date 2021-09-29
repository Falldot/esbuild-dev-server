# esbuild-dev-server

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
		// Important field!!!!!
		incremental: true,
		plugins: [
			// esbuild-dev-server plugin
			esBuildDevServer({
				Port: "8080", 				// Port start local server
				Index: "dist/index.html",	// Root html file
				// Stacic files
                // dist/js, dist/css ...
				StaticDir: "dist",
				// Working directory
				WatchDir: "src",
				// Event reload local server
				OnLoad: async () => {
					// Error handler and rebuild esbuild
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
	// Starting...
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
		// Important field!!!!!
		Incremental: true,
		Plugins: []api.Plugin{
            // esbuild-dev-server plugin
			devserver.Plugin(devserver.Options{
				Port:      ":8080",           // Port start local server
				Index:     "dist/index.html", // Root html file
                // Stacic files
                // dist/js, dist/css ...
				StaticDir: "dist",
                // Working directory
				WatchDir:  "src",
                // Event reload local server
				OnReload: func() {
					result.Rebuild()
				},
			}),
		},
	})
	if len(result.Errors) > 0 {
		log.Fatalln(result.Errors)
	}

	// Starting...
	if err := devserver.Start(); err != nil {
        log.Fatalln(result.Errors)
    }
}
```