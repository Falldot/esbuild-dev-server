package devserver

import (
	"log"

	plugin "github.com/Falldot/esbuild-dev-server/internal/api"
	"github.com/evanw/esbuild/pkg/api"
)

type Options plugin.DevServer

var server plugin.DevServer

func Start() error {
	return server.Start()
}

func Plugin(options Options) api.Plugin {
	server = plugin.DevServer(options)
	return api.Plugin{
		Name: "dev-server",
		Setup: func(pb api.PluginBuild) {
			pb.OnEnd(func(result *api.BuildResult) {
				if server.IsStart {
					if len(result.Errors) > 0 {
						strs := api.FormatMessages(result.Errors, api.FormatMessagesOptions{
							Kind:  api.ErrorMessage,
							Color: true,
						})
						log.Println(strs[0])
						server.SendError(strs[0])
					} else {
						server.SendReload()
					}
				}
			})
		},
	}
}
