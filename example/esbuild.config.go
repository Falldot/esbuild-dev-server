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
			Port:            "8080",            // optional, default: 8080
			WatchDir:        "src",             // optional, default: "src"
			Index:           "dist/index.html", // optional
			StaticDir:       "dist",            // optional
			OnBeforeRebuild: func() {},         // optional
			OnAfterRebuild:  func() {},         // optional
		},
	)
}
