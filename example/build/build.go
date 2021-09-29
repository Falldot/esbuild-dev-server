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
		Engines: []api.Engine{
			{Name: api.EngineChrome, Version: "90"},
		},
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
		Incremental: true,
		Outdir:      "dist/js",
		Write:       true,
	})
	if len(result.Errors) > 0 {
		log.Fatalln(result.Errors)
	}

	if err := devserver.Start(); err != nil {
		log.Fatalln(result.Errors)
	}
}
