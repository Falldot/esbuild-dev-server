package main

import (
	"log"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"lib/src/esbuild-dev-server.ts"},
		Bundle:      true,
		Platform:    api.PlatformNode,
		Engines: []api.Engine{
			{api.EngineNode, "14.18"},
		},
		Tsconfig: "lib/tsconfig.json",
		Write:    true,
		Outdir:   "npm/lib",
	})

	if len(result.Errors) > 0 {
		log.Fatalln(result.Errors)
	}
}
